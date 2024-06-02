package proxyService

import (
	"GeoAPI/geo/internal/models/interfaceModels"
	"GeoAPI/geo/internal/models/structModels"
	"fmt"
	"github.com/ekomobile/dadata/v2/api/model"
	"github.com/go-redis/redis"
)

type proxyService struct {
	service    interfaceModels.Service
	repository interfaceModels.Repository
	redis      *ClientRedis
}

func NewProxyService(realObject interfaceModels.Service, repository interfaceModels.Repository, address, password string) (interfaceModels.Service, error) {
	redis, err := NewRedis(address, password)
	if err != nil {
		return nil, err
	}
	return &proxyService{
		service:    realObject,
		repository: repository,
		redis:      redis,
	}, nil
}

func (p *proxyService) SearchService(request structModels.SearchRequest) (*[]*model.Address, error) {
	CacheRedis, err := p.redis.CheckCacheInRedis(request.Query)
	if err == redis.Nil {
		isCacheBD, err := p.repository.CheckDataAddress(&request)
		if err != nil {
			return nil, err
		}

		if isCacheBD {
			result, err := p.repository.GetDataAddress(&request)
			if err != nil {
				return nil, err
			}

			if err := p.redis.SaveCacheInRedis(request.Query, result); err != nil {
				return nil, err
			}
			return &result, err
		} else {
			result, err := p.service.SearchService(request)
			if err != nil {
				return nil, err
			}

			if err := p.redis.SaveCacheInRedis(request.Query, result); err != nil {
				return nil, err
			}

			if err := p.repository.AddDataAddressToDB(&request, *result); err != nil {
				return nil, err
			}
			return result, nil
		}
	} else if err != nil {
		return nil, err
	} else {
		return CacheRedis, nil
	}
}

func (p *proxyService) GeocodeService(request structModels.GeocodeRequest) (*[]*model.Address, error) {
	geoRequest := fmt.Sprintf("%s, %s", request.Lon, request.Lat)

	CacheRedis, err := p.redis.CheckCacheInRedis(geoRequest)
	if err == redis.Nil {
		isCacheBD, err := p.repository.CheckDataGEO(&request)
		if err != nil {
			return nil, err
		}

		if isCacheBD {
			result, err := p.repository.GetDataGEO(&request)
			if err != nil {
				return nil, err
			}

			if err := p.redis.SaveCacheInRedis(geoRequest, result); err != nil {
				return nil, err
			}
			return &result, err
		} else {
			result, err := p.service.GeocodeService(request)
			if err != nil {
				return nil, err
			}

			if err := p.redis.SaveCacheInRedis(geoRequest, result); err != nil {
				return nil, err
			}

			if err := p.repository.AddDataGEOToDB(&request, *result); err != nil {
				return nil, err
			}
			return result, nil
		}
	} else if err != nil {
		return nil, err
	} else {
		return CacheRedis, nil
	}
}
