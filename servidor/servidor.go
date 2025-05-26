package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"github.com/joho/godotenv"
)

type usuario struct{
	ID uint32 `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil{
		w.Write([]byte("Falha ao ler o corpo da requisição."))
		return
	}

	var usuario usuario

	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil{
		w.Write([]byte("Falha ao converter usuário para struct"))
		return
	}

	fmt.Println(usuario)

	db, err := banco.Conectar()
	if err != nil{
		http.Error(w, "Erro ao se conectar com o banco: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO usuarios (nome, email) VALUES(?, ?)")
	if err != nil{
		w.Write([]byte("Erro ao criar statement"))
	}
	defer statement.Close()

	
	insercao, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar a inserção dos dados na tabela"))
	}

	idInserido, err := insercao.LastInsertId()
	if err != nil{
		w.Write([]byte("Erro ao obter o id inserido"))
	}

	// STATUS CODE 
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))
}