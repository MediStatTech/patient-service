-- name: GetPatientContactInfo :one
SELECT patient_id, contact_id, phone, email, "primary", created_at, updated_at
FROM patient_contact_infos
WHERE contact_id = $1
LIMIT 1;

-- name: GetPatientContactInfoByPatientID :one
SELECT patient_id, contact_id, phone, email, "primary", created_at, updated_at
FROM patient_contact_infos
WHERE patient_id = $1
LIMIT 1;

-- name: ListPatientContactInfos :many
SELECT patient_id, contact_id, phone, email, "primary", created_at, updated_at
FROM patient_contact_infos
ORDER BY created_at DESC;

-- name: ListPatientContactInfosByPatientID :many
SELECT patient_id, contact_id, phone, email, "primary", created_at, updated_at
FROM patient_contact_infos
WHERE patient_id = $1
ORDER BY created_at DESC;

-- name: CountPatientContactInfos :one
SELECT COUNT(*) FROM patient_contact_infos;

-- SQL constants for mutations (used in repository)
-- name: CreatePatientContactInfo :exec
INSERT INTO patient_contact_infos (
    patient_id,
    contact_id,
    phone,
    email,
    "primary",
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdatePatientContactInfo :exec
UPDATE patient_contact_infos
SET
    phone = $2,
    email = $3,
    "primary" = $4,
    updated_at = $5
WHERE contact_id = $1;

-- name: DeletePatientContactInfo :exec
DELETE FROM patient_contact_infos
WHERE contact_id = $1;
