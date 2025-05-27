# 📦 Golang CRUD com MySQL

Este é um projeto básico de API RESTful em Go que realiza operações de **Create**, **Read**, **Update** e **Delete** (CRUD) em um banco de dados **MySQL**.

---

## 🚀 Tecnologias Utilizadas

- [Go (Golang)](https://golang.org/)
- [MySQL](https://www.mysql.com/)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Godotenv](https://github.com/joho/godotenv)

---

## 📁 Estrutura de Pastas

```
📦meu-crud-golang
 ┣ 📂banco
 ┃ ┗ 📄banco.go
 ┣ 📂servidor
 ┃ ┗ 📄servidor.go
 ┣ 📄main.go
 ┣ 📄.env
 ┣ 📄go.sum
 ┣ 📄.gitignore
 ┣ 📄readme
 ┗ 📄go.mod
 
```

---

## ⚙️ Configuração

1. **Clone o repositório**

```bash
git clone https://github.com/seu-usuario/seu-repo.git
cd seu-repo
```


2. **Crie a tabela no MySQL**

```sql
CREATE TABLE usuarios (
  id INT AUTO_INCREMENT PRIMARY KEY,
  nome VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL UNIQUE
);
```

3. **Instale as dependências**

```bash
go mod tidy
```

4. **Execute o projeto**

```bash
go run main.go
```

---

## 📌 Endpoints da API

| Método | Rota               | Descrição                  |
|--------|--------------------|----------------------------|
| POST   | /usuarios          | Cria um novo usuário       |
| GET    | /usuarios          | Lista todos os usuários    |
| GET    | /usuarios/{id}     | Busca usuário por ID       |
| PUT    | /usuarios/{id}     | Atualiza um usuário        |
| DELETE | /usuarios/{id}     | Deleta um usuário          |

---

## 🧪 Exemplo de JSON para criação

```json
{
  "nome": "Fabio Henrique",
  "email": "fabio@email.com"
}
```

---

## 📫 Contato

Entre em contato comigo no [LinkedIn](https://www.linkedin.com/in/fabio-henrique-luz-dev) ou me siga no [Instagram](https://www.instagram.com/fabio.apenas/).

---

## 🧠 Licença

Este projeto está licenciado sob a licença MIT. Sinta-se livre para usar e modificar.
