package groupie

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

var artists = make(map[int]Artist)

func GetJson(link string) ([]byte, error) {
	res, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func InitArtists() error {
	var BaseLinks BaseLink
	links, err := GetJson("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		return err
	}

	err = json.Unmarshal(links, &BaseLinks)
	if err != nil {
		return err
	}

	var Artists []Artist
	content, err := GetJson(BaseLinks.Artists)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &Artists)
	if err != nil {
		return err
	}

	relation, err := GetJson(BaseLinks.Relation)
	if err != nil {
		return err
	}
	var Relations Relation
	err = json.Unmarshal(relation, &Relations)
	if err != nil {
		return err
	}

	for i := range Artists {

		Artists[i].DatesLocations = Relations.Index[i].DatesLocations
		Artists[i].LocationsEdited = ModifyStr(Artists[i].DatesLocations)
		artists[Artists[i].ID] = Artists[i]

	}
	return nil
}

func ModifyStr(m map[string][]string) map[string][]string {
	editedLocations := make(map[string][]string)

	for key := range m {
		newKey := strings.Title(strings.ReplaceAll(strings.ReplaceAll(key, "-", ", "), "_", " "))
		editedLocations[newKey] = m[key]
	}

	return editedLocations
}
