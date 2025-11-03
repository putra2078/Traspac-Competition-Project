CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    user_id REFERENCES INTEGER users(id),
    nip VARCHAR(15) NOT NULL UNIQUE,
    status VARCHAR(15) NOT NULL,
    manager_id REFERENCES INTEGER managers(id),
    contact_id REFERENCES INTEGER contacts(id),
    position_id REFERENCES INTEGER positions(id),
    department_id REFERENCES INTEGER departments(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)