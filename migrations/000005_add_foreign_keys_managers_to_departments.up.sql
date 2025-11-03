-- Menambahkan foreign key relation department -> manager
ALTER TABLE departments
ADD CONSTRAINT fk_department_manager
FOREIGN KEY (head_manager)
REFERENCES managers(id)
ON DELETE SET NULL;

-- Menambahkan foreign key relation manager -> department
ALTER TABLE managers
ADD CONSTRAINT fk_manager_department
FOREIGN KEY (department_id)
REFERENCES departments(id)
ON DELETE SET NULL;