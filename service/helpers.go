package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
