package entity

type Slot struct {
	ID           int    `json:"id"`
	Date         string `json:"date"`
	Time         string `json:"time"`
	Doctor_ID    string `json:"doctorId"`
	Doctor_Name  string `json:"doctorName"`
	Patient_ID   string `json:"patientId"`
	Patient_Name string `json:"patientName"`
	Status       string `json:"status"`
}
