DROP DATABASE FASManagementSystem;
CREATE DATABASE FASManagementSystem;
USE FASManagementSystem;

CREATE TABLE applicants (
    id VARCHAR(36) PRIMARY KEY, 
    name VARCHAR(255) NOT NULL,
    employment_status ENUM('Unemployed','Employed') DEFAULT 'Unemployed',
    marital_status ENUM('Single','Married','Widowed','Divorced') DEFAULT 'Single', 
    sex ENUM('male', 'female', 'other') NOT NULL,
    date_of_birth DATE NOT NULL
);
CREATE TABLE household_members (
    id VARCHAR(36) PRIMARY KEY, 
    applicant_id VARCHAR(36),
    name VARCHAR(255) NOT NULL,
    employment_status ENUM('Unemployed','Employed') DEFAULT 'Unemployed',
	sex ENUM('male', 'female', 'other') NOT NULL,
    date_of_birth DATE NOT NULL,
    relation VARCHAR(50) NOT NULL,
    FOREIGN KEY (applicant_id) REFERENCES applicants(id)
);
CREATE TABLE schemes (
id VARCHAR(36) PRIMARY KEY, 
name VARCHAR(255) NOT NULL,
criteria JSON NOT NULL,
description TEXT

);

CREATE TABLE benefits(
scheme_id VARCHAR(36),
id VARCHAR(36) PRIMARY KEY, 
name VARCHAR(255) NOT NULL,
amount INT NOT NULL,
FOREIGN KEY (scheme_id) REFERENCES schemes(id)
);
CREATE TABLE applications(
id VARCHAR(36) PRIMARY KEY, 
status  ENUM('pending', 'approved', 'rejected') DEFAULT 'pending',
  applicant_id VARCHAR(36),
  scheme_id VARCHAR(36),
  FOREIGN KEY (applicant_id) REFERENCES applicants(id),
  FOREIGN KEY (scheme_id) REFERENCES schemes(id)
);
INSERT INTO applicants (id, name, employment_status, sex, date_of_birth) VALUES
('81913b7a-4493-74b2-93f8-e684c4ca935c', 'James', 'Unemployed', 'male', '1990-07-01'),
('81913b80-2c04-7f9d-86a4-497ef68cb3a0', 'Mary', 'Unemployed', 'female', '1984-10-06');

-- Household Members
INSERT INTO household_members (id, applicant_id, name, employment_status, sex, date_of_birth, relation) VALUES
('81913b88-1d4d-7152-a7ce-75796a2e8ecf', '81913b80-2c04-7f9d-86a4-497ef68cb3a0', 'Gwen', 'Unemployed', 'female', '2016-02-01', 'daughter'),
('81913b88-65c6-7255-828f-9c4dd1e5ce79', '81913b80-2c04-7f9d-86a4-497ef68cb3a0', 'Jayden', 'Unemployed', 'male', '2018-03-15', 'son');

-- Schemes
INSERT INTO schemes (id, name, criteria, description) VALUES
('81913689-9a43-7163-8757-81cc254783f3', 'Retrenchment Assistance Scheme', '{"employment_status": "Unemployed"}', 'Assistance for individuals who have been retrenched.'),
('81913689-befc-7ae3-bb37-3079aa7f1be0', 'Retrenchment Assistance Scheme (families)', '{"employment_status": "Unemployed", "has_children": {"school_level": "primary"}}', 'Assistance for retrenched individuals with children in primary school.');
-- Benefits (added from the JSON)
INSERT INTO benefits (scheme_id, id, name, amount) VALUES
('81913689-9a43-7163-8757-81cc254783f3', '81913b8b-9b12-7d2c-a1fa-ea613b802ebc', 'SkillsFuture Credits', 500.00);

