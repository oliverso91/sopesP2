// File: main.go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/handlers"
)

type Persona struct {
    Nombre string
    Departamento string
    Edad  int
    Contagio string `json:"forma de contagio"`
    Estado string

}

type Parametros struct {
    Url string
    Hilos int
    Solicitud  int

}

type TodaData struct {
  Personas [] Persona
  Parametro Parametros
}


/*func personCreate(w http.ResponseWriter, r *http.Request) {
    // Declare a new Person struct.
    var p TodaData

    // Try to decode the request body into the struct. If there is an error,
    // respond to the client with the error message and a 400 status code.
    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Do something with the Person struct...
  fmt.Println(r.Header.Get("Origin"))
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "https://preview.c9users")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
  //w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "%+v", http.StatusOK)
    DatoPersonas(p);
}*/

func Parametro(w http.ResponseWriter, r *http.Request) {
    // Declare a new Person struct.
    var p TodaData

    // Try to decode the request body into the struct. If there is an error,
    // respond to the client with the error message and a 400 status code.
    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Do something with the Person struct...
  fmt.Println(r.Header.Get("Origin"))
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "https://preview.c9users")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
  //w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "%+v", http.StatusOK)
    fmt.Println(p.Parametro.Url)
  //  fmt.Println(p.Personas)

    pFor := p.Parametro.Solicitud


    for i := 0; i < pFor; i++ {
          fmt.Println(p.Personas[i])
    }

}

///////////////////////enviar valores


/////////////////////////

func main() {
    mux := http.NewServeMux()
  //  mux.HandleFunc("/person/create", personCreate)
    mux.HandleFunc("/parametro", Parametro)

    err := http.ListenAndServe(":4000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(mux))
    log.Fatal(err)



}
