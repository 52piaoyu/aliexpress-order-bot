package db

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func AddUser(db *sql.DB, user User) error {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO users (id, aliexpress_login, aliexpress_password) VALUES (?, ?, ?)")
	_, err := stmt.Exec(user.ID, user.AliexpressLogin, user.AliexpressPassword)
	if err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
	}
	return nil
}

func GetUser(db *sql.DB, id int) (*User, error) {
	rows, err := db.Query("select * from users where id=?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempUser User
		err = rows.Scan(&tempUser.ID, &tempUser.AliexpressLogin, &tempUser.AliexpressPassword)
		if err != nil {
			return nil, err
		}
		return &tempUser, nil
	}
	return nil, nil
}

func UpdateUser(db *sql.DB, id, user User) error {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update users set aliexpress_login=?,aliexpress_password=?, where id=?")
	_, err := stmt.Exec(user.AliexpressLogin, user.AliexpressPassword, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func DeleteUser(db *sql.DB, id2 int) error {
	sid := strconv.Itoa(id2) // int to string
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from users where id=?")
	_, err := stmt.Exec(sid)
	if err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return nil
	}
}

func AddIndex(db *sql.DB, index Index) error {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO users (id, index, url, location, last_modification) VALUES (?, ?, ?, ?, ?)")
	_, err := stmt.Exec(index.ID, index.Index, index.URL, index.Location, index.LastModification.String())
	if err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
	}
	return nil
}

func GetIndex(db *sql.DB, id int) (*Index, error) {
	rows, err := db.Query("select * from indexes where id=?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempIndex Index
		var timeStr string
		err = rows.Scan(&tempIndex.ID, &tempIndex.Index, &tempIndex.URL, &tempIndex.Location, &timeStr)
		if err != nil {
			return nil, err
		}
		tempIndex.LastModification, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timeStr)
		if err != nil {
			return nil, err
		}
		return &tempIndex, nil
	}
	return nil, nil
}

func UpdateIndex(db *sql.DB, id, index Index) error {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update indexes set index=?, url=?, location=?, last_modification=?, where id=?")
	_, err := stmt.Exec(index.Index, index.URL, index.Location, index.LastModification.String(), id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func DeleteIndex(db *sql.DB, id2 int) error {
	sid := strconv.Itoa(id2) // int to string
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from indexes where id=?")
	_, err := stmt.Exec(sid)
	if err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return nil
	}
}
