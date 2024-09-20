package statuses

import (
	"fmt"
	"net/http"
)

/*
 * This file defines the structs and other components needed to represent a
 * status page hosted by InstatusStatus.com.
 */

type inStatusPageField struct {
	Name   string `json:"name"`
	URL    string `url:"url"`
	Status string `url:"status"`
}

type inStatusResponse struct {
	Page *inStatusPageField `json:"page"`
}

/*
 * inStatusPage represents the way to connect to the API and retrieve a
 * response.
 */
type inStatusPage struct {
	name   string
	apiURL string
	data   *inStatusResponse
}

/*
 * Fetch pulls in a response from the API.
 */
func (p *inStatusPage) Fetch(c *http.Client) error {

	return fetchJSON(c, p.apiURL, &p.data)
}

func (p *inStatusPage) Name() string {
	return p.name
}

func (p *inStatusPage) URL() string {
	return p.apiURL
}

func (p *inStatusPage) Status() (string, error) {

	if p.data == nil {
		return "", fmt.Errorf("Data hasn't been retrieved.")
	}

	return p.data.Page.Status, nil
}

/*
 * RegisterStatusPageIOPage creates a new provider.
 */
func RegisterInStatusPage(name, apiURL string) error {
	return registerPage(&inStatusPage{name: name, apiURL: apiURL})
}
