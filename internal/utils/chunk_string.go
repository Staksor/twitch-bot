package utils

import "unicode/utf8"

func ChunkString(longString string, limit int) []string {
	splits := []string{}

	var l, r int
	for l, r = 0, limit; r < len(longString); l, r = r, r+limit {
		for !utf8.RuneStart(longString[r]) {
			r--
		}
		splits = append(splits, longString[l:r])
	}
	splits = append(splits, longString[l:])

	return splits
}
