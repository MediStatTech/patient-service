CREATE TABLE IF NOT EXISTS patients (
    patient_id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name  TEXT NOT NULL,
    last_name   TEXT NOT NULL,
    gender      TEXT NOT NULL,
    dob         DATE NOT NULL,
    staff_id    UUID,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
