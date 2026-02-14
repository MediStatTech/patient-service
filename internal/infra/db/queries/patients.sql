-- name: GetPatient :one
SELECT patient_id, first_name, last_name, gender, dob, staff_id, created_at, updated_at
FROM patients
WHERE patient_id = $1
LIMIT 1;

-- name: ListPatients :many
SELECT patient_id, first_name, last_name, gender, dob, staff_id, created_at, updated_at
FROM patients
ORDER BY last_name ASC, first_name ASC;

-- name: CountPatients :one
SELECT COUNT(*) FROM patients;

-- SQL constants for mutations (used in repository)
-- name: CreatePatient :exec
INSERT INTO patients (
    patient_id,
    first_name,
    last_name,
    gender,
    dob,
    staff_id,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: UpdatePatient :exec
UPDATE patients
SET
    first_name = $2,
    last_name = $3,
    gender = $4,
    dob = $5,
    staff_id = $6,
    updated_at = $7
WHERE patient_id = $1;

-- name: DeletePatient :exec
DELETE FROM patients
WHERE patient_id = $1;
