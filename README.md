# Favorite Games API

This is a RESTful API for managing a collection of favorite games. It's built with Go, using Gin as the web framework and GORM for database operations.

## Features

- List all games
- Add a new game
- Edit an existing game
- Delete a game
- Filter games by year and minimum rating
- Sort games by name, rating, or release year

## Prerequisites

- Go 1.16+
- PostgreSQL

## Installation

1. Clone the repository:
git clone 

github.com


2. Navigate to the project directory:
cd favorites-games-api


3. Install dependencies:
go mod tidy


4. Set up your environment variables in a `.env` file:
DB_HOST=your_db_host DB_PORT=your_db_port DB_USER=your_db_user DB_PASSWORD=your_db_password DB_NAME=your_db_name


5. Run the application:
go run main.go


The server will start on `http://localhost:8080`.

## API Documentation

API documentation is available via Swagger UI at `http://localhost:8080/swagger/index.html` when the application is running.

## Endpoints

- `GET /games`: Get all games
- `POST /games`: Create a new game
- `PUT /games/:id`: Update a game
- `DELETE /games/:id`: Delete a game
- `GET /games/filter`: Filter games by year and minimum rating