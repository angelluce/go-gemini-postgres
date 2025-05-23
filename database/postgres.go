package database

import (
	"database/sql"
	"fmt"
	"log"

	"go-gemini-postgres/config"
	"go-gemini-postgres/models"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(cfg *config.Config) {
	var err error
	DB, err = sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	fmt.Println("¡Conectado exitosamente a PostgreSQL!")
}

func GetItemByID(id int) (*models.Item, error) {
	row := DB.QueryRow("SELECT id, name, description FROM items WHERE id = $1", id)

	item := &models.Item{}
	err := row.Scan(&item.ID, &item.Name, &item.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Item no encontrado con ID: %d", id)
		}
		return nil, fmt.Errorf("Error al obtener el item: %v", err)
	}
	return item, nil
}

func GetAllItems() ([]models.Item, error) {
	rows, err := DB.Query("SELECT id, name, description FROM items")
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta para obtener todos los ítems: %v", err)
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		item := models.Item{}
		err := rows.Scan(&item.ID, &item.Name, &item.Description)
		if err != nil {
			return nil, fmt.Errorf("Error al obtener los items: %v", err)
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error durante la iteración de filas: %v", err)
	}

	return items, nil
}
