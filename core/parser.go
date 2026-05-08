package core

import (
	"encoding/json"
	"fmt"
)

func Parse(jsonStr string) (SearchResults, error) {
	var root map[string]json.RawMessage
	if err := json.Unmarshal([]byte(jsonStr), &root); err != nil {
		return SearchResults{}, fmt.Errorf("unmarshal root: %w", err)
	}
	contentsRaw, ok := root["contents"]
	if !ok {
		return SearchResults{}, fmt.Errorf("no 'contents' key")
	}

	var wrapper struct {
		TwoColumnSearchResultsRenderer struct {
			PrimaryContents struct {
				SectionListRenderer struct {
					Contents []struct {
						ItemSectionRenderer *struct {
							Contents []json.RawMessage `json:"contents"`
						} `json:"itemSectionRenderer"`
					} `json:"contents"`
				} `json:"sectionListRenderer"`
			} `json:"primaryContents"`
		} `json:"twoColumnSearchResultsRenderer"`
	}
	if err := json.Unmarshal(contentsRaw, &wrapper); err != nil {
		return SearchResults{}, fmt.Errorf("unmarshal wrapper: %w", err)
	}

	var results SearchResults
	sections := wrapper.TwoColumnSearchResultsRenderer.
		PrimaryContents.SectionListRenderer.Contents

	for _, section := range sections {
		if section.ItemSectionRenderer == nil {
			continue
		}
		for _, item := range section.ItemSectionRenderer.Contents {

			var vWrap struct {
				VideoRenderer *rawVideo `json:"videoRenderer"`
			}
			if err := json.Unmarshal(item, &vWrap); err == nil && vWrap.VideoRenderer != nil {
				results.Videos = append(results.Videos, mapVideo(vWrap.VideoRenderer))
				continue
			}

			var pWrap struct {
				PlaylistRenderer *rawPlaylist `json:"playlistRenderer"`
			}
			if err := json.Unmarshal(item, &pWrap); err == nil && pWrap.PlaylistRenderer != nil {
				results.Playlists = append(results.Playlists, mapPlaylist(pWrap.PlaylistRenderer))
			}
		}
	}
	return results, nil
}
