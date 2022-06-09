package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertProduct(nome, marca string, quantidade int, preco float64) (bool, error) {
	db, err := sql.Open("mysql", "root:root@/trabalho_poo")

	if err != nil {
		return false, err
	}

	stmt := "insert into produtos (nome, marca, quantidade, preco) values (?, ?, ?, ?)"
	result, err := db.Exec(stmt, nome, marca, quantidade, fmt.Sprintf("%.2f", preco))

	if err != nil {
		return false, err
	}

	rowsAffected, _ := result.RowsAffected()

	return rowsAffected > 0, nil
}