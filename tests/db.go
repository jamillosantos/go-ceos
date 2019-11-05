package tests

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	. "github.com/onsi/gomega"
	"github.com/ory/dockertest"
)

var (
	dockerPool     *dockertest.Pool
	dockerResource *dockertest.Resource
	DB             *sql.DB
)

const (
	dbName = "ceoustest"
	dbUser = "ceous"
)

func DBStart() {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Second * 30
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	dockerPool = pool

	resource, err := pool.Run("postgres", "11-alpine", []string{"POSTGRES_DB=" + dbName, "POSTGRES_USER=" + dbUser})
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	dockerResource = resource

	Expect(pool.Retry(func() error {
		db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s "+
			"dbname=%s sslmode=disable", "localhost", resource.GetPort("5432/tcp"), dbUser, dbName))
		if err != nil {
			return err
		}
		err = db.Ping()
		if err != nil {
			return err
		}
		DB = db
		InitDefault(DB)
		return nil
	})).To(Succeed())
}

func DBUsersCreate() {
	_, err := DB.Exec("CREATE TABLE users (id SERIAL PRIMARY KEY, name varchar(100), password varchar(32) not null default '', role varchar(30) not null default '', created_at timestamptz default now(), updated_at timestamptz default now())")
	ExpectWithOffset(1, err).ToNot(HaveOccurred())
}

func DBUserGroupsCreate() {
	_, err := DB.Exec("CREATE TABLE user_groups (user_id int, group_id int, admin boolean not null default false)")
	ExpectWithOffset(1, err).ToNot(HaveOccurred())
}

func DBUsersInsertJoes() {
	_, err := DB.Exec("INSERT INTO users (name) VALUES ($1),($2),($3),($4);", "Snake Eyes", "Scarlet", "Tank", "Duke")
	ExpectWithOffset(1, err).ToNot(HaveOccurred())
}

func DBUserGroupsInsert() {
	_, err := DB.Exec("INSERT INTO user_groups (user_id, group_id) VALUES ($1, $2),($1, $3),($2, $1),($4, $1);", 1, 2, 3, 4)
	ExpectWithOffset(1, err).ToNot(HaveOccurred())
}

func DBStop() {
	Expect(DB.Close()).To(Succeed())
	ExpectWithOffset(1, dockerPool.Purge(dockerResource)).To(Succeed())
	Default = nil
}
