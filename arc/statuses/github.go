package statuses

/*
 * Register the CircleCI status page.
 */
func init() {
	RegisterStatusPageIOPage("github", "https://www.githubstatus.com/api/v2/status.json")
}
