package groupie

type BaseLink struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type Artist struct {
	ID              int      `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    int      `json:"creationDate"`
	FirstAlbum      string   `json:"firstAlbum"`
	Locations       string   `json:"-"`
	ConcertDates    string   `json:"-"`
	Relations       string   `json:"-"`
	DatesLocations  map[string][]string
	LocationsEdited map[string][]string
}

type Relation struct {
	Index []struct {
		Id             int64               `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type Response struct {
	Main   map[int]Artist
	Search Option
}

type MapAndArtist struct {
	ArtistInfo       Artist
	EncodedLocations string
}

type googleApiResponse struct {
	Results Results `json:"results"`
}

type Results []Geometry

type Geometry struct {
	Geometry Location `json:"geometry"`
}

type Location struct {
	Location Coordinates `json:"location"`
}

type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Option struct {
	All       map[string]struct{}
	Locations map[string]struct{}
}
