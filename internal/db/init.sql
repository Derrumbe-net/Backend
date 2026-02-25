CREATE TABLE admins (
    admin_id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(100),
    is_authorized TINYINT(1) DEFAULT 0,
    is_email_verified TINYINT(1) DEFAULT 0
);

CREATE TABLE stations (
    station_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    land_unit VARCHAR(50),
    geological_unit VARCHAR(50),
    susceptibility VARCHAR(50),
    depth VARCHAR(50),
    landslide_forecast DECIMAL(6,2),
    image_path VARCHAR(100),
    latitude DECIMAL(9,6),
    longitude DECIMAL(9,6),
    elevation INT,
    slope DECIMAL(10,2),
    is_available TINYINT(1) DEFAULT 1,
    collaborator VARCHAR(100),
    station_installation_date DATETIME
);

CREATE TABLE landslides (
    landslide_id INT AUTO_INCREMENT PRIMARY KEY,
    landslide_date DATETIME,
    latitude DECIMAL(9,6),
    longitude DECIMAL(9,6),
    image_path VARCHAR(512)
);

CREATE TABLE faculty_members (
    faculty_member_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    role VARCHAR(100),
    email VARCHAR(255),
    phone VARCHAR(50),
    extension VARCHAR(50),
    linkedin_url VARCHAR(512),
    profile_image_path VARCHAR(100)
);

CREATE TABLE student_members (
    student_member_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(20)
);

CREATE TABLE projects (
    project_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(150),
    description TEXT,
    start_year SMALLINT UNSIGNED,
    end_year SMALLINT UNSIGNED,
    project_status ENUM('planned', 'active', 'paused', 'completed', 'archived') DEFAULT 'active',
    image_path VARCHAR(512)
);

CREATE TABLE publications (
    publication_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(150),
    description TEXT,
    publication_url VARCHAR(512),
    image_path VARCHAR(512),
    published_date DATE
);

CREATE TABLE admin_tokens (
    token_id INT AUTO_INCREMENT PRIMARY KEY,
    admin_id INT NOT NULL,
    verification_token VARCHAR(64) NOT NULL,
    token_expires_at DATETIME NOT NULL,
    FOREIGN KEY (admin_id) REFERENCES admins(admin_id) ON DELETE CASCADE
);

CREATE TABLE station_readings (
    reading_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    station_id INT NOT NULL,
    recorded_at DATETIME NOT NULL,
    precipitation DECIMAL(6,2),
    wc1 DECIMAL(9,6),
    wc2 DECIMAL(9,6),
    wc3 DECIMAL(9,6),
    wc4 DECIMAL(9,6),
    FOREIGN KEY (station_id) REFERENCES stations(station_id) ON DELETE CASCADE
);

CREATE TABLE reports (
    report_id INT AUTO_INCREMENT PRIMARY KEY,
    landslide_id INT,
    reported_at DATETIME,
    latitude DECIMAL(9,6),
    longitude DECIMAL(9,6),
    city VARCHAR(100),
    physical_address VARCHAR(512),
    reporter_name VARCHAR(100),
    reporter_phone VARCHAR(30),
    reporter_email VARCHAR(255),
    description TEXT,
    image_path VARCHAR(512),
    is_validated TINYINT(1) DEFAULT 0,
    FOREIGN KEY (landslide_id) REFERENCES landslides(landslide_id) ON DELETE SET NULL
);
