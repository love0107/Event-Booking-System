CREATE TABLE Event (
                       id BIGINT AUTO_INCREMENT PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       description TEXT,
                       date VARCHAR(255),
                       location VARCHAR(255),
                       createdOn DATETIME DEFAULT CURRENT_TIMESTAMP,
                       updatedOn DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);