package reddit

import (
	"encoding/json"
	"fmt"
	"goat/packages/handlers/custom_error"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// AuthorizeUsingDefaultCredentials authorize the use of reddit API using the default credentials found in the .env file.
func AuthorizeUsingDefaultCredentials() {
	Authorize(Credentials{
		UserAgent:    os.Getenv("REDDIT_USER_AGENT"),
		ClientID:     os.Getenv("REDDIT_CLIENT_ID"),
		ClientSecret: os.Getenv("REDDIT_CLIENT_SECRET"),
		Username:     os.Getenv("REDDIT_USERNAME"),
		Password:     os.Getenv("REDDIT_PASSWORD"),
	})
}

// Authorize the API to access the data from reddit
func Authorize(credentials Credentials) {
	// The form data required to authorize the use of reddit API.
	formData := url.Values{}
	formData.Add("grant_type", "password")
	formData.Add("username", credentials.Username)
	formData.Add("password", credentials.Password)

	request, err := http.NewRequest("POST", EndpointAuth, strings.NewReader(formData.Encode()))

	custom_error.Handle(err, "There was an error setting up the request call for the authorization.")

	// Set the required headers in order to have a successful authorization.
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", os.Getenv("REDDIT_USER_AGENT"))
	request.Header.Set("Authorization", basicAuth(credentials.ClientID, credentials.ClientSecret))

	response, err := http.DefaultClient.Do(request)

	custom_error.Handle(err, "There was an error getting the authorization from the Reddit API.")

	body, err := io.ReadAll(response.Body)

	custom_error.Handle(err, "Could not read the authorization API's response.")

	err = json.Unmarshal(body, &Auth)

	custom_error.Handle(err, "Could not parse the json response.")

	Auth.ExpiringDate = time.Now().Add(time.Second * time.Duration(Auth.ExpiresIn))

	// Close the response body after is has been used to init the auth variable.
	custom_error.Handle(response.Body.Close(), "There was an error closing the authorization response's body.")
}

// get is a function that will call the reddit API with the given endpoint and params in order to fetch data.
// Returns the body of the API call.
func get(endpoint string, params map[string]string) []byte {
	request, err := http.NewRequest("GET", endpoint, nil)

	custom_error.Handle(err, fmt.Sprintf("The getter for the endpoint failed. Endpoint {%s}", endpoint))

	request.Header.Set("User-Agent", os.Getenv("REDDIT_USER_AGENT"))
	request.Header.Set("Authorization", bearerTokenAuth(Auth.AccessToken))

	// Add each one of the query params for the endpoint to the request.
	if len(params) > 0 {
		requestQuery := request.URL.Query()

		for key, value := range params {
			requestQuery.Add(key, value)
		}

		request.URL.RawQuery = requestQuery.Encode()

	}
	response, err := http.DefaultClient.Do(request)

	custom_error.Handle(err, "Could not fetch the response from the reddit API.")

	body, err := io.ReadAll(response.Body)

	custom_error.Handle(err, "Could not read the response given by the reddit API.")

	custom_error.Handle(response.Body.Close(), "There was an error closing the authorization response's body.")

	return body
}

func GetSubredditPosts(subredditName string, subredditType string) ([]Post, error) {
	if time.Now().After(Auth.ExpiringDate) {
		AuthorizeUsingDefaultCredentials()
	}

	defaultSubreddit := "r"

	if subredditType == "" {
		subredditType = defaultSubreddit
	}

	endpoint := EndpointSubreddit + "/" + subredditType + "/" + subredditName

	params := make(map[string]string)

	params["limit"] = "100"

	body := get(endpoint, params)

	var apiPostResponse ApiPostResponse

	err := json.Unmarshal(body, &apiPostResponse)

	custom_error.Handle(err, "There was an error un-parsing the JSON data.")

	var posts []Post

	if err != nil {
		return posts, err
	}

	for _, post := range apiPostResponse.Data.Children {
		// Continue if our post doesn't have an image.
		if len(post.Data.Preview.Images) == 0 {
			continue
		}

		videoSource := post.Data.Preview.Images[0].Variants.Mp4

		if post.Data.IsVideo && len(strings.TrimSpace(videoSource.Source.Url)) == 0 {
			videoSource = Video{
				Source: MediaSource{
					Url:    post.Data.Media.RedditVideo.FallbackUrl,
					Width:  0,
					Height: 0,
				},
			}
		}

		posts = append(posts, Post{
			post.Data.Subreddit,
			post.Data.SubredditNamePrefixed,
			post.Data.SubredditType,
			post.Data.Title,
			post.Data.Hidden,
			post.Data.Edited,
			post.Data.UrlOverriddenByDest,
			post.Data.Author,
			post.Data.Permalink,
			post.Data.Url,
			post.Data.IsVideo,
			post.Data.Preview.Images[0].Id,
			post.Data.Preview.Images[0].Source,
			post.Data.Preview.Images[0].Resolutions,
			post.Data.Preview.Images[0].Variants.Gif,
			videoSource,
		})
	}

	return posts, err
}
