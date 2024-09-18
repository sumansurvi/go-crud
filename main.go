package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/sumansurvi/go-crud/config"
	"github.com/sumansurvi/go-crud/routes"
)

/*
// Retrieve a single item by ID

	func getItem(c *gin.Context) {
		id := c.Param("id")
		var item Item
		err := db.QueryRow("SELECT id, name, description, price, created_at, updated_at FROM items WHERE id = $1", id).Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving data"})
			}
			return
		}
		c.JSON(http.StatusOK, item)
	}

// Create a new item

	func createItem(c *gin.Context) {
		var input Item
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		query := `INSERT INTO items (name, description, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
		createdAt := time.Now()
		updatedAt := createdAt

		err := db.QueryRow(query, input.Name, input.Description, input.Price, createdAt, updatedAt).Scan(&input.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data"})
			return
		}

		c.JSON(http.StatusOK, input)
	}

// Update an existing item

	func updateItem(c *gin.Context) {
		id := c.Param("id")
		var input Item
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		query := `UPDATE items SET name = $1, description = $2, price = $3, updated_at = $4 WHERE id = $5`
		_, err := db.Exec(query, input.Name, input.Description, input.Price, time.Now(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully"})
	}

// Delete an item

	func deleteItem(c *gin.Context) {
		id := c.Param("id")
		_, err := db.Exec("DELETE FROM items WHERE id = $1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
	}
*/
func main() {

	conf, err := config.Init()
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	// Start server

	routes.StartServer(config.NewAppContext(conf)).Run(":" + viper.GetString("app_port"))
}
