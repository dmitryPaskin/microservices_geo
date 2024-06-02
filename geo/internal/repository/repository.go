package repository

import (
	"GeoAPI/geo/internal/models/interfaceModels"
	"GeoAPI/geo/internal/models/structModels"
	"database/sql"
	"encoding/json"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/ekomobile/dadata/v2/api/model"
)

type repository struct {
	db         *sql.DB
	sqlBuilder sq.StatementBuilderType
}

func New(database *sql.DB) interfaceModels.Repository {
	return &repository{
		db:         database,
		sqlBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *repository) GetDataAddress(request *structModels.SearchRequest) ([]*model.Address, error) {
	var result []*model.Address
	var resultString string
	query := r.sqlBuilder.Select("data").
		From("address_data").Where(sq.Eq{"address": request.Query})

	row := query.RunWith(r.db).QueryRow()

	if err := row.Scan(&resultString); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(resultString), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) CheckDataAddress(request *structModels.SearchRequest) (bool, error) {
	query := r.sqlBuilder.Select("COUNT(*)").
		From("address_data").Where(sq.Eq{"address": request.Query})

	row := query.RunWith(r.db).QueryRow()
	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (r *repository) AddDataAddressToDB(request *structModels.SearchRequest, addresses []*model.Address) error {
	data, err := json.Marshal(addresses)
	if err != nil {
		return err
	}

	query := r.sqlBuilder.Insert("address_data").
		Columns("address", "data").
		Values(request.Query, string(data))

	if _, err := query.RunWith(r.db).Exec(); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetDataGEO(request *structModels.GeocodeRequest) ([]*model.Address, error) {
	var result []*model.Address
	var resultString string
	geoRequest := fmt.Sprintf("%s, %s", request.Lon, request.Lat)
	query := r.sqlBuilder.Select("data").
		From("geo_data").Where(sq.Eq{"geo": geoRequest})

	row := query.RunWith(r.db).QueryRow()

	if err := row.Scan(&resultString); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(resultString), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) CheckDataGEO(request *structModels.GeocodeRequest) (bool, error) {
	geoRequest := fmt.Sprintf("%s, %s", request.Lon, request.Lat)
	query := r.sqlBuilder.Select("COUNT(*)").
		From("geo_data").Where(sq.Eq{"geo": geoRequest})

	row := query.RunWith(r.db).QueryRow()
	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (r *repository) AddDataGEOToDB(request *structModels.GeocodeRequest, geo []*model.Address) error {
	data, err := json.Marshal(geo)
	if err != nil {
		return err
	}

	geoRequest := fmt.Sprintf("%s, %s", request.Lon, request.Lat)
	query := r.sqlBuilder.Insert("geo_data").
		Columns("geo", "data").
		Values(geoRequest, data)

	if _, err := query.RunWith(r.db).Exec(); err != nil {
		return err
	}

	return nil
}
