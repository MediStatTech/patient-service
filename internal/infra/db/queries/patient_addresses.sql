-- name: GetPatientAddress :one
SELECT patient_id, place_id, line_1, city, state, created_at, updated_at
FROM patient_addresses
WHERE place_id = $1
LIMIT 1;

-- name: GetPatientAddressByPatientID :one
SELECT patient_id, place_id, line_1, city, state, created_at, updated_at
FROM patient_addresses
WHERE patient_id = $1
LIMIT 1;

-- name: ListPatientAddresses :many
SELECT patient_id, place_id, line_1, city, state, created_at, updated_at
FROM patient_addresses
ORDER BY created_at DESC;

-- name: ListPatientAddressesByPatientID :many
SELECT patient_id, place_id, line_1, city, state, created_at, updated_at
FROM patient_addresses
WHERE patient_id = $1
ORDER BY created_at DESC;

-- name: CountPatientAddresses :one
SELECT COUNT(*) FROM patient_addresses;

-- SQL constants for mutations (used in repository)
-- name: CreatePatientAddress :exec
INSERT INTO patient_addresses (
    patient_id,
    place_id,
    line_1,
    city,
    state,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdatePatientAddress :exec
UPDATE patient_addresses
SET
    line_1 = $2,
    city = $3,
    state = $4,
    updated_at = $5
WHERE place_id = $1;

-- name: DeletePatientAddress :exec
DELETE FROM patient_addresses
WHERE place_id = $1;
