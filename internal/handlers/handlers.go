package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "../index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
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
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	currentTime := time.Now().UTC().Format("2006-01-02_15-04-05")

	filename := fileHeader.Filename
	if filename == "" {
		filename = "file"
	}
	ext := filepath.Ext(filename)
	if ext == "" {
		ext = ".txt"
	}

	// base без расширения
	base := strings.TrimSuffix(filename, ext)

	// если надо — можно добавить префикс "converted_"
	newFilename := fmt.Sprintf("%s_%s%s", base, currentTime, ".txt")

	if err := os.WriteFile(newFilename, []byte(result), 0666); err != nil {
		// log.Fatal(err) — не здесь!
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
