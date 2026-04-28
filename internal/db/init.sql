CREATE TABLE admins (
    admin_id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(100) NULL,
    is_authorized TINYINT(1) DEFAULT 0 NULL,
    is_email_verified TINYINT(1) DEFAULT 0 NULL
);

CREATE TABLE admin_tokens (
    token_id INT AUTO_INCREMENT PRIMARY KEY,
    admin_id INT NOT NULL,
    verification_token VARCHAR(64) NOT NULL,
    token_expires_at DATETIME NOT NULL,
    CONSTRAINT admin_tokens_ibfk_1 FOREIGN KEY (admin_id) REFERENCES admins (admin_id) ON DELETE CASCADE
);

CREATE INDEX admin_id ON admin_tokens (admin_id);

CREATE TABLE faculty_members (
    faculty_member_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    faculty_role VARCHAR(100) NULL,
    email VARCHAR(255) NULL,
    phone VARCHAR(50) NULL,
    extension VARCHAR(50) NULL,
    linkedin_url VARCHAR(512) NULL,
    image_path VARCHAR(100) NULL
);

CREATE TABLE funding_sources (
    funding_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    website_url VARCHAR(512) NULL,
    image_path VARCHAR(512) NULL,
    display_order INT DEFAULT 0 NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NULL
);

CREATE TABLE landslides (
    landslide_id INT AUTO_INCREMENT PRIMARY KEY,
    landslide_date DATETIME NULL,
    latitude DECIMAL(9, 6) NOT NULL,
    longitude DECIMAL(9, 6) NOT NULL,
    image_path VARCHAR(512) NULL
);

CREATE TABLE office_info (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NULL,
    phone VARCHAR(100) NULL,
    phone_ext VARCHAR(100) NULL,
    office_location VARCHAR(512) NULL,
    facebook_url VARCHAR(512) NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE projects (
    project_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description TEXT NULL,
    start_year SMALLINT UNSIGNED NULL,
    end_year SMALLINT UNSIGNED NULL,
    project_status ENUM('planned', 'active', 'paused', 'completed', 'archived') DEFAULT 'planned' NULL,
    image_path VARCHAR(512) NULL
);

CREATE TABLE publications (
    publication_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description TEXT NULL,
    publication_url VARCHAR(512) NULL,
    image_path VARCHAR(512) NULL,
    published_date DATE NULL
);

CREATE TABLE reports (
    report_id INT AUTO_INCREMENT PRIMARY KEY,
    landslide_id INT NULL,
    reported_at DATETIME DEFAULT CURRENT_TIMESTAMP NULL,
    latitude DECIMAL(9, 6) NOT NULL,
    longitude DECIMAL(9, 6) NOT NULL,
    city VARCHAR(100) NULL,
    physical_address VARCHAR(512) NULL,
    reporter_name VARCHAR(100) NULL,
    reporter_phone VARCHAR(30) NULL,
    reporter_email VARCHAR(255) NULL,
    description TEXT NULL,
    image_path VARCHAR(512) NULL,
    is_validated TINYINT(1) DEFAULT 0 NULL,
    CONSTRAINT reports_ibfk_1 FOREIGN KEY (landslide_id) REFERENCES landslides (landslide_id) ON DELETE SET NULL
);

CREATE INDEX landslide_id ON reports (landslide_id);

CREATE TABLE stations (
    station_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    land_unit VARCHAR(50) NULL,
    geological_unit VARCHAR(50) NULL,
    susceptibility VARCHAR(50) NULL,
    depth VARCHAR(50) NULL,
    landslide_forecast DECIMAL(6, 2) NULL,
    image_path VARCHAR(100) NULL,
    latitude DECIMAL(9, 6) NOT NULL,
    longitude DECIMAL(9, 6) NOT NULL,
    elevation INT NULL,
    slope DECIMAL(10, 2) NULL,
    is_available TINYINT(1) DEFAULT 1 NULL,
    collaborator VARCHAR(100) NULL,
    station_installation_date DATETIME NULL
);

CREATE TABLE station_readings (
    reading_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    station_id INT NOT NULL,
    recorded_at DATETIME NOT NULL,
    image_path VARCHAR(100) NULL,
    precipitation DECIMAL(6, 2) DEFAULT 0.00 NULL,
    wc1 DECIMAL(9, 6) NULL,
    wc2 DECIMAL(9, 6) NULL,
    wc3 DECIMAL(9, 6) NULL,
    wc4 DECIMAL(9, 6) NULL,
    CONSTRAINT station_readings_ibfk_1 FOREIGN KEY (station_id) REFERENCES stations (station_id) ON DELETE CASCADE
);

CREATE INDEX station_id ON station_readings (station_id);

CREATE TABLE student_members (
    student_member_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    student_type ENUM('graduate', 'undergraduate') NOT NULL
);
