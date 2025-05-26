package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // Importe Implicíto
	"os"
)

func Conectar() (*sql.DB, error) {
	stringConexao := os.Getenv("CONEXAO") //usuario:senha@/nome_do_banco"

	db, err := sql.Open("mysql", stringConexao)
	if err != nil{
		return nil, err
	}

	if err := db.Ping(); err != nil{
		return nil, err
	}

	return db, nil
}