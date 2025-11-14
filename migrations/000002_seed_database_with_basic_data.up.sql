-- seed table departments --
INSERT INTO departments (name, slug)
VALUES
('Human Resource', 'hr'),
('Engineering', 'engineer'),
('Finance', 'finance')
ON CONFLICT (name) DO NOTHING;

-- seed table positions --
INSERT INTO positions (name, department_id)
VALUES 
('Backend Developer', 1),
('Frontend Developer', 1),
('HR Specialist', 1),
('Finance Analyst', 1)
ON CONFLICT (name) DO NOTHING;

-- seed table attendance_categories --
INSERT INTO attendance_categories (name, code)
VALUES
('Presence', 'P'),
('Absence', 'A'),
('WFH', 'WFH'),
('Sick Leave', 'SL'),
('Leave', 'L'),
('Vacation', 'V'),
('No Information', 'N')
ON CONFLICT (name) DO NOTHING;

-- seed table work_hour --
INSERT INTO work_hour (name, start_time, end_time)
VALUES
    ('Shift Pagi', '08:00:00+07:00', '16:00:00+07:00'),
    ('Shift Siang', '12:00:00+07:00', '20:00:00+07:00'),
    ('Shift Malam', '20:00:00+07:00', '04:00:00+07:00'),
    ('Shift Kantoran', '09:00:00+07:00', '17:00:00+07:00')
ON CONFLICT (name) DO NOTHING;
