package main

import (
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", htmlHandler)
    http.ListenAndServe(":9090", nil)
}

func IsDirectory(path string) (bool, error) {
    fileInfo, err := os.Stat(path)
    return fileInfo.IsDir(), err
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
    file := "pages" + r.URL.Path
    fileInfo, err := os.Stat(file)
    if err != nil {
        return 
    } else if fileInfo.IsDir() {
        if file[len(file) - 1:] != "/"{
            file += "/"
        }
        file += "page.html"
    }
    if _, err := os.Stat(file); os.IsNotExist(err) {
        file = "nogoblet.html"
    }
    http.ServeFile(w, r, file)
}
