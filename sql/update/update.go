package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")

	stmt.Exec("Strawlley Pessoa", 1)
	stmt.Exec("Warysson Gomes", 2)

	stmt2, _ := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	stmt2.Exec(3)

}
