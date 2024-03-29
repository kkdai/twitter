package twitter

type TimelineTweets []struct {
	CreatedAt            string      `json:"created_at"`
	ID                   int64       `json:"id"`
	IDStr                string      `json:"id_str"`
	Text                 string      `json:"text"`
	Source               string      `json:"source"`
	Truncated            bool        `json:"truncated"`
	InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
	InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
	InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
	User                 struct {
		ID          int    `json:"id"`
		IDStr       string `json:"id_str"`
		Name        string `json:"name"`
		ScreenName  string `json:"screen_name"`
		Location    string `json:"location"`
		Description string `json:"description"`
		URL         string `json:"url"`
		Entities    struct {
			URL struct {
				Urls []struct {
					URL         string `json:"url"`
					ExpandedURL string `json:"expanded_url"`
					DisplayURL  string `json:"display_url"`
					Indices     []int  `json:"indices"`
				} `json:"urls"`
			} `json:"url"`
			Description struct {
				Urls []interface{} `json:"urls"`
			} `json:"description"`
		} `json:"entities"`
		Protected                      bool   `json:"protected"`
		FollowersCount                 int    `json:"followers_count"`
		FriendsCount                   int    `json:"friends_count"`
		ListedCount                    int    `json:"listed_count"`
		CreatedAt                      string `json:"created_at"`
		FavouritesCount                int    `json:"favourites_count"`
		UtcOffset                      int    `json:"utc_offset"`
		TimeZone                       string `json:"time_zone"`
		GeoEnabled                     bool   `json:"geo_enabled"`
		Verified                       bool   `json:"verified"`
		StatusesCount                  int    `json:"statuses_count"`
		Lang                           string `json:"lang"`
		ContributorsEnabled            bool   `json:"contributors_enabled"`
		IsTranslator                   bool   `json:"is_translator"`
		IsTranslationEnabled           bool   `json:"is_translation_enabled"`
		ProfileBackgroundColor         string `json:"profile_background_color"`
		ProfileBackgroundImageURL      string `json:"profile_background_image_url"`
		ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https"`
		ProfileBackgroundTile          bool   `json:"profile_background_tile"`
		ProfileImageURL                string `json:"profile_image_url"`
		ProfileImageURLHTTPS           string `json:"profile_image_url_https"`
		ProfileBannerURL               string `json:"profile_banner_url"`
		ProfileLinkColor               string `json:"profile_link_color"`
		ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
		ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
		ProfileTextColor               string `json:"profile_text_color"`
		ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
		HasExtendedProfile             bool   `json:"has_extended_profile"`
		DefaultProfile                 bool   `json:"default_profile"`
		DefaultProfileImage            bool   `json:"default_profile_image"`
		Following                      bool   `json:"following"`
		FollowRequestSent              bool   `json:"follow_request_sent"`
		Notifications                  bool   `json:"notifications"`
	} `json:"user"`
	Geo           interface{} `json:"geo"`
	Coordinates   interface{} `json:"coordinates"`
	Place         interface{} `json:"place"`
	Contributors  interface{} `json:"contributors"`
	IsQuoteStatus bool        `json:"is_quote_status"`
	RetweetCount  int         `json:"retweet_count"`
	FavoriteCount int         `json:"favorite_count"`
	Entities      struct {
		Hashtags     []interface{} `json:"hashtags"`
		Symbols      []interface{} `json:"symbols"`
		UserMentions []interface{} `json:"user_mentions"`
		Urls         []interface{} `json:"urls"`
	} `json:"entities"`
	Favorited bool   `json:"favorited"`
	Retweeted bool   `json:"retweeted"`
	Lang      string `json:"lang"`
}

type Followers struct {
	Users []struct {
		ID                             int64       `json:"id"`
		IDStr                          string      `json:"id_str"`
		Name                           string      `json:"name"`
		ScreenName                     string      `json:"screen_name"`
		Location                       string      `json:"location"`
		ProfileLocation                interface{} `json:"profile_location"`
		URL                            interface{} `json:"url"`
		Description                    string      `json:"description"`
		Protected                      bool        `json:"protected"`
		FollowersCount                 int         `json:"followers_count"`
		FriendsCount                   int         `json:"friends_count"`
		ListedCount                    int         `json:"listed_count"`
		CreatedAt                      string      `json:"created_at"`
		FavouritesCount                int         `json:"favourites_count"`
		UtcOffset                      interface{} `json:"utc_offset"`
		TimeZone                       interface{} `json:"time_zone"`
		GeoEnabled                     bool        `json:"geo_enabled"`
		Verified                       bool        `json:"verified"`
		StatusesCount                  int         `json:"statuses_count"`
		Lang                           string      `json:"lang"`
		ContributorsEnabled            bool        `json:"contributors_enabled"`
		IsTranslator                   bool        `json:"is_translator"`
		IsTranslationEnabled           bool        `json:"is_translation_enabled"`
		ProfileBackgroundColor         string      `json:"profile_background_color"`
		ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
		ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
		ProfileBackgroundTile          bool        `json:"profile_background_tile"`
		ProfileImageURL                string      `json:"profile_image_url"`
		ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
		ProfileLinkColor               string      `json:"profile_link_color"`
		ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
		ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
		ProfileTextColor               string      `json:"profile_text_color"`
		ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
		DefaultProfile                 bool        `json:"default_profile"`
		DefaultProfileImage            bool        `json:"default_profile_image"`
		Following                      bool        `json:"following"`
		FollowRequestSent              bool        `json:"follow_request_sent"`
		Notifications                  bool        `json:"notifications"`
		Muting                         bool        `json:"muting"`
	} `json:"users"`
	NextCursor        int64  `json:"next_cursor"`
	NextCursorStr     string `json:"next_cursor_str"`
	PreviousCursor    int    `json:"previous_cursor"`
	PreviousCursorStr string `json:"previous_cursor_str"`
}

type UserTimeline []struct {
	CreatedAt string `json:"created_at"`
	ID        int64  `json:"id"`
	IDStr     string `json:"id_str"`
	Text      string `json:"text"`
	Truncated bool   `json:"truncated"`
	Entities  struct {
		Hashtags     []interface{} `json:"hashtags"`
		Symbols      []interface{} `json:"symbols"`
		UserMentions []struct {
			ScreenName string `json:"screen_name"`
			Name       string `json:"name"`
			ID         int64  `json:"id"`
			IDStr      string `json:"id_str"`
			Indices    []int  `json:"indices"`
		} `json:"user_mentions"`
		Urls []struct {
			URL         string `json:"url"`
			ExpandedURL string `json:"expanded_url"`
			DisplayURL  string `json:"display_url"`
			Indices     []int  `json:"indices"`
		} `json:"urls"`
	} `json:"entities"`
	Source               string      `json:"source"`
	InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
	InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
	InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
	User                 struct {
		ID          int    `json:"id"`
		IDStr       string `json:"id_str"`
		Name        string `json:"name"`
		ScreenName  string `json:"screen_name"`
		Location    string `json:"location"`
		Description string `json:"description"`
		URL         string `json:"url"`
		Entities    struct {
			URL struct {
				Urls []struct {
					URL         string `json:"url"`
					ExpandedURL string `json:"expanded_url"`
					DisplayURL  string `json:"display_url"`
					Indices     []int  `json:"indices"`
				} `json:"urls"`
			} `json:"url"`
			Description struct {
				Urls []interface{} `json:"urls"`
			} `json:"description"`
		} `json:"entities"`
		Protected                      bool   `json:"protected"`
		FollowersCount                 int    `json:"followers_count"`
		FriendsCount                   int    `json:"friends_count"`
		ListedCount                    int    `json:"listed_count"`
		CreatedAt                      string `json:"created_at"`
		FavouritesCount                int    `json:"favourites_count"`
		UtcOffset                      int    `json:"utc_offset"`
		TimeZone                       string `json:"time_zone"`
		GeoEnabled                     bool   `json:"geo_enabled"`
		Verified                       bool   `json:"verified"`
		StatusesCount                  int    `json:"statuses_count"`
		Lang                           string `json:"lang"`
		ContributorsEnabled            bool   `json:"contributors_enabled"`
		IsTranslator                   bool   `json:"is_translator"`
		IsTranslationEnabled           bool   `json:"is_translation_enabled"`
		ProfileBackgroundColor         string `json:"profile_background_color"`
		ProfileBackgroundImageURL      string `json:"profile_background_image_url"`
		ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https"`
		ProfileBackgroundTile          bool   `json:"profile_background_tile"`
		ProfileImageURL                string `json:"profile_image_url"`
		ProfileImageURLHTTPS           string `json:"profile_image_url_https"`
		ProfileBannerURL               string `json:"profile_banner_url"`
		ProfileLinkColor               string `json:"profile_link_color"`
		ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
		ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
		ProfileTextColor               string `json:"profile_text_color"`
		ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
		HasExtendedProfile             bool   `json:"has_extended_profile"`
		DefaultProfile                 bool   `json:"default_profile"`
		DefaultProfileImage            bool   `json:"default_profile_image"`
		Following                      bool   `json:"following"`
		FollowRequestSent              bool   `json:"follow_request_sent"`
		Notifications                  bool   `json:"notifications"`
		TranslatorType                 string `json:"translator_type"`
	} `json:"user"`
	Geo             interface{} `json:"geo"`
	Coordinates     interface{} `json:"coordinates"`
	Place           interface{} `json:"place"`
	Contributors    interface{} `json:"contributors"`
	RetweetedStatus struct {
		CreatedAt string `json:"created_at"`
		ID        int64  `json:"id"`
		IDStr     string `json:"id_str"`
		Text      string `json:"text"`
		Truncated bool   `json:"truncated"`
		Entities  struct {
			Hashtags     []interface{} `json:"hashtags"`
			Symbols      []interface{} `json:"symbols"`
			UserMentions []interface{} `json:"user_mentions"`
			Urls         []struct {
				URL         string `json:"url"`
				ExpandedURL string `json:"expanded_url"`
				DisplayURL  string `json:"display_url"`
				Indices     []int  `json:"indices"`
			} `json:"urls"`
		} `json:"entities"`
		Source               string      `json:"source"`
		InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
		InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
		InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
		InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
		User                 struct {
			ID          int64  `json:"id"`
			IDStr       string `json:"id_str"`
			Name        string `json:"name"`
			ScreenName  string `json:"screen_name"`
			Location    string `json:"location"`
			Description string `json:"description"`
			URL         string `json:"url"`
			Entities    struct {
				URL struct {
					Urls []struct {
						URL         string `json:"url"`
						ExpandedURL string `json:"expanded_url"`
						DisplayURL  string `json:"display_url"`
						Indices     []int  `json:"indices"`
					} `json:"urls"`
				} `json:"url"`
				Description struct {
					Urls []struct {
						URL         string `json:"url"`
						ExpandedURL string `json:"expanded_url"`
						DisplayURL  string `json:"display_url"`
						Indices     []int  `json:"indices"`
					} `json:"urls"`
				} `json:"description"`
			} `json:"entities"`
			Protected                      bool   `json:"protected"`
			FollowersCount                 int    `json:"followers_count"`
			FriendsCount                   int    `json:"friends_count"`
			ListedCount                    int    `json:"listed_count"`
			CreatedAt                      string `json:"created_at"`
			FavouritesCount                int    `json:"favourites_count"`
			UtcOffset                      int    `json:"utc_offset"`
			TimeZone                       string `json:"time_zone"`
			GeoEnabled                     bool   `json:"geo_enabled"`
			Verified                       bool   `json:"verified"`
			StatusesCount                  int    `json:"statuses_count"`
			Lang                           string `json:"lang"`
			ContributorsEnabled            bool   `json:"contributors_enabled"`
			IsTranslator                   bool   `json:"is_translator"`
			IsTranslationEnabled           bool   `json:"is_translation_enabled"`
			ProfileBackgroundColor         string `json:"profile_background_color"`
			ProfileBackgroundImageURL      string `json:"profile_background_image_url"`
			ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https"`
			ProfileBackgroundTile          bool   `json:"profile_background_tile"`
			ProfileImageURL                string `json:"profile_image_url"`
			ProfileImageURLHTTPS           string `json:"profile_image_url_https"`
			ProfileBannerURL               string `json:"profile_banner_url"`
			ProfileLinkColor               string `json:"profile_link_color"`
			ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
			ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
			ProfileTextColor               string `json:"profile_text_color"`
			ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
			HasExtendedProfile             bool   `json:"has_extended_profile"`
			DefaultProfile                 bool   `json:"default_profile"`
			DefaultProfileImage            bool   `json:"default_profile_image"`
			Following                      bool   `json:"following"`
			FollowRequestSent              bool   `json:"follow_request_sent"`
			Notifications                  bool   `json:"notifications"`
			TranslatorType                 string `json:"translator_type"`
		} `json:"user"`
		Geo               interface{} `json:"geo"`
		Coordinates       interface{} `json:"coordinates"`
		Place             interface{} `json:"place"`
		Contributors      interface{} `json:"contributors"`
		IsQuoteStatus     bool        `json:"is_quote_status"`
		RetweetCount      int         `json:"retweet_count"`
		FavoriteCount     int         `json:"favorite_count"`
		Favorited         bool        `json:"favorited"`
		Retweeted         bool        `json:"retweeted"`
		PossiblySensitive bool        `json:"possibly_sensitive"`
		Lang              string      `json:"lang"`
	} `json:"retweeted_status"`
	IsQuoteStatus     bool   `json:"is_quote_status"`
	RetweetCount      int    `json:"retweet_count"`
	FavoriteCount     int    `json:"favorite_count"`
	Favorited         bool   `json:"favorited"`
	Retweeted         bool   `json:"retweeted"`
	PossiblySensitive bool   `json:"possibly_sensitive,omitempty"`
	Lang              string `json:"lang"`
}
type UserDetail struct {
	ContributorsEnabled bool   `json:"contributors_enabled"`
	CreatedAt           string `json:"created_at"`
	DefaultProfile      bool   `json:"default_profile"`
	DefaultProfileImage bool   `json:"default_profile_image"`
	Description         string `json:"description"`
	Entities            struct {
		Description struct {
			Urls []interface{} `json:"urls"`
		} `json:"description"`
		URL struct {
			Urls []struct {
				DisplayURL  string `json:"display_url"`
				ExpandedURL string `json:"expanded_url"`
				Indices     []int  `json:"indices"`
				URL         string `json:"url"`
			} `json:"urls"`
		} `json:"url"`
	} `json:"entities"`
	FavouritesCount                int         `json:"favourites_count"`
	FollowRequestSent              bool        `json:"follow_request_sent"`
	FollowersCount                 int         `json:"followers_count"`
	Following                      bool        `json:"following"`
	FriendsCount                   int         `json:"friends_count"`
	GeoEnabled                     bool        `json:"geo_enabled"`
	ID                             int64       `json:"id"`
	IDStr                          string      `json:"id_str"`
	IsTranslationEnabled           bool        `json:"is_translation_enabled"`
	IsTranslator                   bool        `json:"is_translator"`
	Lang                           string      `json:"lang"`
	ListedCount                    int         `json:"listed_count"`
	Location                       string      `json:"location"`
	Name                           string      `json:"name"`
	Notifications                  bool        `json:"notifications"`
	ProfileBackgroundColor         string      `json:"profile_background_color"`
	ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool        `json:"profile_background_tile"`
	ProfileBannerURL               string      `json:"profile_banner_url"`
	ProfileImageURL                string      `json:"profile_image_url"`
	ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
	ProfileLinkColor               string      `json:"profile_link_color"`
	ProfileLocation                interface{} `json:"profile_location"`
	ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string      `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
	Protected                      bool        `json:"protected"`
	ScreenName                     string      `json:"screen_name"`
	Status                         struct {
		Contributors interface{} `json:"contributors"`
		Coordinates  interface{} `json:"coordinates"`
		CreatedAt    string      `json:"created_at"`
		Entities     struct {
			Hashtags []interface{} `json:"hashtags"`
			Symbols  []interface{} `json:"symbols"`
			Urls     []struct {
				DisplayURL  string `json:"display_url"`
				ExpandedURL string `json:"expanded_url"`
				Indices     []int  `json:"indices"`
				URL         string `json:"url"`
			} `json:"urls"`
			UserMentions []struct {
				ID         int    `json:"id"`
				IDStr      string `json:"id_str"`
				Indices    []int  `json:"indices"`
				Name       string `json:"name"`
				ScreenName string `json:"screen_name"`
			} `json:"user_mentions"`
		} `json:"entities"`
		FavoriteCount        int         `json:"favorite_count"`
		Favorited            bool        `json:"favorited"`
		Geo                  interface{} `json:"geo"`
		ID                   int64       `json:"id"`
		IDStr                string      `json:"id_str"`
		InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
		InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
		InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
		InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
		Lang                 string      `json:"lang"`
		Place                interface{} `json:"place"`
		PossiblySensitive    bool        `json:"possibly_sensitive"`
		RetweetCount         int         `json:"retweet_count"`
		Retweeted            bool        `json:"retweeted"`
		RetweetedStatus      struct {
			Contributors interface{} `json:"contributors"`
			Coordinates  interface{} `json:"coordinates"`
			CreatedAt    string      `json:"created_at"`
			Entities     struct {
				Hashtags []interface{} `json:"hashtags"`
				Symbols  []interface{} `json:"symbols"`
				Urls     []struct {
					DisplayURL  string `json:"display_url"`
					ExpandedURL string `json:"expanded_url"`
					Indices     []int  `json:"indices"`
					URL         string `json:"url"`
				} `json:"urls"`
				UserMentions []interface{} `json:"user_mentions"`
			} `json:"entities"`
			FavoriteCount        int         `json:"favorite_count"`
			Favorited            bool        `json:"favorited"`
			Geo                  interface{} `json:"geo"`
			ID                   int64       `json:"id"`
			IDStr                string      `json:"id_str"`
			InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
			InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
			InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
			InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
			InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
			Lang                 string      `json:"lang"`
			Place                interface{} `json:"place"`
			PossiblySensitive    bool        `json:"possibly_sensitive"`
			RetweetCount         int         `json:"retweet_count"`
			Retweeted            bool        `json:"retweeted"`
			Source               string      `json:"source"`
			Text                 string      `json:"text"`
			Truncated            bool        `json:"truncated"`
		} `json:"retweeted_status"`
		Source    string `json:"source"`
		Text      string `json:"text"`
		Truncated bool   `json:"truncated"`
	} `json:"status"`
	StatusesCount int    `json:"statuses_count"`
	TimeZone      string `json:"time_zone"`
	URL           string `json:"url"`
	UtcOffset     int    `json:"utc_offset"`
	Verified      bool   `json:"verified"`
}

// Twitte Mention Timeline response.
// Refer: https://developer.twitter.com/en/docs/twitter-api/v1/tweets/timelines/api-reference/get-statuses-mentions_timeline

type MentionsTimeline []struct {
	Coordinates interface{} `json:"coordinates"`
	Favorited   bool        `json:"favorited"`
	Truncated   bool        `json:"truncated"`
	CreatedAt   string      `json:"created_at"`
	IDStr       string      `json:"id_str"`
	Entities    struct {
		Urls         []interface{} `json:"urls"`
		Hashtags     []interface{} `json:"hashtags"`
		UserMentions []struct {
			Name       string `json:"name"`
			IDStr      string `json:"id_str"`
			ID         int    `json:"id"`
			Indices    []int  `json:"indices"`
			ScreenName string `json:"screen_name"`
		} `json:"user_mentions"`
	} `json:"entities"`
	InReplyToUserIDStr   string      `json:"in_reply_to_user_id_str"`
	Contributors         interface{} `json:"contributors"`
	Text                 string      `json:"text"`
	RetweetCount         int         `json:"retweet_count"`
	InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
	ID                   int64       `json:"id"`
	Geo                  interface{} `json:"geo"`
	Retweeted            bool        `json:"retweeted"`
	InReplyToUserID      int         `json:"in_reply_to_user_id"`
	Place                interface{} `json:"place"`
	User                 struct {
		ProfileSidebarFillColor   string `json:"profile_sidebar_fill_color"`
		ProfileSidebarBorderColor string `json:"profile_sidebar_border_color"`
		ProfileBackgroundTile     bool   `json:"profile_background_tile"`
		Name                      string `json:"name"`
		ProfileImageURL           string `json:"profile_image_url"`
		CreatedAt                 string `json:"created_at"`
		Location                  string `json:"location"`
		FollowRequestSent         bool   `json:"follow_request_sent"`
		ProfileLinkColor          string `json:"profile_link_color"`
		IsTranslator              bool   `json:"is_translator"`
		IDStr                     string `json:"id_str"`
		Entities                  struct {
			URL struct {
				Urls []struct {
					ExpandedURL interface{} `json:"expanded_url"`
					URL         string      `json:"url"`
					Indices     []int       `json:"indices"`
				} `json:"urls"`
			} `json:"url"`
			Description struct {
				Urls []interface{} `json:"urls"`
			} `json:"description"`
		} `json:"entities"`
		DefaultProfile                 bool        `json:"default_profile"`
		ContributorsEnabled            bool        `json:"contributors_enabled"`
		FavouritesCount                int         `json:"favourites_count"`
		URL                            string      `json:"url"`
		ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
		UtcOffset                      int         `json:"utc_offset"`
		ID                             int         `json:"id"`
		ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
		ListedCount                    int         `json:"listed_count"`
		ProfileTextColor               string      `json:"profile_text_color"`
		Lang                           string      `json:"lang"`
		FollowersCount                 int         `json:"followers_count"`
		Protected                      bool        `json:"protected"`
		Notifications                  interface{} `json:"notifications"`
		ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
		ProfileBackgroundColor         string      `json:"profile_background_color"`
		Verified                       bool        `json:"verified"`
		GeoEnabled                     bool        `json:"geo_enabled"`
		TimeZone                       string      `json:"time_zone"`
		Description                    string      `json:"description"`
		DefaultProfileImage            bool        `json:"default_profile_image"`
		ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
		StatusesCount                  int         `json:"statuses_count"`
		FriendsCount                   int         `json:"friends_count"`
		Following                      interface{} `json:"following"`
		ShowAllInlineMedia             bool        `json:"show_all_inline_media"`
		ScreenName                     string      `json:"screen_name"`
	} `json:"user"`
	InReplyToScreenName string      `json:"in_reply_to_screen_name"`
	Source              string      `json:"source"`
	InReplyToStatusID   interface{} `json:"in_reply_to_status_id"`
}
type FollowerIDs struct {
	Ids               []interface{} `json:"ids"`
	NextCursor        int64         `json:"next_cursor"`
	NextCursorStr     string        `json:"next_cursor_str"`
	PreviousCursor    int           `json:"previous_cursor"`
	PreviousCursorStr string        `json:"previous_cursor_str"`
}
