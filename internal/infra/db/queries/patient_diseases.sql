-- name: GetPatientDiseas :one
SELECT patient_id, diseas_id, created_at, updated_at
FROM patient_diseases
WHERE patient_id = $1 AND diseas_id = $2
LIMIT 1;

-- name: ListPatientDiseases :many
SELECT patient_id, diseas_id, created_at, updated_at
FROM patient_diseases
ORDER BY created_at DESC;

-- name: ListPatientDiseasesByPatientID :many
SELECT patient_id, diseas_id, created_at, updated_at
FROM patient_diseases
WHERE patient_id = $1
ORDER BY created_at DESC;

-- name: ListPatientDiseasesByDiseasID :many
SELECT patient_id, diseas_id, created_at, updated_at
FROM patient_diseases
WHERE diseas_id = $1
ORDER BY created_at DESC;

-- name: CountPatientDiseases :one
SELECT COUNT(*) FROM patient_diseases;

-- name: CountPatientDiseasesByPatientID :one
SELECT COUNT(*)
FROM patient_diseases
WHERE patient_id = $1;

-- SQL constants for mutations (used in repository)
-- name: CreatePatientDiseas :exec
INSERT INTO patient_diseases (
    patient_id,
    diseas_id,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4);

-- name: UpdatePatientDiseas :exec
UPDATE patient_diseases
SET
    updated_at = $3
WHERE patient_id = $1 AND diseas_id = $2;

-- name: DeletePatientDiseas :exec
DELETE FROM patient_diseases
WHERE patient_id = $1 AND diseas_id = $2;
