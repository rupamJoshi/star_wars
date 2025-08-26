package service

import (
	"fmt"

	"github.com/rupam_joshi/star_wars/config"
	"github.com/rupam_joshi/star_wars/external"
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
	swapi  external.Swapi
}

func NewStarWarsService(config config.Config, repo repo.StarWarRepo, swapi external.Swapi) StarWarsService {
	return &starWarsService{
		repo:   repo,
		config: config,
		swapi:  swapi,
	}
}

// GetCharacter calls SWAPI
func (s *starWarsService) GetCharacter(name string) (*model.Character, error) {
	lastSearchedCharater, err := s.swapi.GetCharacter(name)
	if err != nil {
		return nil, fmt.Errorf("character %q not found", name)
	}

	return lastSearchedCharater, nil

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
