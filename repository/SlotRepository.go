package repository

import (
	entity "Clinic_System/entity/Slot"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// the interface for the Slot Repository
type SlotRepository interface {
	GetAll() ([]entity.Slot, error)
	GetAllAvailableSlots(doctorId string) ([]entity.Slot, error)
	GetAllDoctorSlots(doctorId string) ([]entity.Slot, error)
	NewSlot(slot entity.Slot) (entity.Slot, error)
	DoctorUpdateSlot(slot entity.Slot, slotId int) (entity.Slot, error)
	DoctorDeleteSlot(slotId int) (string, error)
	PatientReserveSlot(slotId int, patientId string, patientName string) (entity.Slot, error)
	PatientCancelReservation(slotId int) (string, error)
	PatientGetAllReservedSlots(patientId string) ([]entity.Slot, error)
	UpdateSlotStatus(slotIdOld int, slotIdNew int) (string, error)
}

// the struct that implements the interface
type slotRepository struct {
	db []entity.Slot // slice that will hold the data after loading it from the db
}

// New creates a new Slot repository object (like constructor in java)
func NewSlotRepository() SlotRepository {

	return &slotRepository{}
}

// implement the methods of the interface
// function that gets all Slots stored in the mysql db
func (r slotRepository) GetAll() ([]entity.Slot, error) {

	//call the FillDB function to fill the db with the data in the database
	r.db, _ = FillSlotDB() // return the db
	print(r.db)
	return r.db, nil

}

// function that gets all Slots stored in the mysql db for a specific doctor
func (r slotRepository) GetAllDoctorSlots(doctorId string) ([]entity.Slot, error) {

	//call the FillDB function to fill the db with the data in the database
	r.db, _ = FillSlotDB() // return the db
	var slots []entity.Slot
	for _, slot := range r.db {
		if slot.Doctor_ID == doctorId {
			slots = append(slots, slot)
		}
	}
	return slots, nil

}

func (r slotRepository) NewSlot(slot entity.Slot) (entity.Slot, error) {
	// call the FillDB function to fill the db with the data in the database
	r.db, _ = FillSlotDB()
	// check if the Slot date and time already exists for the same doctor
	for _, slot1 := range r.db {
		if slot1.Date == slot.Date && slot1.Time == slot.Time && slot1.Doctor_ID == slot.Doctor_ID {
			return entity.Slot{}, errors.New("Slot already exists In the same time and date for the same doctor")
		}
	}

	// call the NewSignUp function to write the new user to the database
	addNewSlotToDB(slot)

	return slot, nil
}

func (r slotRepository) DoctorUpdateSlot(slot entity.Slot, slotId int) (entity.Slot, error) {
	// call the FillDB function to fill the db with the data in the database
	r.db, _ = FillSlotDB()

	// check if the Slot date and time already exists for the same doctor
	for _, slot1 := range r.db {
		if slot1.Date == slot.Date && slot1.Time == slot.Time && slot1.Doctor_ID == slot.Doctor_ID {
			return entity.Slot{}, errors.New("Slot already exists In the same time and date for the same doctor")
		}
	}

	// call the NewSignUp function to write the new user to the database
	updateSlotInDB_DateTime(slot, slotId)

	return slot, nil
}

func (r slotRepository) DoctorDeleteSlot(slotId int) (string, error) {
	// call the FillDB function to fill the db with the data in the database
	r.db, _ = FillSlotDB()

	// call the NewSignUp function to write the new user to the database
	deleteSlotFromDB(slotId)

	return "Slot Deleted", nil
}

// GetAllAvailableSlots function to get all the available slots
func (r slotRepository) GetAllAvailableSlots(doctorId string) ([]entity.Slot, error) {
	// call the FillDB function to fill the db with the data in the database
	r.db, _ = FillSlotDB()
	var slots []entity.Slot
	for _, slot := range r.db {
		if slot.Doctor_ID == doctorId && slot.Status == "Available" {
			slots = append(slots, slot)
		}
	}
	return slots, nil
}

// PatientReserveSlot function to reserve a slot
func (r slotRepository) PatientReserveSlot(slotId int, patientId string, patientName string) (entity.Slot, error) {
	// call the FillDB function to fill the db with the data in the database
	r.db, _ = FillSlotDB()
	var slot entity.Slot
	for _, slot1 := range r.db {
		if slot1.ID == slotId {
			slot = slot1
		}
	}
	if slot.Status == "Available" {
		slot.Patient_ID = patientId
		slot.Patient_Name = patientName
		slot.Status = "Reserved"
		updateSlotInDB(slot)
		return slot, nil
	}
	return entity.Slot{}, errors.New("Slot is not available")
}

// PatientCancelReservation function to cancel a reservation
func (r slotRepository) PatientCancelReservation(slotId int) (string, error) {
	// call the FillDB function to fill the db with the data in the database
	r.db, _ = FillSlotDB()
	var slot entity.Slot
	for _, slot1 := range r.db {
		if slot1.ID == slotId {
			slot = slot1
		}
	}
	if slot.Status == "Reserved" {
		slot.Status = "Available"
		updateSlotInDB(slot)
		return "Slot Canceled Successfully", nil
	}
	return "Error", errors.New("Slot is not available")
}

// PatientGetAllReservedSlots function to get all the reserved slots for a specific patient
func (r slotRepository) PatientGetAllReservedSlots(patientId string) ([]entity.Slot, error) {
	// call the FillDB function to fill the db with the data in the database
	r.db, _ = FillSlotDB()
	var slots []entity.Slot
	for _, slot := range r.db {
		if slot.Patient_ID == patientId {
			slots = append(slots, slot)
		}
	}
	return slots, nil
}

// UpdateSlotStatus function to update the status of a slot
func (r slotRepository) UpdateSlotStatus(slotIdOld int, slotIdNew int) (string, error) {
	// call the FillDB function to fill the db with the data in the database
	r.db, _ = FillSlotDB()
	var slotOld entity.Slot
	var slotNew entity.Slot
	for _, slot := range r.db {
		if slot.ID == slotIdOld {
			slotOld = slot
		}
		if slot.ID == slotIdNew {
			slotNew = slot
		}
	}
	if slotOld.Status == "Reserved" && slotNew.Status == "Available" {
		slotOld.Status = "Available"
		slotNew.Patient_ID = slotOld.Patient_ID
		slotNew.Patient_Name = slotOld.Patient_Name
		slotOld.Patient_ID = "nil"
		slotOld.Patient_Name = "nil"
		slotNew.Status = "Reserved"

		updateSlotInDB(slotOld)
		updateSlotInDB(slotNew)
		return "Slot Updated Successfully", nil
	}
	return "Error", errors.New("Slot is not available")
}

// FillDB function to fill the db with the data in the database
func FillSlotDB() ([]entity.Slot, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("root:root@tcp(%s:%s)/clinic_system", dbHost, dbPort)
	dbcon, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	var slots []entity.Slot
	rows, err := dbcon.Query("SELECT id, date, hour, doctor_id, doctor_Name, status, patient_Name, patient_id FROM slots")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var slot entity.Slot
		err = rows.Scan(&slot.ID, &slot.Date, &slot.Time, &slot.Doctor_ID, &slot.Doctor_Name, &slot.Status, &slot.Patient_Name, &slot.Patient_ID)
		if err != nil {
			panic(err.Error())
		}
		slots = append(slots, slot)
	}
	defer dbcon.Close()

	return slots, nil
}

// NewSignUp function to write the new user to the database
func addNewSlotToDB(slot entity.Slot) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("root:root@tcp(%s:%s)/clinic_system", dbHost, dbPort)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO slots(date, hour, doctor_id, doctor_Name, status, patient_Name, patient_id) VALUES (?,?,?,?,?,?,?)", slot.Date, slot.Time, slot.Doctor_ID, slot.Doctor_Name, "Available", "nil", "nil")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

// DeleteSlot function to write the new user to the database
func deleteSlotFromDB(slotId int) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("root:root@tcp(%s:%s)/clinic_system", dbHost, dbPort)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Query("DELETE FROM slots WHERE id=?", slotId)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

// UpdateSlotDateTime function to write change the date and time of a slot in the database by the doctor
func updateSlotInDB_DateTime(slot entity.Slot, slotId int) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("root:root@tcp(%s:%s)/clinic_system", dbHost, dbPort)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Query("UPDATE slots SET date=?, hour=? WHERE id=?", slot.Date, slot.Time, slotId)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

// UpdateSlot function to Change the status ,Patient Name , Patient_Id of a slot in the database by the patient
func updateSlotInDB(slot entity.Slot) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("root:root@tcp(%s:%s)/clinic_system", dbHost, dbPort)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	if slot.Status == "Available" {
		insert, err := db.Query("UPDATE slots SET status=?, patient_id=?, patient_name=? WHERE id=?", slot.Status, "nil", "nil", slot.ID)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
	} else {
		insert, err := db.Query("UPDATE slots SET status=?, patient_id=?, patient_name=? WHERE id=?", slot.Status, slot.Patient_ID, slot.Patient_Name, slot.ID)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
	}
}
