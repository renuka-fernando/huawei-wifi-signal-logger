package signal

import (
	"encoding/xml"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

const (
	ApiPath   = "api/monitoring/status"
	SessionID = "SessionID"
)

type Client struct {
	IP        string
	SessionId string
}

type State struct {
	SignalStrength int `xml:"SignalIcon"`
}

func (c *Client) GetStrength() (int, error) {
	client := resty.New()
	client.SetCookie(&http.Cookie{
		Name:     SessionID,
		Value:    c.SessionId,
		Path:     "/",
		Domain:   c.IP,
		HttpOnly: true,
		Secure:   false,
	})
	resp, err := client.R().Get(fmt.Sprintf("http://%s/%s", c.IP, ApiPath))
	if err != nil {
		return 0, err
	}

	st := &State{}
	if err = xml.Unmarshal(resp.Body(), st); err != nil {
		return 0, err
	}

	return st.SignalStrength, nil
}
