package spotify

import (
	"net/url"
	"strconv"
)

type PersonalizedUserArtists struct {
	Items []struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Followers struct {
			Href  interface{} `json:"href"`
			Total int         `json:"total"`
		} `json:"followers"`
		Genres []string `json:"genres"`
		Href   string   `json:"href"`
		ID     string   `json:"id"`
		Images []struct {
			Height int    `json:"height"`
			URL    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name       string `json:"name"`
		Popularity int    `json:"popularity"`
		Type       string `json:"type"`
		URI        string `json:"uri"`
	} `json:"items"`
	Total    int         `json:"total"`
	Limit    int         `json:"limit"`
	Offset   int         `json:"offset"`
	Href     string      `json:"href"`
	Previous interface{} `json:"previous"`
	Next     string      `json:"next"`
}

type PersonalizedUserTracks struct {
	Items []struct {
		Album struct {
			AlbumType string `json:"album_type"`
			Artists   []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			ExternalUrls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			TotalTracks          int    `json:"total_tracks"`
			Type                 string `json:"type"`
			URI                  string `json:"uri"`
		} `json:"album"`
		Artists []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"artists"`
		AvailableMarkets []string `json:"available_markets"`
		DiscNumber       int      `json:"disc_number"`
		DurationMs       int      `json:"duration_ms"`
		Explicit         bool     `json:"explicit"`
		ExternalIds      struct {
			Isrc string `json:"isrc"`
		} `json:"external_ids"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href        string `json:"href"`
		ID          string `json:"id"`
		IsLocal     bool   `json:"is_local"`
		Name        string `json:"name"`
		Popularity  int    `json:"popularity"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
	} `json:"items"`
	Total    int         `json:"total"`
	Limit    int         `json:"limit"`
	Offset   int         `json:"offset"`
	Href     string      `json:"href"`
	Previous interface{} `json:"previous"`
	Next     string      `json:"next"`
}

func (c *Client) GetUserTopArtistsOpt(opt *Options) (*PersonalizedUserArtists, error) {
	spotifyURL := c.baseURL + "me/top/artists"

	if opt != nil {
		v := url.Values{}
		if opt.Timerange != nil {
			v.Set("time_range", *opt.Timerange)
		}
		if opt.Limit != nil {
			v.Set("limit", strconv.Itoa(*opt.Limit))
		}
		if opt.Offset != nil {
			v.Set("offset", strconv.Itoa(*opt.Offset))
		}
		if params := v.Encode(); params != "" {
			spotifyURL += "?" + params
		}
	}

	var result PersonalizedUserArtists

	err := c.get(spotifyURL, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetUserTopArtists() (*PersonalizedUserArtists, error) {
	return c.GetUserTopArtistsOpt(nil)
}

func (c *Client) GetUserTopTracks2Opt(opt *Options) (*PersonalizedUserTracks, error) {
	spotifyURL := c.baseURL + "me/top/tracks"

	if opt != nil {
		v := url.Values{}
		if opt.Timerange != nil {
			v.Set("time_range", *opt.Timerange+"_term")
		}
		if opt.Limit != nil {
			v.Set("limit", strconv.Itoa(*opt.Limit))
		}
		if opt.Offset != nil {
			v.Set("offset", strconv.Itoa(*opt.Offset))
		}
		if params := v.Encode(); params != "" {
			spotifyURL += "?" + params
		}
	}

	var result PersonalizedUserTracks

	err := c.get(spotifyURL, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetUserTopTracks2() (*PersonalizedUserTracks, error) {
	return c.GetUserTopTracks2Opt(nil)
}
