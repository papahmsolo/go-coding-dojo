CREATE DATABASE bank;

CREATE TABLE bank.account(
	id INT NOT NULL AUTO_INCREMENT,
	balance INT NOT NULL DEFAULT 0,
	status ENUM("frozen", "audited", "regular") NOT NULL DEFAULT "regular",
	opened TIMESTAMP NOT NULL,
	PRIMARY KEY(id)
);

