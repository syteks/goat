package reddit

// All the API constants used to construct our reddit endpoints.
const (
	authUrl   = "https://www.reddit.com"
	baseUrl   = "https://oauth.reddit.com"
	apiPrefix = "api"
	version   = "v1"
)

// All the endpoints for the reddit API.
var (
	EndpointAuth      = authUrl + "/" + apiPrefix + "/" + version + "/access_token"
	EndpointSubreddit = baseUrl
)
