package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type client struct {
	c *http.Client
}

func (this *client) getJSON(url string, payload interface{}) error {

	resp, err := this.c.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error: API response was not HTTP 200.")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, payload)
}

func New() *client {

	return &client{&http.Client{Timeout: 4 * time.Second}}

}
