
-- name: DeductFromAccount :exec
UPDATE bank.account 
	SET balance = balance - ?
	WHERE balance > ? AND status = ? AND opened > ?;

-- name: CreateAccount :execresult  
INSERT INTO bank.account (balance, status, opened)
VALUES (?, ?, ?);

-- name: FreezeAccounts :exec
UPDATE bank.account 
	SET status = "frozen"
	WHERE balance <= ? AND status = "regular";

-- name: ListAccounts :many
SELECT * FROM bank.account WHERE status = ?;