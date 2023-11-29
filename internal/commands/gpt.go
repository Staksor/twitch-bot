package commands

import (
	"bot/internal/structs"
	"bot/internal/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

// Makes a request to an API of a GPT chat bot
func Gpt(
	message twitch.PrivateMessage,
	client *twitch.Client,
	chatMessage string,
	gptResponseStates map[string]*structs.GptResponseState,
	cooldowns map[string]*time.Time,
) {
	if len(chatMessage) <= 3 {
		client.Reply(message.Channel, message.ID, "Please write a longer message")

		return
	}

	iniData := utils.GetIniData()

	gptCooldown, _ := iniData.Section("cooldowns").Key("gpt").Int()
	if !utils.CheckCooldown("gpt", gptCooldown, message, client, cooldowns) {
		return
	}

	Request := structs.GptRequest{
		Application: iniData.Section("api").Key("gpt_bot_key").String(),
		Instance:    iniData.Section("api").Key("gpt_bot_id").String(),
		Message:     fmt.Sprintf("(my name is %s. answer only to the message inside the quotes as if you just met me, but you know my name) \"%s\" (!!!be brief, don't write answers longer than 700 characters!!!)", message.User.DisplayName, message.Message),
	}

	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(&Request)
	var jsonData = []byte(buffer.String())

	httpClient := &http.Client{}
	req, _ := http.NewRequest("POST", "https://www.botlibre.com/rest/json/chat", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	res, _ := httpClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	bodyString := string(body)

	var gptResponse structs.GptResponse

	if err := json.Unmarshal([]byte(bodyString), &gptResponse); err == nil {
		var gptResponseMessages []string = utils.ChunkString(gptResponse.Message, 200)
		client.Reply(message.Channel, message.ID, "🤖 "+gptResponseMessages[0])

		if len(gptResponseMessages) > 1 {
			userState := new(structs.GptResponseState)
			userState.Messages = gptResponseMessages[1:]
			gptResponseStates[message.User.ID] = userState
			client.Reply(message.Channel, message.ID, "type !continue for more")
		}
	} else {
		client.Reply(message.Channel, message.ID, fmt.Sprintf("There was an error in the API response PoroSad (%s)", bodyString))
		fmt.Println(err)
	}
}
