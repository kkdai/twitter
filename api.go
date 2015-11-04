package twitter

import "errors"

func NewTwitterClient(consumerKey, consumerSecret string, useServerClient bool) *Twitter {
	newClient := new(Twitter)

	if useServerClient {
		newClient.client = NewServerClient(consumerKey, consumerSecret)
	} else {
		newClient.client = NewDesktopClient(consumerKey, consumerSecret)
	}

	newClient.useServerClient = useServerClient
	return newClient
}

type Twitter struct {
	client          interface{}
	useServerClient bool
}

func (t *Twitter) getServerClient() *ServerClient {
	if srv, ok := t.client.(ServerClient); ok {
		return &srv
	}

	return nil
}

func (t *Twitter) getDesktopClient() *DesktopClient {
	if dst, ok := t.client.(DesktopClient); ok {
		return &dst
	}

	return nil
}

func (t *Twitter) GetAuthURL(tokenUrl string) string {
	if t.useServerClient {
		return t.getServerClient().GetAuthURL(tokenUrl)
	} else {
		return ""
	}

}

func (t *Twitter) CompleteAuth(tokenKey, verificationCode string) error {
	if t.useServerClient {
		return t.getServerClient().CompleteAuth(tokenKey, verificationCode)
	} else {
		return errors.New("New server client")
	}
}

func (t *Twitter) QueryTimeLine(count int) (interface{}, error) {
	if t.useServerClient {
		return t.getServerClient().QueryTimeLine(count)
	} else {
		return t.getDesktopClient().QueryTimeLine(count)

	}
}

func (t *Twitter) QueryFollower(count int) (interface{}, error) {
	if t.useServerClient {
		return t.getServerClient().QueryTimeLine(count)
	} else {
		return t.getDesktopClient().QueryTimeLine(count)

	}
}

func (t *Twitter) QueryTest(url string) (interface{}, error) {
	if t.useServerClient {
		return t.getServerClient().BasicQuery(url)
	} else {
		return t.getDesktopClient().BasicQuery(url)

	}
}
