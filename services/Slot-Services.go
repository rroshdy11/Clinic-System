package service

import (
	entity "Clinic_System/entity/Slot"
	"Clinic_System/repository"
	"context"
)

// services needed to control the solt
type SlotService interface {
	GetAll(ctx context.Context) ([]entity.Slot, error)
	GetAllDoctorSlots(ctx context.Context, doctorId string) ([]entity.Slot, error)
	CreateNewSlot(ctx context.Context, slot entity.Slot) (entity.Slot, error)
	DoctorDeleteSlot(ctx context.Context, slotId int) (string, error)
	DoctorUpdateSlot(ctx context.Context, slotId int, slot entity.Slot) (entity.Slot, error)

	PatientCancelSlot(ctx context.Context, slotId int) (string, error)
	PatientReserveSlot(ctx context.Context, slot entity.Slot) (entity.Slot, error)
	GetAllAvailableSlots(ctx context.Context, doctorId string) ([]entity.Slot, error)
	PatientGetAllReservedSlots(ctx context.Context, patientId string) ([]entity.Slot, error)
	UpdateSlotStatus(ctx context.Context, slotIdOld int, slotIdNew int) (string, error)
}

// SlotService implements the SlotService interface
type SlotServiceImpl struct {
	slotRepo repository.SlotRepository
}

// NewSlotService will create new SlotService object representation of SlotService interface
func NewSlotService(slotRepo repository.SlotRepository) SlotService {
	return &SlotServiceImpl{slotRepo: slotRepo}
}

// GetAll will return all the slots
func (ssi *SlotServiceImpl) GetAll(ctx context.Context) ([]entity.Slot, error) {
	slots, err := ssi.slotRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return slots, nil
}

// GetAllDoctorSlots will return all the slots for a specific doctor
func (ssi *SlotServiceImpl) GetAllDoctorSlots(ctx context.Context, doctorId string) ([]entity.Slot, error) {
	slots, err := ssi.slotRepo.GetAllDoctorSlots(doctorId)
	if err != nil {
		return nil, err
	}
	return slots, nil
}

// CreateNewSlot will create a new slot
func (ssi *SlotServiceImpl) CreateNewSlot(ctx context.Context, slot entity.Slot) (entity.Slot, error) {
	slot, err := ssi.slotRepo.NewSlot(slot)
	if err != nil {
		return entity.Slot{}, err
	}
	return slot, nil
}

// DoctorDeleteSlot will delete a slot
func (ssi *SlotServiceImpl) DoctorDeleteSlot(ctx context.Context, slotId int) (string, error) {
	slot, err := ssi.slotRepo.DoctorDeleteSlot(slotId)
	if err != nil {
		return "", err
	}
	return slot, nil
}

// DoctorUpdateSlot will update a slot
func (ssi *SlotServiceImpl) DoctorUpdateSlot(ctx context.Context, slotId int, slot entity.Slot) (entity.Slot, error) {
	slot, err := ssi.slotRepo.DoctorUpdateSlot(slot, slotId)
	if err != nil {
		return entity.Slot{}, err
	}
	return slot, nil
}

// PatientCancelSlot will cancel a slot
func (ssi *SlotServiceImpl) PatientCancelSlot(ctx context.Context, slotId int) (string, error) {
	_, err := ssi.slotRepo.PatientCancelReservation(slotId)
	if err != nil {
		return "", err
	}
	return "Slot Canceled Successfully", nil
}

// PatientReserveSlot will reserve a slot
func (ssi *SlotServiceImpl) PatientReserveSlot(ctx context.Context, slot entity.Slot) (entity.Slot, error) {
	slot, err := ssi.slotRepo.PatientReserveSlot(slot.ID, slot.Patient_ID, slot.Patient_Name)
	if err != nil {
		return entity.Slot{}, err
	}
	return slot, nil
}

// GetAllAvailableSlots will return all the available slots
func (ssi *SlotServiceImpl) GetAllAvailableSlots(ctx context.Context, doctorId string) ([]entity.Slot, error) {
	slots, err := ssi.slotRepo.GetAllAvailableSlots(doctorId)
	if err != nil {
		return nil, err
	}
	return slots, nil
}

// PatientGetAllReservedSlots will return all the reserved slots for a specific patient
func (ssi *SlotServiceImpl) PatientGetAllReservedSlots(ctx context.Context, patientId string) ([]entity.Slot, error) {
	slots, err := ssi.slotRepo.PatientGetAllReservedSlots(patientId)
	if err != nil {
		return nil, err
	}
	return slots, nil
}

// UpdateSlotStatus will update the status of a slot
func (ssi *SlotServiceImpl) UpdateSlotStatus(ctx context.Context, slotIdOld int, slotIdNew int) (string, error) {
	_, err := ssi.slotRepo.UpdateSlotStatus(slotIdOld, slotIdNew)
	if err != nil {
		return "", err
	}
	return "Slot Updated Successfully", nil
}
