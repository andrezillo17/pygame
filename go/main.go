package main

import (
  "html/template"
  "net/http"
  "path"
)

type Profile struct {//se define la estructura que se enviará a la plantilla
  Name    string
  Hobbies []string
  Facultad []string
}

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe(":3000", nil)//se debe acceder a localhost:3000
}

func foo(w http.ResponseWriter, r *http.Request) {
  profile := Profile{"Alejandro Daza", []string{"Cóndor", "Bienestar"}, []string{"Ingeniería", "Educación","Artes","Tecnologica"}}

  lp := path.Join("templates", "layout.html")//se cargan las plantillas
  fp := path.Join("templates", "index.html")

  
  tmpl, err := template.ParseFiles(lp, fp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := tmpl.Execute(w, profile); err != nil {//Se escribe hacia el navegador
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
