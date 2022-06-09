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

func SelectProducts() (*sql.Rows, error) {
	db, err := sql.Open("mysql", "root:root@/trabalho_poo")

	if err != nil {
		return nil, err
	}

	stmt := "select * from produtos"
	rows, _ := db.Query(stmt)

	return rows, nil
}

func DeleteProduct(id int) (bool, error) {
	db, err := sql.Open("mysql", "root:root@/trabalho_poo")

	if err != nil {
		return false, err
	}

	stmt := "delete from produtos where id = ?"
	result, err := db.Exec(stmt, id)

	if err != nil {
		return false, err
	}

	rowsAffected, _ := result.RowsAffected()

	return rowsAffected > 0, nil
}

func UpdateProduct(id int, nome, marca string, quantidade int, preco float64) (bool, error) {
	db, err := sql.Open("mysql", "root:root@/trabalho_poo")

	if err != nil {
		return false, err
	}

	stmt := "update produtos set nome = ?, marca = ?, quantidade = ?, preco = ? where id = ?"
	result, err := db.Exec(stmt, nome, marca, quantidade, fmt.Sprintf("%.2f", preco), id)

	if err != nil {
		return false, err
	}

	rowsAffected, _ := result.RowsAffected()

	return rowsAffected > 0, nil
}
