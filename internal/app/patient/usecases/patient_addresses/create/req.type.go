package create

type Request struct {
	PatientID string
	Line1     string
	City      string
	State     string
}

type Response struct {
	PlaceID string
}
