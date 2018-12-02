package main

type hearStruct struct {
	regex string
	text  string
}

func hear(word string, userid string) string {
	if word == "hello" {
		SendMessage(userid, "Hi there!!")
		return "In this case, we'll trigger anything that will be handled when 'Hello' is triggered"
	}
	return ""
}

func (h *hearStruct) listen(userid string) (string, string) {
	if h.text == "" && h.regex == "" {
		return "", "Oops! Nothing to listen for"
	}

	if h.regex != "" || h.text != "" {
		hear(h.regex, userid)
		hear(h.text, userid)
		return h.regex, h.text
	}
	return "", ""
}
