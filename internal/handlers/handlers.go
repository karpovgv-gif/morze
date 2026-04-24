package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "../index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	input := string(data)

	result, err := service.Convert(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	currentTime := time.Now().UTC().String()

	safeName := strings.ReplaceAll(currentTime, " ", "_")

	safeName = strings.ReplaceAll(safeName, ":", "-")
	filename := fileHeader.Filename
	ext := filepath.Ext(filename) // ".txt"
	newFilename := "converted_" + strings.ReplaceAll(filename, ext, "") + "_morse" + ext

	if err := os.WriteFile(newFilename, []byte(result), 0666); err != nil {
		log.Fatal(err)
	}
}
