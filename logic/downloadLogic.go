package logic

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/razpinator/gmx/models"
)

func DownloadFile(item models.Item) error {
	url, output := item.Source, item.OutputFile
	// Create the HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making the request: %v", err)
	}
	defer response.Body.Close()

	// Check if the response status code is successful (200 OK)
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", response.Status)
	}

	// If output is not provided, use the filename from the URL
	if output == "" {
		output = filepath.Base(url)
	}

	// Create the file
	file, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("error creating the file: %v", err)
	}
	defer file.Close()

	// Write the body to file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("error writing the file: %v", err)
	}

	return nil
}
