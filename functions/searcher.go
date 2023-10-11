package groupie

import (
	"strconv"
	"strings"
)

func SearchIn(list map[int]Artist, targetStr string) map[int]Artist {
	newList := make(map[int]Artist)

	target := strings.ToLower(targetStr)

	flag := false

	for _, val := range list {

		flag = false
		// search by name

		if strings.Contains(strings.ToLower(val.Name), target) {
			flag = true
		}
		// search by members

		for _, mems := range val.Members {
			if strings.Contains(strings.ToLower(mems), target) {
				flag = true
			}
		}

		// search by creation date
		if strings.Contains(strings.ToLower(strconv.Itoa(val.CreationDate)), target) {
			flag = true
		}

		// first albume date
		if strings.Contains(strings.ToLower(val.FirstAlbum), target) {
			flag = true
		}

		// dates and locations search
		for key, element := range val.LocationsEdited {
			if strings.Contains(strings.ToLower(key), target) {
				flag = true
			}

			for _, vals := range element {
				if strings.Contains(strings.ToLower(vals), target) {
					flag = true
				}
			}
		}

		if flag {
			newList[val.ID] = val
		}

	}

	return newList
}

func GetOptions(dataset map[int]Artist) Option {
	All := map[string]struct{}{}
	Locations := map[string]struct{}{}
	opt := Option{All: All, Locations: Locations}

	for _, val := range dataset {

		opt.All[val.Name] = struct{}{}
		// search by members

		for _, mems := range val.Members {
			opt.All[mems] = struct{}{}
		}
		opt.All[strconv.Itoa(val.CreationDate)] = struct{}{}

		opt.All[val.FirstAlbum] = struct{}{}
		// dates and locations search

		for key, element := range val.LocationsEdited {
			opt.All[key] = struct{}{}
			opt.Locations[key] = struct{}{}

			for _, vals := range element {
				opt.All[vals] = struct{}{}
			}
		}
	}

	return opt
}
