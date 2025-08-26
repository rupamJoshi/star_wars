package service

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rupam_joshi/star_wars/graph/model"
)

type mockSWAPI struct {
	characters map[string]*model.Character
}

func (m *mockSWAPI) GetCharacter(name string) (*model.Character, error) {
	if c, ok := m.characters[name]; ok {
		return c, nil
	}
	return nil, fmt.Errorf("not found")
}

func Test_starWarsService_GetCharacter(t *testing.T) {
	mock := &mockSWAPI{
		characters: map[string]*model.Character{
			"Luke Skywalker": {
				Name:     "Luke Skywalker",
				Films:    []string{"A New Hope"},
				Vehicles: []string{"Snowspeeder"},
			},
		},
	}

	s := &starWarsService{
		swapi: mock,
	}

	tests := []struct {
		name    string
		argName string
		want    *model.Character
		wantErr bool
	}{
		{
			name:    "existing character",
			argName: "Luke Skywalker",
			want: &model.Character{
				Name:     "Luke Skywalker",
				Films:    []string{"A New Hope"},
				Vehicles: []string{"Snowspeeder"},
			},
			wantErr: false,
		},
		{
			name:    "non-existing character",
			argName: "Darth Vader",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetCharacter(tt.argName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCharacter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCharacter() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
