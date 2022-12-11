package reddit

// Post contains the information about a poss of a subreddit.
type Post struct {
	SubredditName         string             `json:"subreddit"`
	SubredditNamePrefixed string             `json:"subreddit_name_prefixed"`
	SubredditType         string             `json:"subreddit_type"`
	Title                 string             `json:"title"`
	Hidden                bool               `json:"hidden"`
	Edited                interface{}        `json:"edited"`
	UrlOverriddenByDest   string             `json:"url_overridden_by_dest"`
	Author                string             `json:"author"`
	Permalink             string             `json:"permalink"`
	Url                   string             `json:"url"`
	IsVideo               bool               `json:"is_video"`
	MediaId               string             `json:"id"`
	Source                MediaSource        `json:"source"`
	Resolutions           []MediaResolutions `json:"resolutions"`
	Gif                   Gif
	Mp4                   Video
}

// Preview the preview containing all the information needed for a preview.
type Preview struct {
	Images  []Media `json:"images"`
	Enabled bool    `json:"enabled"`
}

// Media type used to store information about the media source and resolutions.
type Media struct {
	Id          string             `json:"id"`
	Source      MediaSource        `json:"source"`
	Resolutions []MediaResolutions `json:"resolutions"`
	Variants    Variants           `json:"variants"`
}

// Variants type used to store information about the variants of media for a specific reddit post.
type Variants struct {
	Gif Gif   `json:"gif"`
	Mp4 Video `json:"mp4"`
}

// Gif type structure that will contain the information about a gif source.
type Gif struct {
	Source      MediaSource
	Resolutions []MediaResolutions
}

// Video type structure that will contain the information about the video source.
type Video struct {
	Source      MediaSource
	Resolutions []MediaResolutions
}

// MediaSource type used to define the source of the media.
type MediaSource struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// MediaResolutions type use to define multiple resolutions for the same source.
type MediaResolutions struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type ApiPostData struct {
	After     interface{}       `json:"after"`
	Dist      int               `json:"dist"`
	Modhash   string            `json:"modhash"`
	GeoFilter string            `json:"geo_filter"`
	Children  []ApiPostChildren `json:"children"`
	Before    interface{}       `json:"before"`
}

type ApiPostChildren struct {
	Kind string              `json:"kind"`
	Data ApiPostChildrenData `json:"data"`
}

type ApiPostChildrenData struct {
	ApprovedAtUtc              interface{}   `json:"approved_at_utc"`
	Subreddit                  string        `json:"subreddit"`
	Selftext                   string        `json:"selftext"`
	UserReports                []interface{} `json:"user_reports"`
	Saved                      bool          `json:"saved"`
	ModReasonTitle             interface{}   `json:"mod_reason_title"`
	Gilded                     int           `json:"gilded"`
	Clicked                    bool          `json:"clicked"`
	Title                      string        `json:"title"`
	LinkFlairRichtext          []interface{} `json:"link_flair_richtext"`
	SubredditNamePrefixed      string        `json:"subreddit_name_prefixed"`
	Hidden                     bool          `json:"hidden"`
	Pwls                       int           `json:"pwls"`
	LinkFlairCssClass          interface{}   `json:"link_flair_css_class"`
	Downs                      int           `json:"downs"`
	ThumbnailHeight            int           `json:"thumbnail_height"`
	TopAwardedType             interface{}   `json:"top_awarded_type"`
	ParentWhitelistStatus      string        `json:"parent_whitelist_status"`
	HideScore                  bool          `json:"hide_score"`
	Name                       string        `json:"name"`
	Quarantine                 bool          `json:"quarantine"`
	LinkFlairTextColor         string        `json:"link_flair_text_color"`
	UpvoteRatio                float64       `json:"upvote_ratio"`
	AuthorFlairBackgroundColor string        `json:"author_flair_background_color"`
	SubredditType              string        `json:"subreddit_type"`
	Ups                        int           `json:"ups"`
	TotalAwardsReceived        int           `json:"total_awards_received"`
	ThumbnailWidth             int           `json:"thumbnail_width"`
	AuthorFlairTemplateId      interface{}   `json:"author_flair_template_id"`
	IsOriginalContent          bool          `json:"is_original_content"`
	AuthorFullName             string        `json:"author_fullname"`
	SecureMedia                interface{}   `json:"secure_media"`
	IsRedditMediaDomain        bool          `json:"is_reddit_media_domain"`
	IsMeta                     bool          `json:"is_meta"`
	Category                   interface{}   `json:"category"`
	LinkFlairText              interface{}   `json:"link_flair_text"`
	CanModPost                 bool          `json:"can_mod_post"`
	Score                      int           `json:"score"`
	ApprovedBy                 interface{}   `json:"approved_by"`
	IsCreatedFromAdsUi         bool          `json:"is_created_from_ads_ui"`
	AuthorPremium              bool          `json:"author_premium"`
	Thumbnail                  string        `json:"thumbnail"`
	Edited                     interface{}   `json:"edited"`
	AuthorFlairCssClass        string        `json:"author_flair_css_class"`
	AuthorFlairRichtext        []interface{} `json:"author_flair_richtext"`
	PostHint                   string        `json:"post_hint"`
	ContentCategories          interface{}   `json:"content_categories"`
	IsSelf                     bool          `json:"is_self"`
	ModNote                    interface{}   `json:"mod_note"`
	Created                    float64       `json:"created"`
	LinkFlairType              string        `json:"link_flair_type"`
	Wls                        int           `json:"wls"`
	RemovedByCategory          interface{}   `json:"removed_by_category"`
	BannedBy                   interface{}   `json:"banned_by"`
	AuthorFlairType            string        `json:"author_flair_type"`
	Domain                     string        `json:"domain"`
	AllowLiveComments          bool          `json:"allow_live_comments"`
	SelftextHtml               interface{}   `json:"selftext_html"`
	Likes                      interface{}   `json:"likes"`
	SuggestedSort              interface{}   `json:"suggested_sort"`
	BannedAtUtc                interface{}   `json:"banned_at_utc"`
	UrlOverriddenByDest        string        `json:"url_overridden_by_dest"`
	ViewCount                  interface{}   `json:"view_count"`
	Archived                   bool          `json:"archived"`
	NoFollow                   bool          `json:"no_follow"`
	IsCrosspostable            bool          `json:"is_crosspostable"`
	Pinned                     bool          `json:"pinned"`
	Over18                     bool          `json:"over_18"`
	Preview                    Preview       `json:"preview"`
	AllAwardings               []interface{} `json:"all_awardings"`
	Awarders                   []interface{} `json:"awarders"`
	MediaOnly                  bool          `json:"media_only"`
	CanGild                    bool          `json:"can_gild"`
	Spoiler                    bool          `json:"spoiler"`
	Locked                     bool          `json:"locked"`
	AuthorFlairText            string        `json:"author_flair_text"`
	TreatmentTags              []interface{} `json:"treatment_tags"`
	Visited                    bool          `json:"visited"`
	RemovedBy                  interface{}   `json:"removed_by"`
	NumReports                 interface{}   `json:"num_reports"`
	Distinguished              interface{}   `json:"distinguished"`
	SubredditId                string        `json:"subreddit_id"`
	AuthorIsBlocked            bool          `json:"author_is_blocked"`
	ModReasonBy                interface{}   `json:"mod_reason_by"`
	RemovalReason              interface{}   `json:"removal_reason"`
	LinkFlairBackgroundColor   string        `json:"link_flair_background_color"`
	Id                         string        `json:"id"`
	IsRobotIndexable           bool          `json:"is_robot_indexable"`
	NumDuplicates              int           `json:"num_duplicates"`
	ReportReasons              interface{}   `json:"report_reasons"`
	Author                     string        `json:"author"`
	DiscussionType             interface{}   `json:"discussion_type"`
	NumComments                int           `json:"num_comments"`
	SendReplies                bool          `json:"send_replies"`
	Media                      RedditMedia   `json:"media"`
	ContestMode                bool          `json:"contest_mode"`
	AuthorPatreonFlair         bool          `json:"author_patreon_flair"`
	AuthorFlairTextColor       string        `json:"author_flair_text_color"`
	Permalink                  string        `json:"permalink"`
	WhitelistStatus            string        `json:"whitelist_status"`
	Stickied                   bool          `json:"stickied"`
	Url                        string        `json:"url"`
	SubredditSubscribers       int           `json:"subreddit_subscribers"`
	CreatedUtc                 float64       `json:"created_utc"`
	NumCrossposts              int           `json:"num_crossposts"`
	ModReports                 []interface{} `json:"mod_reports"`
	IsVideo                    bool          `json:"is_video"`
}

// RedditMedia it is the structure for the root video information.
type RedditMedia struct {
	RedditVideo RedditVideo `json:"reddit_video"`
}

// RedditVideo it is the structure for video information.
type RedditVideo struct {
	BitrateKbps       int    `json:"bitrate_kbps"`
	FallbackUrl       string `json:"fallback_url"`
	Height            int    `json:"height"`
	Width             int    `json:"width"`
	ScrubberMediaUrl  string `json:"scrubber_media_url"`
	DashUrl           string `json:"dash_url"`
	Duration          int    `json:"duration"`
	HlsUrl            string `json:"hls_url"`
	IsGif             bool   `json:"is_gif"`
	TranscodingStatus string `json:"transcoding_status"`
}

// ApiPostResponse is the default structure that the reddit API will return for a post.
type ApiPostResponse struct {
	Kind string      `json:"kind"`
	Data ApiPostData `json:"data"`
}
