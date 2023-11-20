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
func Gpt(message twitch.PrivateMessage, client *twitch.Client, chatMessage string) {
	if len(chatMessage) <= 3 {
		client.Reply(message.Channel, message.ID, "Please write a longer message")

		return
	}

	iniData := utils.GetIniData()

	Request := structs.GptRequest{
		Application: iniData.Section("api").Key("gpt_bot_key").String(),
		Instance:    iniData.Section("api").Key("gpt_bot_id").String(),
		Message:     chatMessage,
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
		gptResponseMessages := utils.ChunkString(gptResponse.Message, 450)
		for _, gptResponseMessage := range gptResponseMessages {
			client.Reply(message.Channel, message.ID, gptResponseMessage)
			time.Sleep(200 * time.Millisecond)
		}
	} else {
		fmt.Println(err)
	}
}
