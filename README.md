# API de Jogos Favoritos

Este projeto é uma API RESTful para gerenciar jogos favoritos dos usuários. Ele fornece funcionalidades para registro de usuários, autenticação, e manipulação de dados de jogos. A API é construída usando Go, Gin Framework, e GORM para comunicação com um banco de dados PostgreSQL.


## Funcionalidades

- Registro de Usuários: Criação de novos usuários com validação de entrada.
- Autenticação: Sistema de login com geração de token JWT.
- Gerenciamento de Jogos: CRUD (Create, Read, Update, Delete) para jogos, incluindo filtragem e ordenação.
- Status de Jogos: Atribuição de status aos jogos (Não jogado, Jogando, Zerado, Platinado).

## Tecnologias Utilizadas

- Go: Linguagem de programação utilizada para o desenvolvimento da API.
- Gin Framework: Framework web para Go, utilizado para criar roteamentos e middleware.
- GORM: ORM para Go, utilizado para interagir com o banco de dados PostgreSQL.
- PostgreSQL: Sistema de gerenciamento de banco de dados relacional.
- JWT: JSON Web Tokens para autenticação segura.
- Docker: Para execução do banco de dados.

## Prerequisites

- Go 1.16+
- PostgreSQL
- Configurar variáveis de ambiente no arquivo .env e docker-compose:

  DB_HOST=localhost <br/>
  DB_PORT=5432 <br/>
  DB_USER=root <br/>
  DB_PASSWORD=root <br/>
  DB_NAME=root <br/>
  JWT_SECRET_KEY=sua_secret_key

## Installation

1. **Clone o repositório:**
  ```
  git clone https://github.com/seu-usuario/seu-repositorio.git
  cd seu-repositorio
  ```

2. **Instale as dependências:**
  ```
  go mod tidy
  ```


3. **Configure o Banco de dados:**
  - Execute o Docker compose.


4. **Execute a aplicação:**
  ```
  go run main.go
  ```

## Uso

**Rotas Principais**
- Registro de Usuário: POST /auth/register
- Login de Usuário: POST /auth/login
- Listar Jogos: GET /api/games
- Filtrar Jogos: GET /api/games/filter
- Criar Jogo: POST /api/games
- Editar Jogo: PUT /api/games/:id
- Deletar Jogo: DELETE /api/games/:id

**Exemplo de requisição para registro:**

```
POST /auth/register
Content-Type: application/json

{
  "username": "exampleUser",
  "email": "user@example.com",
  "password": "securePassword"
}
```

**Exemplo de requisição pra login:**
```
POST /auth/login
Content-Type: application/json

{
  "username": "exampleUser",
  "password": "securePassword"
}
```

**Autenticação**

- As rotas da API (exceto registro e login) requerem um token JWT.

- Envie o token no cabeçalho de autorização:

- Authorization: Bearer {token}

## Contribuição

1. Fork o projeto.
2. Crie uma nova branch: 
  ```
  git checkout -b feature/nova-funcionalidade.
  ```
3. Faça suas alterações e commite:
  ```
  git commit -m 'Adiciona nova funcionalidade'.
  ```
4. Faça o push para a branch:
  ```
  git push origin feature/nova-funcionalidade.
  ```
5. Envie um Pull Request.