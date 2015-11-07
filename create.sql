CREATE TABLE comments (
    id INT UNIQUE NOT NULL,
    body VARCHAR(255),
    user_name VARCHAR(255),
    file_path VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
