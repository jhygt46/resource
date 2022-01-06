package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"resource/utils"
	"runtime"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Filtro struct {
	C  []Campos `json:"C"`
	E  []Evals  `json:"E"`
	Id int32    `json:"Id"`
}
type Campos struct {
	T int      `json:"T"`
	N string   `json:"N"`
	V []string `json:"V"`
}
type Evals struct {
	T int    `json:"T"`
	N string `json:"N"`
}

func GetDbFiltroString(file string) (*sql.DB, error) {

	var path string
	if runtime.GOOS == "windows" {
		path = "C:/Allin/db/"
	} else {
		path = "/var/db/"
	}

	if !utils.FileExists(path + file) {
		db, err := sql.Open("sqlite3", path+file)
		if err == nil {
			stmt, err := db.Prepare("create table if not exists filtros (id integer not null primary key autoincrement,filtro text)")
			if err != nil {
				return nil, err
			}
			stmt.Exec()
			return db, nil
		} else {
			return nil, err
		}
	} else {
		db, err := sql.Open("sqlite3", path+file)
		if err == nil {
			return db, nil
		} else {
			return nil, err
		}
	}
}
func GetDbFiltroBytes(file string) (*sql.DB, error) {

	var path string
	if runtime.GOOS == "windows" {
		path = "C:/Allin/db/"
	} else {
		path = "/var/db/"
	}

	if !utils.FileExists(path + file) {
		db, err := sql.Open("sqlite3", path+file)
		if err == nil {
			stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS filtros (filtro BLOB NOT NULL, id INTEGER NOT NULL, PRIMARY KEY (id))")
			if err != nil {
				return nil, err
			}
			stmt.Exec()
			return db, nil
		} else {
			return nil, err
		}
	} else {
		db, err := sql.Open("sqlite3", path+file)
		if err == nil {
			return db, nil
		} else {
			return nil, err
		}
	}
}
func FiltroStringInit(db *sql.DB, filtro Filtro, total int64) {

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.
	stmt, err := tx.Prepare("INSERT INTO filtros (filtro) VALUES(?)")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	now := time.Now()
	for i := 1; i <= int(total); i++ {
		filtro.Id = int32(i)
		u, err := json.Marshal(filtro)
		if err == nil {
			if _, err := stmt.Exec(string(u)); err != nil {
				fmt.Println(err)
			}
		}
	}
	elapsed := time.Since(now)
	fmt.Printf("CREATE DB => TOTAL %v [%s] c/u total %v\n", total, utils.Time_cu(elapsed, int(total)), elapsed)

	if err := tx.Commit(); err != nil {
		fmt.Println(err)
	}
}
func FiltroBytesInit(db *sql.DB, filtro Filtro, total int64) {

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.
	stmt, err := tx.Prepare("INSERT INTO filtros (filtro) VALUES(?)")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	now := time.Now()
	for i := 1; i <= int(total); i++ {
		filtro.Id = int32(i)
		u, err := json.Marshal(filtro)
		if err == nil {
			if _, err := stmt.Exec(u); err != nil {
				fmt.Println(err)
			}
		}
	}
	elapsed := time.Since(now)
	fmt.Printf("CREATE DB => TOTAL %v [%s] c/u total %v\n", total, utils.Time_cu(elapsed, int(total)), elapsed)

	if err := tx.Commit(); err != nil {
		fmt.Println(err)
	}
}
func GetFiltroByteContent(db *sql.DB, id int64) ([]byte, error) {
	rows, err := db.Query("SELECT filtro FROM db2.filtros WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	var filtro []byte
	for rows.Next() {
		err := rows.Scan(&filtro)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()
	return filtro, nil
}
func GetFiltroStringContent(db *sql.DB, id int64) (string, error) {

	rows, err := db.Query("SELECT filtro FROM db1.filtros WHERE id=?", id)
	if err != nil {
		return "", err
	}
	var filtro string
	for rows.Next() {
		err := rows.Scan(&filtro)
		if err != nil {
			return "", err
		}
	}
	defer rows.Close()
	return filtro, nil
}
