package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"testing/todo_app/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB
var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid STRING NOT NULL UNIQUE,
	name STREING,
	email STRING,
	password STRING,
	created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
