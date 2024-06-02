package geo

import (
	"GEO_API/proxy/internal/controller/responder"
	"GEO_API/proxy/internal/models"
	"GEO_API/proxy/internal/service/proxyService"
	"encoding/json"
	"net/http"
)

type HandlerGeo struct {
	s proxyService.Geo
	r responder.Responder
}

func New(service proxyService.Geo, responder responder.Responder) HandlerGeo {
	return HandlerGeo{service, responder}
}

// @Summary Search for an address
// @ID search_address
// @Tags geo
// @Accept json
// @Produce json
// @Param request body models.SearchRequest true "Search request"
// @Success 200 {object} models.AddressSearch "get data"
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /address/search [post]
func (h *HandlerGeo) SearchAddressHandler(w http.ResponseWriter, r *http.Request) {
	var searchRequest models.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&searchRequest); err != nil {
		h.r.ErrorBedRequest(w, err)
		return
	}

	address, err := h.s.Address(searchRequest)
	if err != nil {
		h.r.ErrorInternal(w, err)
		return
	}

	h.r.OutputJSON(w, responder.Response{
		Success: true,
		Message: "address get",
		Data:    address,
	})
}

// @Summary Search for an GEO
// @ID GEO_address
// @Tags geo
// @Accept json
// @Produce json
// @Param request body models.GeocodeRequest true "Geocode request"
// @Success 200 {object} models.AddressGeo "get data"
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /address/geocode [post]
func (h *HandlerGeo) GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	var geocodeRequest models.GeocodeRequest
	if err := json.NewDecoder(r.Body).Decode(&geocodeRequest); err != nil {
		h.r.ErrorBedRequest(w, err)
		return
	}

	geocode, err := h.s.Geocode(geocodeRequest)
	if err != nil {
		h.r.ErrorInternal(w, err)
		return
	}
	h.r.OutputJSON(w, responder.Response{
		Success: true,
		Message: "address get",
		Data:    geocode,
	})
}
