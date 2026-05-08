package core

import "encoding/json"

type Video struct {
	ID          string
	Title       string
	URL         string
	Duration    string
	Views       string
	PublishedAt string
	Thumbnail   string
	Channel     Channel
}

type Channel struct {
	ID   string
	Name string
	URL  string
}

type Playlist struct {
	ID         string
	Title      string
	URL        string
	VideoCount string
	Thumbnail  string
	Channel    Channel
}

type SearchResults struct {
	Videos    []Video
	Playlists []Playlist
}

type ytData struct {
	Contents struct {
		TwoColumnSearchResultsRenderer struct {
			PrimaryContents struct {
				SectionListRenderer struct {
					Contents []struct {
						ItemSectionRenderer struct {
							Contents []json.RawMessage `json:"contents"`
						} `json:"itemSectionRenderer"`
					} `json:"itemSectionRenderer"`
				} `json:"contents"`
			} `json:"sectionListRenderer"`
		} `json:"primaryContents"`
	} `json:"twoColumnSearchResultsRenderer"`
}

type rawVideo struct {
	VideoID string `json:"videoId"`
	Title   struct {
		Runs []struct {
			Text string `json:"text"`
		} `json:"runs"`
	} `json:"title"`
	LengthText *struct {
		SimpleText string `json:"simpleText"`
	} `json:"lengthText"`
	ViewCountText *struct {
		SimpleText string `json:"simpleText"`
	} `json:"viewCountText"`
	PublishedTimeText *struct {
		SimpleText string `json:"simpleText"`
	} `json:"publishedTimeText"`
	Thumbnail struct {
		Thumbnails []struct {
			URL string `json:"url"`
		} `json:"thumbnails"`
	} `json:"thumbnail"`
	OwnerText struct {
		Runs []struct {
			Text               string `json:"text"`
			NavigationEndpoint struct {
				BrowseEndpoint struct {
					BrowseID         string `json:"browseId"`
					CanonicalBaseURL string `json:"canonicalBaseUrl"`
				} `json:"browseEndpoint"`
			} `json:"navigationEndpoint"`
		} `json:"runs"`
	} `json:"ownerText"`
}

type rawPlaylist struct {
	PlaylistID string `json:"playlistId"`
	Title      struct {
		SimpleText string `json:"simpleText"`
	} `json:"title"`
	VideoCount string `json:"videoCount"`
	Thumbnails []struct {
		Thumbnails []struct {
			URL string `json:"url"`
		} `json:"thumbnails"`
	} `json:"thumbnails"`
	ShortBylineText struct {
		Runs []struct {
			Text               string `json:"text"`
			NavigationEndpoint struct {
				BrowseEndpoint struct {
					BrowseID         string `json:"browseId"`
					CanonicalBaseURL string `json:"canonicalBaseUrl"`
				} `json:"browseEndpoint"`
			} `json:"navigationEndpoint"`
		} `json:"runs"`
	} `json:"shortBylineText"`
}
