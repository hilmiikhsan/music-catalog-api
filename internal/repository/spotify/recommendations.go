package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rs/zerolog/log"
)

func (o *outbound) GetRecommendation(ctx context.Context, limit int, trackID string) (*SpotifyRecommendationResponse, error) {
	params := url.Values{}

	params.Set("limit", strconv.Itoa(limit))
	params.Set("market", "ID")
	params.Set("seed_tracks", trackID)

	basePath := "https://api.spotify.com/v1/recommendations"
	urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())

	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create recommendation request for spotify")
		return nil, err
	}

	accessToken, tokenType, err := o.GetTokenDetails()
	if err != nil {
		log.Error().Err(err).Msg("failed to get token details")
		return nil, err
	}

	bearerToken := fmt.Sprintf("%s %s", tokenType, accessToken)

	req.Header.Set("Authorization", bearerToken)

	resp, err := o.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("failed to do recommendation request for spotify")
		return nil, err
	}
	defer resp.Body.Close()

	var response SpotifyRecommendationResponse

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Error().Err(err).Msg("failed to decode response")
		return nil, err
	}

	return &response, nil
}
