-- name: CreatePatient :one
INSERT INTO patients (full_name, age, gender, address, phone, status, created_by)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetPatient :one
SELECT * FROM patients WHERE id = $1;

-- name: ListPatients :many
SELECT * FROM patients ORDER BY created_at DESC;

-- name: UpdatePatient :one
UPDATE patients
SET full_name = $2,
    age = $3,
    gender = $4,
    address = $5,
    phone = $6,
    status = $7,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePatient :exec
DELETE FROM patients WHERE id = $1;
