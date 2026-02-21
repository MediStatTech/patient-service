package repo

import (
	"context"

	"github.com/google/uuid"
)

const ListPatientsByStaffID = `-- name: ListPatientsByStaffID :many
SELECT patient_id, first_name, last_name, gender, dob, staff_id, created_at, updated_at
FROM patients
WHERE staff_id = $1
ORDER BY last_name ASC, first_name ASC
`

func (q *Queries) ListPatientsByStaffID(ctx context.Context, staffID uuid.UUID) ([]Patient, error) {
	rows, err := q.db.QueryContext(ctx, ListPatientsByStaffID, staffID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Patient{}
	for rows.Next() {
		var i Patient
		if err := rows.Scan(
			&i.PatientID,
			&i.FirstName,
			&i.LastName,
			&i.Gender,
			&i.Dob,
			&i.StaffID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
