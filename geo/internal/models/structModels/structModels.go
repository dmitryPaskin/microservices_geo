package structModels

// @name SearchRequest
// @description SearchRequest represents the request body for address search
type SearchRequest struct {
	Query string `json:"query"`
}

// @name GeocodeRequest
// @description GeocodeRequest represents the request body for address geocoding.
type GeocodeRequest struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}
