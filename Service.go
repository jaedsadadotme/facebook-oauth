package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var client_id = os.Getenv("CILENT_ID")
var client_secret = os.Getenv("CILENT_SECRET")
var redirect_uri = "http://localhost:1323/"

func getTokenBearer(code string) (Token, ErrorResponse) {
	uri := fmt.Sprintf(`https://graph.facebook.com/v8.0/oauth/access_token?client_id=%s&redirect_uri=%s&client_secret=%s&code=%s`, client_id, redirect_uri, client_secret, code)
	resp, err := http.Get(uri)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	if err := data["error"]; err != nil {
		err := data["error"].(map[string]interface{})
		return Token{}, ErrorResponse{
			Error:        http.StatusText(http.StatusBadRequest),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err["message"],
		}
	}
	return Token{AccessToken: data["access_token"].(string)}, ErrorResponse{}
}
func getUserInfo(token string) Response {
	uri := "https://graph.facebook.com/v8.0/me?fields=id,name,email"
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body2, _ := ioutil.ReadAll(response.Body)
	/// return
	var result Response
	json.Unmarshal(body2, &result.Datas)
	result.Datas.Image = fmt.Sprintf("https://graph.facebook.com/%s/picture?type=large", result.Datas.ID)
	return result
}
