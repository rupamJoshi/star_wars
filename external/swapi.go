package external

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/rupam_joshi/star_wars/config"
	"github.com/rupam_joshi/star_wars/graph/model"
)

type Swapi interface {
	GetCharacter(name string) (*model.Character, error)
}

type SwapiImp struct {
	config *config.Config
}

func (s SwapiImp) GetCharacter(name string) (*model.Character, error) {
	resp, err := http.Get(s.config.SWAPIConfig.SWAPIURLPeople)
	if err != nil {
		return nil, fmt.Errorf("failed to call SWAPI: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	characters := []model.Character{}
	if err := json.NewDecoder(resp.Body).Decode(&characters); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// filter by name
	for _, c := range characters {
		if strings.EqualFold(c.Name, name) {

			var resolvedFilms []string
			for _, filmURL := range c.Films {
				title, err := fetchFilmTitle(filmURL)
				if err != nil {
					return nil, err
				}
				resolvedFilms = append(resolvedFilms, title)
			}
			var resolvedVehicles []string
			for _, vURL := range c.Vehicles {
				vName, err := fetchVehicleName(vURL)
				if err != nil {
					return nil, err
				}
				resolvedVehicles = append(resolvedVehicles, vName)
			}
			return &model.Character{
				Name:     c.Name,
				Films:    resolvedFilms,
				Vehicles: resolvedVehicles,
			}, nil
		}
	}
	return nil, err
}

func fetchFilmTitle(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch film: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code for film: %d", resp.StatusCode)
	}

	var film struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&film); err != nil {
		return "", fmt.Errorf("failed to decode film: %w", err)
	}

	return film.Title, nil
}

func fetchVehicleName(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch vehicle: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code for vehicle: %d", resp.StatusCode)
	}

	var vehicle struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&vehicle); err != nil {
		return "", fmt.Errorf("failed to decode vehicle: %w", err)
	}

	return vehicle.Name, nil
}

func NewSWAPI(config *config.Config) Swapi {
	return &SwapiImp{
		config: config,
	}
}
