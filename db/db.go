package db

import (
	"database/sql"
	"proxy/model"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUrl  string = "root:GHkoyt12@tcp(localhost:3306)/proxy"
	server string = "mysql"
)

func AddRedirect(r model.Redirect) error {
	db, err := sql.Open(server, dbUrl)
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO redirect(from_url,to_url) VALUES(?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(r.From_URL, r.To_URL)
	if err != nil {
		return err
	}
	return err
}

func GetAllUrls() (*[]model.Redirect, error) {
	var results []model.Redirect
	db, err := sql.Open(server, dbUrl)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT * FROM redirect"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var r model.Redirect
		err = rows.Scan(&r.ID, &r.From_URL, &r.To_URL)
		if err != nil {
			return nil, err
		}
		results = append(results, r)
	}
	return &results, err
}
