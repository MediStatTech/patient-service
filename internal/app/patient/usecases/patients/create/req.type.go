package create

import "time"

type Request struct {
	FirstName string
	LastName  string
	Gender    string
	Dob       time.Time
	StaffID   *string
}

type Response struct {
	PatientID string
}
