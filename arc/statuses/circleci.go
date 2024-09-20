package statuses

/*
 * Register the CircleCI status page.
 */
func init() {
	RegisterStatusPageIOPage("circleci", "https://status.circleci.com/api/v2/status.json")
}
