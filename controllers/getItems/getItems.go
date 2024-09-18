package getItems

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sumansurvi/go-crud/config"
	"github.com/sumansurvi/go-crud/structures"
)

// Retrieve all items
func GetAllItems(c *gin.Context, appCtx *config.AppContext) {
	fmt.Println("GetALlItems", c)
	rows, err := appCtx.Config.Db.Query("SELECT id, name, description, price, created_at, updated_at FROM items")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving data"})
		return
	}
	defer rows.Close()

	var items []structures.Item
	for rows.Next() {
		var item structures.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.CreatedAt, &item.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning data"})
			return
		}
		items = append(items, item)
	}
	c.JSON(http.StatusOK, items)

}
