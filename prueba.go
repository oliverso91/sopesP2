package main

import (
    "net/http"
    "fmt"
    "log"
    "io/ioutil"
    "strings"
)

func main() {

    // request body (payload)
    requestBody := strings.NewReader(`
      {
"Nombre": "Preuba",
"Departamento" : "Guatemala",
"Edad": 12,
"forma de contagio" : "positivo",
"Estado" :"recuperado"
}
    `)

    // post some data
    res, err := http.Post(
        "http://localhost:4001/persona",
        "application/json; charset=UTF-8",
        requestBody,
    )

    // check for response error
    if err != nil {
        log.Fatal( err )
    }

    // read response data
    data, _ := ioutil.ReadAll( res.Body )

    // close response body
    res.Body.Close()

    // print request `Content-Type` header
    requestContentType := res.Request.Header.Get( "Content-Type" )
    fmt.Println( "Request content-type:", requestContentType )

    // print response body
    fmt.Printf( "%s\n", data )

}
