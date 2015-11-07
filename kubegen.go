package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

type Form struct {
  Name string
  ApiVersion []byte
  //Selectors string
  // Template string
  // Spec string
}

func main() {
  // take a form struct and compose a manifest file
  // returns manifest file in yaml format
  http.HandleFunc("/", viewHandler)
  http.ListenAndServe(":8080", nil)
}

func (f *Form) save() error {
  filename := f.Name + ".yaml"
  return ioutil.WriteFile(filename, f.ApiVersion, 0600)
}

func loadForm(name string) *Form{
  filename := name + ".yaml"
  apiVersion, _ := ioutil.ReadFile(filename)
  return &Form{Name: name, ApiVersion: apiVersion}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  name := "name"
  f := loadForm(name)
  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", f.Name, f.ApiVersion)
}
