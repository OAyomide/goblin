package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goblin/parser"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// type config struct {
// 	VerifyToken string
// 	AccessToken string
// 	AppSecret   string
// }

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

/**
* Our Error json struct
 */
type Error struct {
	message string `json:"message", omitempty`
}
type ErrorString struct {
	Er Error `json:"error"`
}
type ErrorMessageStruct struct {
	Error ErrorString `json:"error"`
}

//======STRUCT FOR OUR BODY=======
type Postbackstruct struct {
	Payload string `json:"payload"`
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
type MessagingStruct struct {
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
			Postback  *Postbackstruct `json:"postback"`
		}
	}
}

type GetStarted struct {
	Payload string `json:"payload"`
}

type getStartedButton struct {
	GetS GetStarted `json:"get_started"`
}

var t = parser.GetAccessToken()

func webhookGetHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Here is the token! ==> %s\n", t)

	token := t
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

//+14172892061
func webhookPostHandler(w http.ResponseWriter, r *http.Request) {

	//we want to parse our request object
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error parsing the body: %s \n", err)
	}

	var body Body
	json.Unmarshal([]byte(data), &body)

	if body.Object == "page" {
		fmt.Printf("The whole object received::%v\n\n", string(data))
		for _, entries := range body.Entry {
			for _, messaging := range entries.Messaging {
				if messaging.Message != nil {
					attachments := messaging.Message.Attachments
					text := messaging.Message.Text

					if attachments != nil {
						SendMessage(messaging.Sender.ID, "Oops! Cant do that yet")
						fmt.Println("Attachment. Cannot process")
					} else if messaging.Message.Text != "" {
						fmt.Printf("THE USER ID IS::%s", messaging.Sender.ID)
						vr := hearStruct{text: text}
						vr.listen(messaging.Sender.ID)
					}

				} else if messaging.Postback != nil {
					fmt.Println("Yay! We have a postback event!")

					if messaging.Postback.Payload == "GET_STARTED" {
						SendMessage(messaging.Sender.ID, "Hello! It seems this is the first time we're talking! Call me Goblin!")
					}
				}
			}
		}
		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Printf("SEEMS LIKE THE POSTBACK OF THE GET STARTED BUTTON:: %v", data)
	}
}

func callSendAPI(data []byte) {
	//accessToken := tk.AccessToken
	accessToken := t
	response, err := http.Post("https://graph.facebook.com/v2.6/me/messages?access_token="+accessToken, "application/json", bytes.NewBuffer(data))

	if err != nil {
		log.Printf("Error here!! %s", err)
		panic("CALLSEND API ERR:" + err.Error())
	}

	res, _ := ioutil.ReadAll(response.Body)
	var rs CallSendApiResponse
	json.Unmarshal([]byte(res), &rs)

	if rs.MessageID == "" {
		fmt.Printf("Error with CallsendAPI here:%s", string(res))
	}
}

func getUserProfile(userID string) string {
	//accessToken := tk.AccessToken
	accessToken := t

	profileFields := []string{"first_name", "last_name", "name", "profile_pic"}
	separatedUserFields := strings.Join(profileFields, ",")
	response, err := http.Get("https://graph.facebook.com/v3.1/" + userID + "?fields=" + separatedUserFields + "&access_token=" + accessToken)

	if err != nil {
		fmt.Printf(">>ERROR ACCESSING THE USER PROFILE: %s<<\n\n", err)
		panic(err)
	}

	res, _ := ioutil.ReadAll(response.Body)
	// fmt.Printf("All green! Got user profile: %v\n\n", string(res))
	//defer response.Body.Close()
	return string(res)
}

//SendMessage is used to send the text to the user on messenger
func SendMessage(UserID string, text string) {
	userProfile := getUserProfile(UserID)
	var prof Profile

	json.Unmarshal([]byte(userProfile), &prof)
	i := TextReplyRecipientstruct{UserID}
	t := TextReplystruct{text}
	send, _ := json.Marshal(Vertex{"RESPONSE", i, t})
	callSendAPI(send)
}

func setKV(ppt string, val string) string {
	rt := map[string]string{ppt: val}
	return createKeyValuePairs(rt)
}

func createKeyValuePairs(val map[string]string) string {
	v := new(bytes.Buffer)
	for key, value := range val {
		fmt.Fprintf(v, "\"%s:\"%s\"\n", key, value)
	}
	return v.String()
}

//set the get started payload for the GET_STARTED button
func setGetStartedPayload(value string) {

	accessToken := t
	fn := getStartedButton{
		GetStarted{
			Payload: value,
		},
	}

	data, err := json.Marshal(fn)

	if err != nil {
		fmt.Print("Error marshalling GET STARTED")
		panic(err)
	}
	//send the request
	response, err := http.Post("https://graph.facebook.com/v3.1/me/messenger_profile?access_token="+accessToken, "application/json", bytes.NewBuffer(data))

	if err != nil {
		fmt.Printf("Error setting get started payload here: %s", err.Error())
	}

	rs, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Error parsing POST response here: %s", err.Error())
	}

	fmt.Printf("HERE IS THE RESPONSE AFTER MAKING POST REQUEST TO SET PAYLOAD: %s\n\n", string(rs))
}
