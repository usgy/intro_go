package main

import "net/http"

func viewHandler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello, %q", r.URL.Path[1:])
}

func main() {
  http.HandleFunc("/", viewHandler)
  http.ListenAndServe(":8080", nil)
}

