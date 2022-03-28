CREATE DATABASE bank;

CREATE TABLE bank.account(
	id INT NOT NULL AUTO_INCREMENT,
	balance INT NOT NULL DEFAULT 0,
	status ENUM("frozen", "audited", "regular") NOT NULL DEFAULT "regular",
	opened TIMESTAMP NOT NULL,
	PRIMARY KEY(id)
);

INSERT INTO bank.account (balance, status, opened)
VALUES (15, "regular", '2010-01-01 00:00:01.000000');

INSERT INTO bank.account (balance, status, opened)
VALUES (27, "regular", '2010-01-01 00:00:01.000000');