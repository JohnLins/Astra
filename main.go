package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/AstraApp/server/scraper"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func section(x, y, z int) string {
	if x < 0 {
		if y < 0 {
			if z < 0 {
				return "3"
			}
			if z >= 0 {
				return "7"
			}
		}
		if y >= 0 {
			if z < 0 {
				return "1"
			}
			if z >= 0 {
				return "5"
			}
		}
	}
	if x >= 0 {
		if y < 0 {
			if z < 0 {
				return "4"
			}
			if z >= 0 {
				return "8"
			}
		}
		if y >= 0 {
			if z < 0 {
				return "2"
			}
			if z >= 0 {
				return "6"
			}
		}
	}
	return ""
}

// Section .
type Section struct {
	Name     string `json:"satellitename"`
	Velocity string `json:"velocity"`
	Image    string `json:"image"`
}

//section ranges from "1" to "8" (The sections in json)
//	Name, Velocity, ImageURL
func fetchSection(section string) (*Section, error) {
	//section := section(x, y, z)
	file, err := os.Open("scraper/satellites.json")
	if err != nil {
		return nil, err
	}

	var satellites map[string][]Section
	if err := json.NewDecoder(file).Decode(&satellites); err != nil {
		return nil, err
	}

	for key, sections := range satellites {
		// Satellite name
		if key != section {
			continue
		}

		// Satellite sections
		for i := 0; i < len(sections); i++ {
			fmt.Println("Name: " + sections[i].Name)
			fmt.Println("Velocity: " + sections[i].Velocity)
			fmt.Println("Image: " + sections[i].Image)
			return &sections[i], nil
		}
	}

	return nil, nil
}

func main() {
	fmt.Println("Astra Server")

	s := scraper.New()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//http://127.0.0.1:8080/?x=5&y=-8&z=3
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		x, err := strconv.Atoi(r.URL.Query().Get("x"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		y, err := strconv.Atoi(r.URL.Query().Get("y"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		z, err := strconv.Atoi(r.URL.Query().Get("z"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		section, err := fetchSection(section(x, y, z))
		if err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		satellite, err := s.Scrape(section.Name)
		if err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.MarshalIndent(struct {
			Name         string
			Velocity     string
			Image        string
			ResourceID   string
			ResourceName string
			Description  string
		}{
			section.Name,
			section.Velocity,
			section.Image,
			satellite.Spase.Observatory.ResourceID,
			satellite.Spase.Observatory.ResourceHeader.ResourceName,
			satellite.Spase.Observatory.ResourceHeader.Description,
		}, "", "\t")
		if err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	})

	http.ListenAndServe(":8080", r)
}
