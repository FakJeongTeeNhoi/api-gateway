package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

const headerAuthorization = "authorization"

func SetAccountInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token from header
		token := c.GetHeader(headerAuthorization)

		if token == "" {
			c.Next()
			return
		}

		// call UMS service to get account info with
		req, err := http.NewRequest("GET", os.Getenv("UMS_URL")+"/api/auth/account-info", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		req.Header.Set(headerAuthorization, token)
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if resp.StatusCode != http.StatusOK {
			c.AbortWithStatusJSON(resp.StatusCode, gin.H{"error": "Failed to get account info"})
			return
		}

		accountInfo := make(map[string]interface{})
		// read response body
		err = json.NewDecoder(resp.Body).Decode(&accountInfo)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// delete field success in accountInfo
		delete(accountInfo, "success")

		// add account info to context
		c.Set("user", accountInfo)

		c.Next()
	}
}
