package repository

import (
	"errors"
	"strings"
	"sync"

	"CarCharger-CC/model"
)

type StationRepository interface {
	FindByLocation(location string) ([]model.Station, error)
	FindByID(stationID string) (*model.Station, error)
	UpdateStatus(stationID string, status model.StationStatus) error
	Save(station model.Station) error
}

type InMemoryStationRepository struct {
	mu       sync.RWMutex
	stations map[string]model.Station
}

func NewInMemoryStationRepository() *InMemoryStationRepository {
	return &InMemoryStationRepository{stations: make(map[string]model.Station)}
}

func (r *InMemoryStationRepository) FindByLocation(location string) ([]model.Station, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []model.Station
	for _, station := range r.stations {
		if location == "" || strings.Contains(strings.ToLower(station.Location), strings.ToLower(location)) {
			result = append(result, station)
		}
	}
	return result, nil
}

func (r *InMemoryStationRepository) FindByID(stationID string) (*model.Station, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	station, ok := r.stations[stationID]
	if !ok {
		return nil, errors.New("station not found")
	}
	return &station, nil
}

func (r *InMemoryStationRepository) UpdateStatus(stationID string, status model.StationStatus) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	station, ok := r.stations[stationID]
	if !ok {
		return errors.New("station not found")
	}
	station.StationStatus = status
	r.stations[stationID] = station
	return nil
}

func (r *InMemoryStationRepository) Save(station model.Station) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.stations[station.StationID] = station
	return nil
}
