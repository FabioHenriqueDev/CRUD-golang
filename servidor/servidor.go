package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type usuario struct{
	ID uint32 `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
}

// Cria um usuario e adiciona no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
		return
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

// Busca todos usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
		return
	}
	
	db, err := banco.Conectar()
	if err != nil{
		http.Error(w, "Erro ao se conectar com o banco: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()


	// SELECT * FROM usuarios
	linhas, err := db.Query("SELECT * FROM usuarios")
	if err != nil{
		w.Write([]byte("Erro ao fazer a consulta"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario
		if err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil{
			w.Write([]byte("Erro ao escanear usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
		
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil{
		w.Write([]byte("Erro ao converter usuarios ara json"))
		return
	}
}


// Traz um usuario especifico de um banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
		return
	}

	parametros := mux.Vars(r)

	ID, err := strconv.ParseInt(parametros["id"], 10, 32)
	if err != nil{
		w.Write([]byte("Erro ao converter o id para inteiro"))
		return
	}

	db, err := banco.Conectar()
	if err != nil{
		w.Write([]byte("Erro ao se conectar com o banco de dados"))
		return
	}

	linha, err := db.Query("SELECT * FROM usuarios WHERE ID = ?", ID)
	if err != nil{
		w.Write([]byte("Erro na consulta"))
		return
	}

	var usuario usuario
	if linha.Next(){
		if err = linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil{
			w.Write([]byte("Erro ao adicionar valores da consulta na variavel"))
			return
		}
	}

	if err = json.NewEncoder(w).Encode(usuario); err != nil{
		w.Write([]byte("Erro ao converer struct em json"))
		return
	}
}