package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	geo "github.com/martinlindhe/google-geolocate"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error get env")
	}
	keyApi := os.Getenv("API_KEYS")
	client := geo.NewGoogleGeo(keyApi)


	// geolocate
	geolocateRes, _ := client.Geolocate()
	fmt.Println(geolocateRes)


	// geo reverse
	p := geo.Point{Lat: 0.7104045, Lng: 123.1463992}
	geoReverseRes, _ := client.ReverseGeocode(&p)
	fmt.Println(geoReverseRes)


	geoCodeRes, _ := client.Geocode("New York City")
	// Output: &{40.7127837 -74.0059413 New York, NY, USA}
	fmt.Println(geoCodeRes)
}
