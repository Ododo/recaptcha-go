package recaptcha

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Response struct {
	success    bool
	errorCodes []string
}

var api_url = "https://www.google.com/recaptcha/api/siteverify"

var privatekey string

func SetPrivateKey(key string) {
	privatekey = key
}

func CheckResponse(response string) bool {
	var r Response
	resp, _ := http.PostForm(api_url,
		url.Values{"secret": {privatekey}, "response": {response}})
	decoder := json.NewDecoder(resp.Body)
	jsonErr := decoder.Decode(&r)
	if jsonErr != nil {
		fmt.Printf("Couldnt decode http response into JSON")
		return false
	}
	if len(r.errorCodes) > 0 {
		fmt.Printf("Received error codes from google api :")
		for _, v := range r.errorCodes {
			fmt.Printf(v)
		}
	}
	return r.success
}
