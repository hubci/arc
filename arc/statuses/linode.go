package statuses

/*
 * Register the CircleCI status page.
 */
func init() {
	RegisterStatusPageIOPage("linode", "https://status.linode.com/api/v2/status.json")
}
