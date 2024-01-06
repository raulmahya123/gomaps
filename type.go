package peda

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBInfo struct {
	DBString string
	DBName   string
}

type GeometryPolygon struct {
	Coordinates [][][]float64 `json:"coordinates" bson:"coordinates"`
	Type        string        `json:"type,omitempty" bson:"type,omitempty"`
}
type GeometryLineString struct {
	Coordinates [][]float64 `json:"coordinates" bson:"coordinates"`
	Type        string      `json:"type" bson:"type"`
}

type GeometryPoint struct {
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
	Type        string    `json:"type" bson:"type"`
}

type GeoJsonLineString struct {
	Type       string             `json:"type" bson:"type"`
	Properties Properties         `json:"properties" bson:"properties"`
	Geometry   GeometryLineString `json:"geometry" bson:"geometry"`
}

type GeoJsonPolygon struct {
	Type       string          `json:"type" bson:"type"`
	Properties Properties      `json:"properties" bson:"properties"`
	Geometry   GeometryPolygon `json:"geometry" bson:"geometry"`
}

type Geometry struct {
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
	Type        string      `json:"type" bson:"type"`
}
type GeoJson struct {
	Type       string     `json:"type" bson:"type"`
	Properties Properties `json:"properties" bson:"properties"`
	Geometry   Geometry   `json:"geometry" bson:"geometry"`
}

type Properties struct {
	Name string `json:"name" bson:"name"`
}

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role,omitempty" bson:"role,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Lokasi struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Properties Name               ` bson:"properties,omitempty"`
	Geometry   Geometry           `bson:"geometry,omitempty"`
	Kategori   string             `bson:"kategori,omitempty"`
}
type Name struct {
	Name string `bson:"name,omitempty"`
}

type LongLat struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Lokasii struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Properties Namee              `bson:"properties,omitempty"`
	Geometry   Geometryy          `bson:"geometry,omitempty"`
	Kategori   string             `bson:"kategori,omitempty"`
}

type Namee struct {
	Name string `bson:"name,omitempty"`
}

type Geometryy struct {
	Type        string      `bson:"type,omitempty"`
	Coordinates [][]float64 `bson:"coordinates,omitempty"`
}

type Pesan struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message" bson:"message"`
}

type GeoBorder struct {
	Type        string        `bson:"type"`
	Coordinates [][][]float64 `bson:"coordinates"`
}

type LocationData struct {
	ID          string    `bson:"_id"`
	Province    string    `bson:"province"`
	District    string    `bson:"district"`
	SubDistrict string    `bson:"sub_district"`
	Village     string    `bson:"village"`
	Border      GeoBorder `bson:"border"`
}
