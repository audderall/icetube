package core

func mapVideo(r *rawVideo) Video {
	video := Video{
		ID:  r.VideoID,
		URL: "https://www.youtube.com/watch?v=" + r.VideoID,
	}
	if len(r.Title.Runs) > 0 {
		video.Title = r.Title.Runs[0].Text
	}
	if r.LengthText != nil {
		video.Duration = r.LengthText.SimpleText
	}
	if r.ViewCountText != nil {
		video.Views = r.ViewCountText.SimpleText
	}
	if r.PublishedTimeText != nil {
		video.PublishedAt = r.PublishedTimeText.SimpleText
	}
	if len(r.Thumbnail.Thumbnails) > 0 {
		video.Thumbnail = r.Thumbnail.Thumbnails[len(r.Thumbnail.Thumbnails)-1].URL
	}
	if len(r.OwnerText.Runs) > 0 {
		run := r.OwnerText.Runs[0]
		video.Channel = Channel{
			ID:   run.NavigationEndpoint.BrowseEndpoint.BrowseID,
			Name: run.Text,
			URL:  "https://www.youtube.com" + run.NavigationEndpoint.BrowseEndpoint.CanonicalBaseURL,
		}
	}
	return video
}

func mapPlaylist(r *rawPlaylist) Playlist {
	playlist := Playlist{
		ID:         r.PlaylistID,
		Title:      r.Title.SimpleText,
		VideoCount: r.VideoCount,
		URL:        "https://www.youtube.com/playlist?list=" + r.PlaylistID,
	}
	if len(r.Thumbnails) > 0 && len(r.Thumbnails[0].Thumbnails) > 0 {
		playlist.Thumbnail = r.Thumbnails[0].Thumbnails[0].URL
	}
	if len(r.ShortBylineText.Runs) > 0 {
		run := r.ShortBylineText.Runs[0]
		playlist.Channel = Channel{
			ID:   run.NavigationEndpoint.BrowseEndpoint.BrowseID,
			Name: run.Text,
			URL:  "https://www.youtube.com" + run.NavigationEndpoint.BrowseEndpoint.CanonicalBaseURL,
		}
	}
	return playlist
}
