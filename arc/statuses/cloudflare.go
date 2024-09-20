package statuses

/*
 * Register the CircleCI status page.
 */
func init() {
	RegisterStatusPageIOPage("cloudflare", "https://www.cloudflarestatus.com/api/v2/status.json")
}
