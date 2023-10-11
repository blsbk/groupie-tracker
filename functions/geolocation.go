package groupie

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const googleApiUri = "https://maps.googleapis.com/maps/api/geocode/json?key=AIzaSyDZ65aafPZP_EQ6292Zb-5_GvZsKG_RE68&address="

func GetCityCoordinates(address string) Coordinates {
	// fmt.Println("Fetching langitude and longitude of the city ...")
	resp, err := http.Get(googleApiUri + address)
	if err != nil {
		log.Fatal("Fetching google api uri data error: ", err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("Reading google api data error: ", err)
	}

	var data googleApiResponse
	json.Unmarshal(bytes, &data)

	return Coordinates{data.Results[0].Geometry.Location.Latitude, data.Results[0].Geometry.Location.Longitude}
}

func PrepareLocations(strs []string) []Coordinates {
	var coords []Coordinates
	for _, val := range strs {
		cor := GetCityCoordinates(val)
		coords = append(coords, cor)
	}
	return coords
}

func ConcertPlaces(data Artist) []string {
	addresses := data.DatesLocations

	var adress []string
	for key := range addresses {
		adress = append(adress, key)
	}

	return adress
}
