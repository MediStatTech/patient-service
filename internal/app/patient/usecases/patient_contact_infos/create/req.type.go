package create

type Request struct {
	PatientID string
	Phone     string
	Email     string
	Primary   bool
}

type Response struct {
	ContactID string
}
