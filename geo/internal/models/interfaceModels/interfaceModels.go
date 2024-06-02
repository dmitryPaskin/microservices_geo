package interfaceModels

import (
	"GeoAPI/geo/internal/models/structModels"
	"github.com/ekomobile/dadata/v2/api/model"
)

type Repository interface {
	GetDataAddress(request *structModels.SearchRequest) ([]*model.Address, error)
	CheckDataAddress(request *structModels.SearchRequest) (bool, error)
	AddDataAddressToDB(request *structModels.SearchRequest, addresses []*model.Address) error

	GetDataGEO(request *structModels.GeocodeRequest) ([]*model.Address, error)
	CheckDataGEO(request *structModels.GeocodeRequest) (bool, error)
	AddDataGEOToDB(request *structModels.GeocodeRequest, geo []*model.Address) error
}

type Service interface {
	SearchService(request structModels.SearchRequest) (*[]*model.Address, error)
	GeocodeService(request structModels.GeocodeRequest) (*[]*model.Address, error)
}
