package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mrjones/oauth"
)

const (
	//Basic OAuth related URLs
	OAUTH_REQUES_TOKEN string = "https://api.twitter.com/oauth/request_token"
	OAUTH_AUTH_TOKEN   string = "https://api.twitter.com/oauth/authorize"
	OAUTH_ACCESS_TOKEN string = "https://api.twitter.com/oauth/access_token"

	//List API URLs
	API_BASE     string = "https://api.twitter.com/1.1/"
	API_TIMELINE string = API_BASE + "statuses/home_timeline.json"
	API_FOLLOWER string = API_BASE + "followers/ids.json"
)

type Client struct {
	HttpConn      *http.Client
	OAuthConsumer *oauth.Consumer
}

func NewClient(consumerKey, consumerSecret string) *Client {
	newClient := &Client{}

	newClient.OAuthConsumer = oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   OAUTH_REQUES_TOKEN,
			AuthorizeTokenUrl: OAUTH_AUTH_TOKEN,
			AccessTokenUrl:    OAUTH_ACCESS_TOKEN,
		},
	)

	//Enable debug info
	newClient.OAuthConsumer.Debug(true)

	return newClient
}

func (c *Client) BasicQuery(queryString string) (interface{}, error) {
	if c.HttpConn == nil {
		return nil, errors.New("No Client OAuth")
	}

	response, err := c.HttpConn.Get(queryString)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var data map[string]interface{}
	bits, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(bits, &data)

	return data, err
}

func (c *Client) QueryTimeLine(count int) (interface{}, error) {
	requesURL := fmt.Sprintf("%s?count=%d", API_TIMELINE, count)
	return c.BasicQuery(requesURL)
}

func (c *Client) QueryFollower(count int) (interface{}, error) {
	requesURL := fmt.Sprintf("%s?count=%d", API_FOLLOWER, count)
	return c.BasicQuery(requesURL)
}
