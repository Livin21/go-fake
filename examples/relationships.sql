-- Example SQL schema with relationship constraints
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    age INTEGER CHECK (age >= 18 AND age <= 80)
);

CREATE TABLE departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    budget DECIMAL(10,2) CHECK (budget >= 50000 AND budget <= 500000)
);

CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    department_id INTEGER NOT NULL REFERENCES departments(id),
    salary DECIMAL(10,2) CHECK (salary >= 30000 AND salary <= 150000),
    hire_date DATE NOT NULL
);

-- Additional constraints for demonstration
CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    department_id INTEGER REFERENCES departments(id),
    manager_id INTEGER REFERENCES employees(id),
    start_date DATE NOT NULL,
    end_date DATE,
    budget DECIMAL(12,2)
);
