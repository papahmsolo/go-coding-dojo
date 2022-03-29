package service

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
	"github.com/vrublevski/sqlc/query"

	_ "github.com/go-sql-driver/mysql"
)

var (
	testDb *sql.DB
)

func TestMain(m *testing.M) {
	if err := os.Chdir(".."); err != nil {
		log.Println("unable to switch to parent directory:", err)
		return
	}
	pwd, err := os.Getwd()
	if err != nil {
		log.Println("unable to get working directory:", err)
		return
	}

	sqlPath := path.Join(pwd, "testdata")
	log.Println("using testdata in ", sqlPath)
	mysqlCfg := dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "5.7",
		Env:        []string{"MYSQL_ROOT_PASSWORD=password"},
		Mounts:     []string{sqlPath + ":/docker-entrypoint-initdb.d"},
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Println("unable to connect to docker:", err)
		return
	}
	mysqlResource, err := pool.RunWithOptions(&mysqlCfg)
	if err != nil {
		log.Printf("starting container: %s", err)
	}
	defer func() {
		if err := mysqlResource.Close(); err != nil {
			log.Println("unable to purge mysql resource:", err)
		}
	}()

	var trier trier
	trier.mysqlResource = mysqlResource
	if err := pool.Retry(trier.retry); err != nil {
		log.Println("unable to connect to mysql server:", err)
		return
	}
	testDb = trier.db
	defer testDb.Close()

	m.Run()
}

func TestDeductFromAccounts(t *testing.T) {
	accountant := NewAccountant(testDb)
	tm := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	accounts, err := accountant.DeductFromAccounts(10, tm)
	require.NoError(t, err)
	require.Equal(t, 1, len(accounts), "expected one account")

	account := accounts[0]
	require.Equal(t, account.Balance, int32(5))
	require.Equal(t, account.Status, query.AccountStatusFrozen)
}

type trier struct {
	db            *sql.DB
	mysqlResource *dockertest.Resource
}

func (t *trier) retry() error {
	var err error
	dataSrc := fmt.Sprintf("root:password@(localhost:%s)/mysql?parseTime=true", t.mysqlResource.GetPort("3306/tcp"))
	t.db, err = sql.Open("mysql", dataSrc)
	if err != nil {
		log.Println(err)
		return err
	}
	err = t.db.Ping()
	if err != nil {
		log.Println(err)
	}
	return err
}
