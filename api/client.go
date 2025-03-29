package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// searches for a package
func searchPackage(packageName string, c *gin.Context) {
	url := ""
}

func FetchData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch data, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
