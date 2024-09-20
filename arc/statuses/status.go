package statuses

import (
	"fmt"
	"net/url"
)

// This is how the package keeps track of all the status pages that it knows
// about.
var registeredPages map[string]StatusPage = map[string]StatusPage{}

/*
 * GetPageURL returns a status page's API URL.
 */
func Page(name string) (StatusPage, error) {

	sp, ok := registeredPages[name]
	if !ok {
		return nil, fmt.Errorf("Not a registered page.")
	}

	return sp, nil
}

/*
 * registerPage adds a new status page to the list for a specific company /
 * provider.
 *
 * name - should be a unique slug
 */
func registerPage(provider StatusPage) error {

	name := provider.Name()

	// If the name has already been used.
	if _, ok := registeredPages[name]; ok {
		return newNameTakenError(name)
	}

	// If the URL provided is no good.
	if _, err := url.Parse(provider.URL()); err != nil {
		return err
	}

	registeredPages[name] = provider

	return nil
}
