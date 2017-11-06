package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	//Basic OAuth related URLs
	OAUTH_REQUES_TOKEN string = "https://api.twitter.com/oauth/request_token"
	OAUTH_AUTH_TOKEN   string = "https://api.twitter.com/oauth/authorize"
	OAUTH_ACCESS_TOKEN string = "https://api.twitter.com/oauth/access_token"

	//List API URLs
	API_BASE           string = "https://api.twitter.com/1.1/"
	API_TIMELINE       string = API_BASE + "statuses/home_timeline.json"
	API_FOLLOWERS_IDS  string = API_BASE + "followers/ids.json"
	API_FOLLOWERS_LIST string = API_BASE + "followers/list.json"
	API_FOLLOWER_INFO  string = API_BASE + "users/show.json"
)

type Client struct {
	HttpConn *http.Client
}

func (c *Client) HasAuth() bool {
	return c.HttpConn != nil
}

func (c *Client) BasicQuery(queryString string) ([]byte, error) {
	if c.HttpConn == nil {
		return nil, errors.New("No Client OAuth")
	}

	response, err := c.HttpConn.Get(queryString)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	return bits, err
}

func (c *Client) QueryTimeLine(count int) (TimelineTweets, []byte, error) {
	requesURL := fmt.Sprintf("%s?count=%d", API_TIMELINE, count)
	data, err := c.BasicQuery(requesURL)
	ret := TimelineTweets{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (c *Client) QueryFollower(count int) (Followers, []byte, error) {
	requesURL := fmt.Sprintf("%s?count=%d", API_FOLLOWERS_LIST, count)
	data, err := c.BasicQuery(requesURL)
	ret := Followers{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (c *Client) QueryFollowerIDs(count int) (FollowerIDs, []byte, error) {
	requesURL := fmt.Sprintf("%s?count=%d", API_FOLLOWERS_IDS, count)
	data, err := c.BasicQuery(requesURL)
	var ret FollowerIDs
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (c *Client) QueryFollowerById(id int) (UserDetail, []byte, error) {
	requesURL := fmt.Sprintf("%s?user_id=%d", API_FOLLOWER_INFO, id)
	data, err := c.BasicQuery(requesURL)
	var ret UserDetail
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}
