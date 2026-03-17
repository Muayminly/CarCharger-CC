package service

import (
	"CarCharger-CC/model"
	"CarCharger-CC/repository"
)

type StationService interface {
	SearchStations(location string) ([]model.Station, error)
	GetStationByID(stationID string) (*model.Station, error)
	GetAvailableSlots(stationID string, dateTime string) ([]model.ChargingSlot, error)
	UpdateStationStatus(stationID string, status model.StationStatus) error
}

type stationService struct {
	stationRepo repository.StationRepository
	slotRepo    repository.ChargingSlotRepository
}

func NewStationService(stationRepo repository.StationRepository, slotRepo repository.ChargingSlotRepository) StationService {
	return &stationService{stationRepo: stationRepo, slotRepo: slotRepo}
}

func (s *stationService) SearchStations(location string) ([]model.Station, error) {
	return s.stationRepo.FindByLocation(location)
}

func (s *stationService) GetStationByID(stationID string) (*model.Station, error) {
	return s.stationRepo.FindByID(stationID)
}

func (s *stationService) GetAvailableSlots(stationID string, dateTime string) ([]model.ChargingSlot, error) {
	return s.slotRepo.FindAvailableByStationID(stationID)
}

func (s *stationService) UpdateStationStatus(stationID string, status model.StationStatus) error {
	return s.stationRepo.UpdateStatus(stationID, status)
}
