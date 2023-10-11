package groupie

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errors(w, http.StatusNotFound, "Incorrect Request")
		return
	}

	searchList := GetOptions(artists)

	switch r.Method {
	case "GET":
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := t.Execute(w, Response{artists, searchList}); err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
	case "POST":

		t, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}

		target := r.FormValue("target")
		debutDateFrom := r.FormValue("debutDatef")
		debutDateTo := r.FormValue("debutDatet")
		albumDateFrom := r.FormValue("albumDatef")
		albumDateTo := r.FormValue("albumDatet")
		members := r.Form["members"]
		location := r.FormValue("location")

		selected := Selector(artists, target, debutDateFrom, debutDateTo, albumDateFrom, albumDateTo, location, members)

		if len(selected) == 1 {
			url := ""
			for id, _ := range selected {
				url = "/artist/" + strconv.Itoa(id)
			}
			http.Redirect(w, r, url, http.StatusSeeOther)
		}

		if err := t.Execute(w, Response{selected, searchList}); err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}

	default:
		Errors(w, http.StatusMethodNotAllowed, "")
		return
	}
}

func OneArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Errors(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	url := strings.Split(r.URL.Path, "/")

	id, err := strconv.Atoi(url[2])
	if err != nil {
		Errors(w, http.StatusNotFound, "Incorrect Artist Path")
		return
	}
	if len(url) > 3 {
		Errors(w, http.StatusBadRequest, "Bad Request")
		return
	}

	t, err := template.ParseFiles("./templates/artist.html")
	if err != nil {
		Errors(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := r.ParseForm(); err != nil {
		Errors(w, http.StatusInternalServerError, err.Error())
		return
	}

	if id <= 0 || id > len(artists) {
		Errors(w, http.StatusNotFound, "Incorrect Artist Path")
		return
	}

	coords := PrepareLocations(ConcertPlaces(artists[id]))
	encodedCoords, err := json.Marshal(coords)
	if err != nil {
		Errors(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := t.Execute(w, MapAndArtist{artists[id], string(encodedCoords)}); err != nil {
		Errors(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func Errors(w http.ResponseWriter, status int, message string) {
	// w.WriteHeader(status)
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := struct {
		StatusCodeAndText string
		MessageError      string
	}{
		StatusCodeAndText: strconv.Itoa(status) + " " + http.StatusText(status),
		MessageError:      message,
	}
	if err := t.Execute(w, res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
