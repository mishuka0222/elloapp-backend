CREATE TABLE IF NOT EXISTS users_ello
(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT(20) NOT NULL,
    username VARCHAR(200) NOT NULL,
    password VARCHAR(64) NOT NULL,
    email VARCHAR(64) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    avatar JSON,
    link VARCHAR(60),
    date_of_birth DATETIME NOT NULL,
    email_confirmed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE (email),
    UNIQUE INDEX users_ello_user_id_IDX (user_id),
    CONSTRAINT users_ello_user_idFKusers_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS confirmation_codes
(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT(20) NOT NULL,
    code VARCHAR(6) DEFAULT NULL,
    expired_at TIMESTAMP NOT NULL,

    UNIQUE INDEX confirmation_codes_user_id_IDX (user_id),
    CONSTRAINT confirmation_codes_user_idFKusers_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE users MODIFY COLUMN phone varchar(32);
ALTER TABLE confirmation_codes MODIFY COLUMN code varchar(6);