package statuses

import (
	"fmt"
	"net/http"
	"time"
)

/*
 * This file defines the structs and other components needed to represent a
 * status page hosted by StatusPage.io.
 */

type statusPageIOPageField struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `url:"url"`
	TimeZone  string    `json:"time_zone"`
	UpdatedAt time.Time `json:"updated_at"`
}

type statusPageIOStatus struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}

type statusPageIOResponse struct {
	Page   *statusPageIOPageField `json:"page"`
	Status *statusPageIOStatus    `json:"status"`
}

/*
 * statusPageIOPage represents the way to connect to the API and retrieve a
 * response.
 */
type statusPageIOPage struct {
	name   string
	apiURL string
	data   *statusPageIOResponse
}

/*
 * Fetch pulls in a response from the API.
 */
func (p *statusPageIOPage) Fetch(c *http.Client) error {

	return fetchJSON(c, p.apiURL, &p.data)
}

func (p *statusPageIOPage) Name() string {
	return p.name
}

func (p *statusPageIOPage) URL() string {
	return p.apiURL
}

func (p *statusPageIOPage) Status() (string, error) {

	if p.data == nil {
		return "", fmt.Errorf("Data hasn't been retrieved.")
	}

	return p.data.Status.Description, nil
}

/*
 * RegisterStatusPageIOPage creates a new provider.
 */
func RegisterStatusPageIOPage(name, apiURL string) error {
	return registerPage(&statusPageIOPage{name: name, apiURL: apiURL})
}
