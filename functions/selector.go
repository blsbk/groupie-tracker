package groupie

import (
	"strconv"
	"strings"
)

func Selector(list map[int]Artist, target, debutDateFrom, debutDateTo, albumDateFrom, albumDateTo, location string, members []string) map[int]Artist {
	searchRes := SearchIn(list, target)

	filterRes := DebutFilter(searchRes, debutDateFrom, debutDateTo)

	filterRes = AlbumFilter(filterRes, albumDateFrom, albumDateTo)

	filterRes = LocationFilter(filterRes, location)

	if len(members) != 0 {
		filterRes = MembersFilter(filterRes, members)
	}

	return filterRes
}

func AlbumFilter(list map[int]Artist, albumDateFrom, albumDateTo string) map[int]Artist {
	newList := make(map[int]Artist)

	if (albumDateFrom == "1957" && albumDateTo == "2023") || (albumDateFrom == "" && albumDateTo == "") {
		return list
	}
	dateFrom, err := strconv.Atoi(albumDateFrom)
	if err != nil {
		return newList
	}

	dateTo, err := strconv.Atoi(albumDateTo)
	if err != nil {
		return newList
	}

	flag := false

	for _, val := range list {
		flag = false
		date, _ := strconv.Atoi(strings.Split(val.FirstAlbum, "-")[2])

		if date >= dateFrom && date <= dateTo {
			flag = true
		}
		if flag {
			newList[val.ID] = val
		}

	}
	return newList
}

func DebutFilter(list map[int]Artist, debutDateFrom, debutDateTo string) map[int]Artist {
	newList := make(map[int]Artist)

	if (debutDateFrom == "1957" && debutDateTo == "2023") || (debutDateFrom == "" && debutDateTo == "") {
		return list
	}
	dateFrom, err := strconv.Atoi(debutDateFrom)
	if err != nil {
		return newList
	}
	dateTo, err := strconv.Atoi(debutDateTo)
	if err != nil {
		return newList
	}

	flag := false

	for _, val := range list {
		flag = false

		if val.CreationDate >= dateFrom && val.CreationDate <= dateTo {
			flag = true
		}

		if flag {
			newList[val.ID] = val
		}

	}
	return newList
}

func MembersFilter(list map[int]Artist, members []string) map[int]Artist {
	newList := make(map[int]Artist)
	flag := false

	for _, val := range list {
		flag = false
		for _, v := range members {
			if strconv.Itoa(len(val.Members)) == v {
				flag = true
			}
		}

		if flag {
			newList[val.ID] = val
		}

	}
	return newList
}

func LocationFilter(list map[int]Artist, location string) map[int]Artist {
	newList := make(map[int]Artist)

	target := strings.ToLower(location)

	flag := false

	for _, val := range list {

		flag = false

		// dates and locations search
		for key := range val.LocationsEdited {
			if strings.Contains(strings.ToLower(key), target) {
				flag = true
			}
		}

		if flag {
			newList[val.ID] = val
		}

	}

	return newList
}
