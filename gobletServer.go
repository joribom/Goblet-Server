/*
 *  DEPRECATED
 *  Has since 2016-12-28 been replaced with apache2
 */

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
        http.Redirect(w, r, "http://gobletdeathandrebirth.com/goblet-not-found", 301)
        return
    } else if fileInfo.IsDir() {
        if file[len(file) - 1:] != "/"{
            file += "/"
        }
        file += "page.html"
    }
    http.ServeFile(w, r, file)
}
