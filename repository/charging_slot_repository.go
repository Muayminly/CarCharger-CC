package repository

import (
	"errors"
	"sync"

	"CarCharger-CC/model"
)

type ChargingSlotRepository interface {
	FindAvailableByStationID(stationID string) ([]model.ChargingSlot, error)
	FindByID(slotID string) (*model.ChargingSlot, error)
	UpdateStatus(slotID string, status model.SlotStatus) error
	Save(slot model.ChargingSlot) error
}

type InMemoryChargingSlotRepository struct {
	mu    sync.RWMutex
	slots map[string]model.ChargingSlot
}

func NewInMemoryChargingSlotRepository() *InMemoryChargingSlotRepository {
	return &InMemoryChargingSlotRepository{slots: make(map[string]model.ChargingSlot)}
}

func (r *InMemoryChargingSlotRepository) FindAvailableByStationID(stationID string) ([]model.ChargingSlot, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := []model.ChargingSlot{}
	for _, slot := range r.slots {
		if slot.StationID == stationID && slot.SlotStatus == model.SlotAvailable {
			result = append(result, slot)
		}
	}
	return result, nil
}

func (r *InMemoryChargingSlotRepository) FindByID(slotID string) (*model.ChargingSlot, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	slot, ok := r.slots[slotID]
	if !ok {
		return nil, errors.New("slot not found")
	}
	return &slot, nil
}

func (r *InMemoryChargingSlotRepository) UpdateStatus(slotID string, status model.SlotStatus) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	slot, ok := r.slots[slotID]
	if !ok {
		return errors.New("slot not found")
	}
	slot.SlotStatus = status
	r.slots[slotID] = slot
	return nil
}

func (r *InMemoryChargingSlotRepository) Save(slot model.ChargingSlot) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.slots[slot.SlotID] = slot
	return nil
}
