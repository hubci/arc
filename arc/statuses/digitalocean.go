package statuses

/*
 * Register the CircleCI status page.
 */
func init() {
	RegisterStatusPageIOPage("digitalocean", "https://status.digitalocean.com/api/v2/status.json")
}
