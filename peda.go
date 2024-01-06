package peda

import (
	"encoding/json"
	"net/http"
)

func PostGeoIntersects(mongoenv, dbname string, r *http.Request) string {
	var longlat LongLat
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)

	err := json.NewDecoder(r.Body).Decode(&longlat)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		response.Status = true
		response.Message = GeoIntersects(mconn, longlat.Longitude, longlat.Latitude)
	}
	return ReturnStruct(response)
}

func PostGeoWithin(mongoenv, dbname string, r *http.Request) string {
	var coordinate GeometryPolygon
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)

	err := json.NewDecoder(r.Body).Decode(&coordinate)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		response.Status = true
		response.Message = GeoWithin(mconn, coordinate.Coordinates)
	}
	return ReturnStruct(response)
}

func PostNear(mongoenv, dbname string, r *http.Request) string {
	var longlat LongLat
	var response Pesan
	response.Status = false
	mconn := SetConnection2dsphereTest(mongoenv, dbname)

	err := json.NewDecoder(r.Body).Decode(&longlat)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		response.Status = true
		response.Message = Near(mconn, longlat.Longitude, longlat.Latitude)
	}
	return ReturnStruct(response)
}
