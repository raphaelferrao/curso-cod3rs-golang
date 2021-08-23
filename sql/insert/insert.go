package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO usuarios(nome) VALUES(?)")
	stmt.Exec("Maria")
	stmt.Exec("Joao")

	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId()
	fmt.Println("ultimo id inserido: ", id)

	linhas, _ := res.RowsAffected()
	fmt.Println("linhas", linhas)
}
