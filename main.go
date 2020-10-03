package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AstraApp/server/scraper"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

/*
func section(x, y, x int)int {
	if x < 0 {
		if y < 0 {
			if z < 0 {
				//
				return //selection number
			}
			if z >= 0 {
				//
				return
			}
		}
		if y >= 0 {
			if z < 0 {
				//
				return
			}
			if z >= 0 {
				//
				return
			}
		}
	}
	if x >= 0 {
		if y < 0 {
			if z < 0 {
				//
				return
			}
			if z >= 0 {
				//
				return
			}
		}
		if y >= 0 {
			if z < 0 {
				//
				return
			}
			if z >= 0 {
				//
				return
			}
		}
	}
}
*/

func fetchSatelliteName(x, y, cax, cay int) string {

	/*satellites, err := ioutil.ReadFile("./satellites.json")
	if err != nil {
		fmt.Println(err)
	}*/

	return ""
}

func main() {
	fmt.Println("Astra Server")

	s := scraper.New()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		latitude, err := strconv.Atoi(r.URL.Query().Get("latitude"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		longitude, err := strconv.Atoi(r.URL.Query().Get("longitude"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		laCameraAngle, err := strconv.Atoi(r.URL.Query().Get("laCameraAngle"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		loCameraAngle, err := strconv.Atoi(r.URL.Query().Get("loCameraAngle"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		satelliteName := fetchSatelliteName(latitude, longitude, laCameraAngle, loCameraAngle)
		fmt.Println(satelliteName)

		// TODO: Replace "GOES" with satelliteName
		satellite, err := s.Scrape("GOES")
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.MarshalIndent(satellite, "", "\t")
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
	})

	http.ListenAndServe(":8080", r)
}
