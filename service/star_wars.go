package service

import (
	"github.com/rupam_joshi/star_wars/graph/model"
	"github.com/rupam_joshi/star_wars/repo"
)

type StarWarsService interface {
	GetCharacter(name string) (*model.Character, error)
	SaveSearchResult(result model.CharacterResult) (*model.Character, error)
	GetAllSavedCharacters() ([]*model.Character, error)
}

type starWarsService struct {
	repo repo.StarWarRepo
}

func NewStarWarsService(repo repo.StarWarRepo) StarWarsService {
	return &starWarsService{repo: repo}
}

// GetCharacter would normally call SWAPI or another external API.
// For now, let’s just build a Character from the name.
func (s *starWarsService) GetCharacter(name string) (*model.Character, error) {
	// TODO: Call SWAPI resolver here instead of dummy implementation
	return &model.Character{
		Name: name,
	}, nil
}

func (s *starWarsService) SaveSearchResult(result model.CharacterResult) (*model.Character, error) {
	char := &model.Character{
		ID:   result.ID,
		Name: result.Name,
	}

	if err := s.repo.Save(char); err != nil {
		return nil, err
	}

	return char, nil
}

func (s *starWarsService) GetAllSavedCharacters() ([]*model.Character, error) {
	return s.repo.GetAll()
}
