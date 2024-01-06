package peda

import (
	"fmt"
	"testing"
)

func TestUpdateGetData(t *testing.T) {
	mconn := SetConnection("mongodb+srv://raulgantengbanget:0nGCVlPPoCsXNhqG@cluster0.9ofhjs3.mongodb.net/?retryWrites=true&w=majority", "petapedia")
	datagedung := GeoIntersects(mconn, 107.66417814690686, -6.901637720654366)
	fmt.Println(datagedung)
}

func TestGeoWithin(t *testing.T) {
	mconn := SetConnection("mongodb+srv://raulgantengbanget:0nGCVlPPoCsXNhqG@cluster0.9ofhjs3.mongodb.net/?retryWrites=true&w=majority", "petapediaaa")
	coordinates := [][][]float64{
		{
			{103.62052506248301,
				-1.6105001000148462},
			{103.62061804929925,
				-1.6106710617710007},
			{103.62071435707355,
				-1.6106229269090022},
			{103.62061472834131,
				-1.6104420062116702},
			{103.62052506248301,
				-1.6105001000148462},
		},
	}

	datagedung := GeoWithin(mconn, coordinates)
	fmt.Println(datagedung)
}

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphereTest("mongodb+srv://raulgantengbanget:0nGCVlPPoCsXNhqG@cluster0.9ofhjs3.mongodb.net/?retryWrites=true&w=majority", "petapediaaa")
	datagedung := Near(mconn, 103.61028753823456, -1.6246312350403684)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection2dsphereTestPoint("mongodb+srv://raulgantengbanget:0nGCVlPPoCsXNhqG@cluster0.9ofhjs3.mongodb.net/?retryWrites=true&w=majority", "petapediaaa")
	datagedung := NearSpehere(mconn, 40.78, -73.9667)
	fmt.Println(datagedung)
}

func TestPolygon(t *testing.T) {
	// Set up MongoDB connection for testing
	mconn := SetConnection("mongodb+srv://raulgantengbanget:0nGCVlPPoCsXNhqG@cluster0.9ofhjs3.mongodb.net/?retryWrites=true&w=majority", "petapediaaa")

	// Example coordinates for a polygon
	coordinates := [][][]float64{
		{
			{103.62052506248301, -1.6105001000148462},
			{103.62061804929925, -1.6106710617710007},
			{103.62071435707355, -1.6106229269090022},
			{103.62061472834131, -1.6104420062116702},
			{103.62052506248301, -1.6105001000148462},
		},
	}

	// Call the function being tested
	result := Polygon(mconn, coordinates)

	// Add your assertions based on expected behavior
	expectedResult := ""
	if result != expectedResult {
		t.Errorf("Expected '%s', got '%s'", expectedResult, result)
	}
}

// func TestPolygon(t *testing.T) {
// 	mconn := SetConnection("mongodb+srv://raulgantengbanget:0nGCVlPPoCsXNhqG@cluster0.9ofhjs3.mongodb.net/?retryWrites=true&w=majority", "petapediaaa")
// 	coordinates := [][][]float64{
// 		{
// 			{103.62052506248301,
// 				-1.6105001000148462},
// 			{103.62061804929925,
// 				-1.6106710617710007},
// 			{103.62071435707355,
// 				-1.6106229269090022},
// 			{103.62061472834131,
// 				-1.6104420062116702},
// 			{103.62052506248301,
// 				-1.6105001000148462},
// 		},
// 	}

// 	datagedung := gq.Polygon(mconn, coordinates)
// 	fmt.Println(datagedung)
// }
