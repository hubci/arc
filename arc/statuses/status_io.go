package statuses

import (
	"fmt"
	"net/http"
	"time"
)

/*
 * This file defines the structs and other components needed to represent a
 * status page hosted by Status.io.
 */

/*
 * statusIOStatus represents the individual component status as well as most
 * of the overall status.
 */
type StatusIOStatus struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
}

/*
 * statusIOReponse represents the JSON response from the StatusPage.io API.
 */
type statusIOResponse struct {
	Result struct {
		StatusOverall struct {
			Updated time.Time `json:"updated"`
			*StatusIOStatus
		} `json:"status_overall"`
	} `json:"result"`
}

/*
 * statusIOProvider represents the way to connect to the API and retrieve a
 * response.
 */
type statusIOProvider struct {
	name   string
	apiURL string
	data   *statusIOResponse
}

/*
 * Fetch pulls in a response from the API.
 */
func (p *statusIOProvider) Fetch(c *http.Client) error {

	return fetchJSON(c, p.apiURL, &p.data)
}

func (p *statusIOProvider) Name() string {
	return p.name
}

func (p *statusIOProvider) URL() string {
	return p.apiURL
}

func (p *statusIOProvider) Status() (string, error) {

	if p.data == nil {
		return "", fmt.Errorf("Data hasn't been retrieved.")
	}

	return p.data.Result.StatusOverall.Status, nil
}

/*
 * RegisterStatusIOPage creates a new provider.
 */
func RegisterStatusIOPage(name, apiURL string) error {
	return registerPage(&statusIOProvider{name: name, apiURL: apiURL})
}
