package twitter

import (
	"fmt"
	"log"
)

func NewDesktopClient(consumerKey, consumerSecret string) *DesktopClient {
	newClient := NewClient(consumerKey, consumerKey)
	newServer := new(DesktopClient)
	newServer.Client = *newClient
	return newServer
}

type DesktopClient struct {
	Client
}

func (d *DesktopClient) DoAuth() error {
	requestToken, u, err := d.OAuthConsumer.GetRequestTokenAndUrl("oob")
	fmt.Println("rest token=", requestToken, " err=", err)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("(1) Go to: " + u)
	fmt.Println("(2) Grant access, you should get back a verification code.")
	fmt.Println("(3) Enter that verification code here: ")

	verificationCode := ""
	fmt.Scanln(&verificationCode)

	accessToken, err := d.OAuthConsumer.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		log.Fatal(err)
	}

	d.HttpConn, err = d.OAuthConsumer.MakeHttpClient(accessToken)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
