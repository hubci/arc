package statuses

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StatusPage interface {
	Fetch(*http.Client) error
	Name() string
	URL() string
	Status() (string, error)
}

func fetchJSON(c *http.Client, url string, payload interface{}) error {

	resp, err := c.Get(url)
	if err != nil {
		return fmt.Errorf("Failed to GET URL. Err: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error: API response was not HTTP 200.")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to read GET response. Err: %s", err)
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal GET response. Err: %s", err)
	}

	return nil
}
