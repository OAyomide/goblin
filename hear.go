package main

import (
	"fmt"
)

type hearStruct struct {
	regex string
	text  string
}

func hear(word string, userid string) string {
	if word == "hello" {
		SendMessage(userid, "Hi there!!")
		return "In this case, we'll trigger anything that will be handled when 'Hello' is triggered"
	}

	if word == "GET_STARTED" {
		SendMessage(userid, "Hello there! Seems like this is the first time we're talking! Call me Go Goblin or maybe Goblonio")
		return "GET STARTED TRIGGERED!!"
	}

	if word == "gallery" {
		ddg(userid)
		return "MEDIA QUERY"
	}

	return ""
}

func (h *hearStruct) listen(userid string) {
	if h.text == "" && h.regex == "" {
		panic("Oops! Nothing to listen for")
	}

	if h.regex != "" {
		hear(h.regex, userid)
		fmt.Printf("REGEX PASSED:::%s\n\n", h.regex)
	} else if h.text != "" {
		hear(h.text, userid)
		fmt.Printf("TEXT PASSED:::%s\n\n", h.text)
	}
}
