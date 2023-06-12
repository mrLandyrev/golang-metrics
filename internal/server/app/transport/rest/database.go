package rest

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuildPingHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := db.Ping()

		if err == nil {
			c.Status(http.StatusOK)

			return
		}

		c.Status(http.StatusInternalServerError)
	}
}
