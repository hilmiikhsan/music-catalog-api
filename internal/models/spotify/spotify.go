package spotify

type SearchRespose struct {
	Items  []SpotifyTrackObject `json:"items"`
	Total  int                  `json:"total"`
	Limit  int                  `json:"limit"`
	Offset int                  `json:"offset"`
}

type SpotifyTrackObject struct {
	AlbumType        string   `json:"album_type"`
	AlbumTotalTracks int      `json:"album_total_tracks"`
	AlbumImagesUrl   []string `json:"album_images_url"`
	AlbumName        string   `json:"album_name"`
	ArtistsName      []string `json:"artists_name"`
	Explicit         bool     `json:"explicit"`
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	IsLike           *bool    `json:"is_like"`
}

type RecommendationResponse struct {
	Items []SpotifyTrackObject `json:"items"`
}
