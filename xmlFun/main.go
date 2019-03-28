package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Human struct {
	FName 	string
	LName 	string
	Age 	int
}

func (h Human)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	h.FName = "Hasan"
	h.LName	= "URAL"
	h.Age	= 2

	r.ParseForm()

	fmt.Println(r.Form)

	fmt.Println("path",r.URL.Path)

	strBody := `<table>
					<thead><tr><th>İsim</th><th>Soyisim</th><th>Yaş</th></tr></thead>
					<tbody><tr><td>`+h.FName+`</td><td>`+h.LName+`</td><td>`+strconv.Itoa(h.Age)+`</td></tr></tbody>
				</table>`
	fmt.Fprintf(w,strBody)

}
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Merhaba %s",r.URL.Path[1:])
}
func CheckError(err error){
	if err != nil{
		panic(err)
	}
}
func main() {
	//http.HandleFunc("/",handler)
	var h Human
	http.ListenAndServe(":2525",h)
	//fmt.Println("Web Server")
}
