package reddit

const (
	authUrl   = "https://www.reddit.com"
	baseUrl   = "https://oauth.reddit.com"
	apiPrefix = "api"
	version   = "v1"
)

var (
	EndpointAuth      = authUrl + "/" + apiPrefix + "/" + version + "/access_token"
	EndpointSubreddit = baseUrl
)
