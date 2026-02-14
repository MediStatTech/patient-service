CREATE TABLE IF NOT EXISTS patient_diseases (
    patient_id  UUID NOT NULL REFERENCES patients(patient_id) ON DELETE CASCADE,
    diseas_id   UUID NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (patient_id, diseas_id)
);

CREATE INDEX idx_patient_diseases_patient_id ON patient_diseases(patient_id);
CREATE INDEX idx_patient_diseases_diseas_id ON patient_diseases(diseas_id);
