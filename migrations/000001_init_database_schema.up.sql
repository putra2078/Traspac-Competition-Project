-- table users --
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email TEXT UNIQUE NOT NULL,
    role VARCHAR(15) CHECK (role IN ('admin', 'manager', 'employee')),
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);

-- table contacts --
CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    photo VARCHAR(500) NOT NULL,
    email TEXT UNIQUE NOT NULL,
    phone_number VARCHAR(20),
    address TEXT,
    gender VARCHAR(15) CHECK (gender IN ('male', 'female', 'other')),
    birth_date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- table departments --
CREATE TABLE departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    slug VARCHAR(20) UNIQUE NOT NULL,
    manager_id INT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- table positions --
CREATE TABLE positions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    department_id INTEGER REFERENCES departments(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- table work_hour --
CREATE TABLE work_hour (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- table managers --
CREATE TABLE managers (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    contact_id INT NOT NULL,
    nip VARCHAR(20) UNIQUE NOT NULL,
    status VARCHAR(20) NOT NULL,
    work_time INTEGER REFERENCES work_hour(id),
    position_id INTEGER REFERENCES positions(id),
    department_id INTEGER REFERENCES departments(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_manager_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,

    CONSTRAINT fk_manager_contact
    FOREIGN KEY (contact_id)
    REFERENCES contacts(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

-- add constraint department manager_id
ALTER TABLE departments
ADD CONSTRAINT fk_department_manager
FOREIGN KEY (manager_id)
REFERENCES managers(id)
ON UPDATE CASCADE
ON DELETE SET NULL;

-- table employee --
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    contact_id INT NOT NULL,
    nip VARCHAR(20) UNIQUE NOT NULL,
    status VARCHAR(20) NOT NULL,
    work_time INTEGER REFERENCES work_hour(id),
    manager_id INT,
    position_id INTEGER REFERENCES positions(id),
    department_id INTEGER REFERENCES departments(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_employee_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,

    CONSTRAINT fk_employee_contact
    FOREIGN KEY (contact_id)
    REFERENCES contacts(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,

    CONSTRAINT fk_employee_manager
    FOREIGN KEY (manager_id)
    REFERENCES managers(id)
    ON UPDATE CASCADE
    ON DELETE SET NULL
);

-- table admin --
CREATE TABLE admins (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    contact_id INT NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_admin_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,

    CONSTRAINT fk_admin_contact
    FOREIGN KEY (contact_id)
    REFERENCES contacts(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

-- table presence --
CREATE TABLE presences (
    id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL,
    date DATE NOT NULL,
    check_in_hour TIME NOT NULL,
    lat_check_in DECIMAL NOT NULL,
    long_check_in DECIMAL NOT NULL,
    check_in_status VARCHAR(20) NOT NULL,
    check_out_hour TIME  NULL,
    lat_check_out DECIMAL  NULL,
    long_check_out DECIMAL  NULL,
    check_out_status VARCHAR(20) NULL,
    CONSTRAINT fk_presences_employee
    FOREIGN KEY (employee_id)
    REFERENCES employees(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

-- table attendance_categories --
CREATE TABLE attendance_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    code VARCHAR(15) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- table attendances --
CREATE TABLE attendances (
    id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL,
    category_id INT NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_attendances_employee
    FOREIGN KEY (employee_id)
    REFERENCES employees(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,

    CONSTRAINT fk_attendance_category
    FOREIGN KEY (category_id)
    REFERENCES attendance_categories(id)
    ON UPDATE CASCADE
    ON DELETE SET NULL
);

-- table absences --
CREATE TABLE absences (
    id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL,
    category_id INT NOT NULL,
    date DATE NOT NULL,
    description VARCHAR(255) NOT NULL,
    status VARCHAR(15) CHECK (status IN ('pending', 'appproved', 'rejected')),
    reject_reason VARCHAR(255) NULL,
    CONSTRAINT fk_absences_employee
    FOREIGN KEY (employee_id)
    REFERENCES employees(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,

    CONSTRAINT fk_absence_category
    FOREIGN KEY (category_id)
    REFERENCES attendance_categories(id)
    ON UPDATE CASCADE
    ON DELETE SET NULL
);

-- table overtimes --
CREATE TABLE overtimes (
    id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL,
    department_id INT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    description VARCHAR(255) NOT NULL,
    status VARCHAR(15) CHECK (status IN ('pending', 'approved', 'rejected')),
    reject_reason VARCHAR(255) NULL,

    CONSTRAINT fk_overtime_employee
    FOREIGN KEY (employee_id)
    REFERENCES employees(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,

    CONSTRAINT fk_overtime_department
    FOREIGN KEY (department_id)
    REFERENCES departments(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

-- table settings --
CREATE TABLE settings (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    language VARCHAR(15) CHECK (language IN ('en_us', 'indo')),
    notification VARCHAR(15) CHECK (notification IN ('allowed', 'not_allowed')),

    CONSTRAINT fk_users_settings
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

-- table documents --
CREATE TABLE documents (
    id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    path VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_employee_id
    FOREIGN KEY (employee_id)
    REFERENCES employees(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);


-- table workspaces --
CREATE TABLE workspaces (
    id SERIAL PRIMARY KEY,
    pass_code VARCHAR(255) NULL,
    created_by INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    privacy VARCHAR(15) CHECK (privacy IN ('private', 'public')),
    join_link VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_workspace_creator
    FOREIGN KEY (created_by)
    REFERENCES users(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

-- table task_tab --
CREATE TABLE task_tab (
    id SERIAL PRIMARY KEY,
    workspace_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    position INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_task_workspace
    FOREIGN KEY (workspace_id)
    REFERENCES workspaces(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

-- table task_card --
CREATE TABLE task_card (
    id SERIAL PRIMARY KEY,
    task_tab_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    content VARCHAR(255) NULL,
    comment VARCHAR(255) NULL,
    date DATE NOT NULL,
    status BOOLEAN NOT NULL,
    employee_id INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_task_tab
    FOREIGN KEY (task_tab_id)
    REFERENCES task_tab(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,

    CONSTRAINT fk_employee_task
    FOREIGN KEY (employee_id)
    REFERENCES employees(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

-- table servers --
CREATE TABLE servers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_by INT NOT NULL,
    privacy VARCHAR(15) CHECK (privacy IN ('public', 'private')),
    pass_code VARCHAR(255) NULL,
    link_join VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_server_creator
    FOREIGN KEY (created_by)
    REFERENCES users(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

-- table room_chat --
CREATE TABLE rooms_chat (
    id SERIAL PRIMARY KEY,
    server_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_by INT NOT NULL,
    pass_code VARCHAR(255) NULL,
    link_join VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_rooms_creator
    FOREIGN KEY (created_by)
    REFERENCES users(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,

    CONSTRAINT fk_room_server
    FOREIGN KEY (server_id)
    REFERENCES servers(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

-- table room_messages --
CREATE TABLE room_messages (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE RESTRICT,
    room_id INT NOT NULL,
    message_text VARCHAR(500) NOT NULL,
    message_content VARCHAR(500) NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_messages_room
    FOREIGN KEY(room_id)
    REFERENCES rooms_chat(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

-- table direct_message --
CREATE TABLE direct_message (
    id SERIAL PRIMARY KEY,
    user_receiver INTEGER REFERENCES users(id) ON DELETE RESTRICT,
    user_sender INTEGER REFERENCES users(id) ON DELETE RESTRICT,
    message_text VARCHAR(500) NOT NULL,
    message_content VARCHAR(500) NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);