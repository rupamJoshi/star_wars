package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/rupam_joshi/star_wars/config"
	"github.com/rupam_joshi/star_wars/graph/model"
	"github.com/rupam_joshi/star_wars/repo"
)

var lastSearchedCharater model.Character

type StarWarsService interface {
	GetCharacter(name string) (*model.Character, error)
	SaveSearchResult() (*model.FavoriteCharacter, error)
	GetAllSavedCharacters() ([]*model.FavoriteCharacter, error)
}

type starWarsService struct {
	config config.Config
	repo   repo.StarWarRepo
}

func NewStarWarsService(config config.Config, repo repo.StarWarRepo) StarWarsService {
	return &starWarsService{
		repo:   repo,
		config: config,
	}
}

// GetCharacter calls SWAPI
func (s *starWarsService) GetCharacter(name string) (*model.Character, error) {
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
			lastSearchedCharater = model.Character{
				Name:     c.Name,
				Films:    resolvedFilms,
				Vehicles: resolvedVehicles,
			}
			return &lastSearchedCharater, nil
		}
	}
	return nil, fmt.Errorf("character %q not found", name)

}

func (s *starWarsService) SaveSearchResult() (*model.FavoriteCharacter, error) {
	char := &model.FavoriteCharacter{
		Name:     lastSearchedCharater.Name,
		Films:    lastSearchedCharater.Films,
		Vehicles: lastSearchedCharater.Vehicles,
	}

	if err := s.repo.Save(char); err != nil {
		return nil, err
	}

	return char, nil
}

func (s *starWarsService) GetAllSavedCharacters() ([]*model.FavoriteCharacter, error) {
	return s.repo.GetAll()
}
