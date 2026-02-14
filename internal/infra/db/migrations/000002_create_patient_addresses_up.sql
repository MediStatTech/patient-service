CREATE TABLE IF NOT EXISTS patient_addresses (
    patient_id  UUID NOT NULL REFERENCES patients(patient_id) ON DELETE CASCADE,
    place_id    UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    line_1      TEXT NOT NULL,
    city        TEXT NOT NULL,
    state       TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_patient_addresses_patient_id ON patient_addresses(patient_id);
