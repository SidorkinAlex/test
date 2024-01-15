package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	question := "Вопрос для чат GPT что тяжелее килограм пуха или килограмм бетона"
	jsonString := textToBody(question)
	answer := request(jsonString)
	fmt.Println("question:")
	fmt.Println(question)
	fmt.Println("answer:")
	fmt.Println(answer)

}
func request(bodyString string) string {
	url := "url_to_service"

	payload := strings.NewReader(bodyString)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("Authorization", "Basic dXNlcjE6Y2JzRUFMMTJlTUJqblRaUQ==") //код базовой авторизации
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("accept", "text/event-stream")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func textToBody(textQuestion string) string {
	jsonData := `{
		"action": "_ask",
		"model": "gpt-4-turbo-stream-you",
		"jailbreak": "default",
		"meta": {
			"id": "7322001447494789451",
			"content": {
				"conversation": [],
				"internet_access": false,
				"content_type": "text",
				"parts": [
					{
						"content": "П",
						"role": "user"
					}
				]
			}
		}
	}`

	var data JSONData
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	// Modify the Go structures as needed
	data.Meta.Content.Parts[0].Content = textQuestion

	// Marshal the modified Go structures back into JSON
	modifiedJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	return string(modifiedJSON)
}

type Part struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type Content struct {
	Conversation   []interface{} `json:"conversation"`
	InternetAccess bool          `json:"internet_access"`
	ContentType    string        `json:"content_type"`
	Parts          []Part        `json:"parts"`
}

type Meta struct {
	ID      string  `json:"id"`
	Content Content `json:"content"`
}

type JSONData struct {
	Action    string `json:"action"`
	Model     string `json:"model"`
	Jailbreak string `json:"jailbreak"`
	Meta      Meta   `json:"meta"`
}
