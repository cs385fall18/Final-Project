
CREATE DATABASE IF NOT EXISTS minibank DEFAULT CHARACTER SET utf8mb4;

USE minibank;

CREATE TABLE IF NOT EXISTS account (
    id INTEGER NOT NULL AUTO_INCREMENT,
    username CHAR(30) NOT NULL,
    password CHAR(60) NOT NULL,
    timestamp BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY(id),
    UNIQUE KEY(username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS sessions (
    session CHAR(36) NOT NULL,
    username CHAR(30) not null,
    expiration BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY(session)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS token_user {

	token CHAR(200) NOT NULL,
	username CHAR(30) NOT NULL,
	PRIMARY KEY (token),
	/* do we need 1 token per user?  */
}	ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


/*type JWTToken struct {
	Token string `json:"token"`
	Username string `json:"username"`
}*/


CREATE USER IF NOT EXISTS 'minibank'@'%' IDENTIFIED by 'minibank';

GRANT ALL PRIVILEGES ON minibank.* to 'minibank'@'%';

FLUSH PRIVILEGES;
