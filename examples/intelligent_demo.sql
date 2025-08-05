CREATE TABLE users (
    user_id VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email_address VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255),
    home_address VARCHAR(255),
    birth_date VARCHAR(255),
    is_active VARCHAR(255) NOT NULL,
    created_at VARCHAR(255) NOT NULL
);

CREATE TABLE employees (
    employee_id VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    job_title VARCHAR(255) NOT NULL,
    department VARCHAR(255) NOT NULL,
    salary_amount VARCHAR(255) NOT NULL,
    hire_date VARCHAR(255) NOT NULL,
    work_email VARCHAR(255) NOT NULL,
    office_phone VARCHAR(255),
    profile_picture VARCHAR(255),
    has_benefits VARCHAR(255) NOT NULL
);
