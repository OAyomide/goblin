package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type config struct {
	VerifyToken string
	AccessToken string
	AppSecret   string
}

type TextReplystruct struct {
	Text string `json:"text"`
}
type TextReplyRecipientstruct struct {
	ID string `json:"id"`
}
type Vertex struct {
	X string                   `json:"message_type"`
	I TextReplyRecipientstruct `json:"recipient"`
	Y TextReplystruct          `json:"message"`
}

type Profile struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name:`
	ProfilePic string `json:"profile_pic"`
}

type CallSendApiResponse struct {
	RecID     string `json:"recipient_id"`
	MessageID string `json:"message_id"`
}

var tk config
var v = getToken()
var marshalError = json.Unmarshal([]byte(v), &tk)

func webhookGetHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Here is the token! ==> %s\n", tk.VerifyToken)

	token := tk.VerifyToken
	//the token gotten from our req object
	tokenTrue := r.URL.Query().Get("hub.verify_token")
	hubChallenge := r.URL.Query().Get("hub.challenge")
	if tokenTrue == token {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(hubChallenge))
	} else {
		fmt.Fprint(w, "Nay! Tokens don't match")
	}
}

func webhookPostHandler(w http.ResponseWriter, r *http.Request) {

	type Postbackstruct struct {
		Payload string
	}

	type Payloadstruct struct {
		URL      string
		Reusable bool `json:"is_reusable"`
	}
	type Attachmentstruct struct {
		Type    string
		Payload Payloadstruct
	}
	type Messagestruct struct {
		Text        string
		Attachments []*Attachmentstruct
	}

	type Recipientstruct struct {
		ID string
	}

	type Senderstruct struct {
		ID string
	}

	type Body struct {
		Object string
		Entry  []struct {
			ID        string
			Time      int64
			Messaging []struct {
				Timestamp int64
				Sender    Senderstruct
				Recipient Recipientstruct
				Message   *Messagestruct
				Postback  *Postbackstruct
			}
		}
	}
	//we want to parse our request object
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error parsing the body: %s \n", err)
	}

	var body Body
	json.Unmarshal([]byte(data), &body)

	if body.Object == "page" {
		for _, entries := range body.Entry {
			for _, messaging := range entries.Messaging {
				if messaging.Message != nil {
					attachments := messaging.Message.Attachments
					text := messaging.Message.Text

					if attachments != nil {
						SendMessage(messaging.Sender.ID, "Oops! Cant do that yet")
						fmt.Println("Attachment. Cannot process")
					} else if messaging.Message.Text != "" {
						SendMessage(messaging.Sender.ID, text)
					}

				} else if messaging.Postback != nil {
					fmt.Println("Yay! We have a postback event!")
				}
			}
		}
	}
}

func callSendAPI(data []byte) {
	accessToken := tk.AccessToken
	response, err := http.Post("https://graph.facebook.com/v2.6/me/messages?access_token="+accessToken, "application/json", bytes.NewBuffer(data))

	if err != nil {
		log.Printf("Error here!! %s", err)
		panic(err)
	}

	res, _ := ioutil.ReadAll(response.Body)
	var rs CallSendApiResponse
	json.Unmarshal([]byte(res), &rs)
	fmt.Printf("The Message ID is: %s and recipient ID is: %s", rs.MessageID, rs.RecID)
	fmt.Printf("Log of the transaction here!! Response: %s\n\n", string(res))
}

func getUserProfile(userID string) string {
	accessToken := tk.AccessToken

	profileFields := []string{"first_name", "last_name", "name", "profile_pic"}
	separatedUserFields := strings.Join(profileFields, ",")
	response, err := http.Get("https://graph.facebook.com/v3.1/" + userID + "?fields=" + separatedUserFields + "&access_token=" + accessToken)

	if err != nil {
		fmt.Printf(">>ERROR ACCESSING THE USER PROFILE: %s<<\n\n", err)
		panic(err)
	}

	defer response.Body.Close()
	res, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("All green! Got user profile: %v\n\n", string(res))
	return string(res)
}

func SendMessage(UserID string, text string) {
	userProfile := getUserProfile(UserID)
	var prof Profile

	json.Unmarshal([]byte(userProfile), &prof)
	i := TextReplyRecipientstruct{UserID}
	t := TextReplystruct{text}
	send, _ := json.Marshal(Vertex{"RESPONSE", i, t})
	callSendAPI(send)
}
