package util
import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"log"
	"net/url"
)

func auth() string {
	var token = os.Getenv("TWITTER_BEARER_TOKEN")
	var bearerToken = fmt.Sprintf("Bearer %s", token)
	return bearerToken
}

func create_url() string {
	// Query params
	params := url.Values{}
	params.Add("tweet.fields", "lang,author_id,in_reply_to_user_id")
	params.Add("ids", "3165112641,14877483,21126305")
	params.Add("expansions", "pinned_tweet_id")
	url := url.URL{
		Scheme: "https",
		Host:   "api.twitter.com",
		Path: "/2/users",
		RawQuery: params.Encode(),
	}
	var t_url = url.String()
	return t_url
}


func MakeRequest () ([]uint8, error) {
	var url = create_url()

	// Create a Bearer string by appending string access token
	var bearer = auth()

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
