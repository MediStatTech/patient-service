CREATE TABLE IF NOT EXISTS patient_contact_infos (
    patient_id  UUID NOT NULL REFERENCES patients(patient_id) ON DELETE CASCADE,
    contact_id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    phone       TEXT NOT NULL,
    email       TEXT NOT NULL,
    "primary"   BOOLEAN NOT NULL DEFAULT false,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_patient_contact_infos_patient_id ON patient_contact_infos(patient_id);
