CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    nip VARCHAR(15) NOT NULL UNIQUE,
    status VARCHAR(15) NOT NULL,
    manager_id INTEGER REFERENCES managers(id),
    contact_id INTEGER REFERENCES contacts(id),
    position_id INTEGER REFERENCES positions(id),
    department_id INTEGER REFERENCES departments(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)