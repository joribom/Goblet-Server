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
        req, _ := http.NewRequest("GET", "http://gobletdeathandrebirth.com/goblet-not-found", nil)
        htmlHandler(w, req)
        return
    } else if fileInfo.IsDir() {
        if file[len(file) - 1:] != "/"{
            file += "/"
        }
        file += "page.html"
    }
    http.ServeFile(w, r, file)
}
