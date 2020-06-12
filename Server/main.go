package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	var (
		i ironman
		w wolverine
	)
	message1 := &messageHandler{"Mesaj 1"}
	message2 := &messageHandler{"Mesaj 2"}

	mux := http.NewServeMux()

	mux.HandleFunc("/",indexHandler)
	mux.HandleFunc("/index",indexHandler)
	mux.HandleFunc("/iletisim",contactHandler)
	mux.Handle("/ironman",i)
	mux.Handle("/wolverine",w)
	mux.Handle("/message1",message1)
	mux.Handle("/message2",message2)
	mux.HandleFunc("/search",search)
	_ = http.ListenAndServe(":8000", mux)
}
type ironman int
func(x ironman)ServeHTTP(rw http.ResponseWriter, r *http.Request){
	io.WriteString(rw,"I am Ironman!")
}
type wolverine int
func(x wolverine)ServeHTTP(rw http.ResponseWriter, r *http.Request){
	io.WriteString(rw,"I am Wolverine!")
}

func contactHandler(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("İletişim sayfası"))
}
func indexHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Index"))
	x := r.URL.Path[1:]
	w.Write([]byte(x))
}

type messageHandler struct {
	message string
}
func (m messageHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,m.message)
}

func search(wr http.ResponseWriter, r *http.Request){
	hl := r.FormValue("hl")
	source := r.FormValue("source")
	q := r.FormValue("q")

	data := "hl:"+hl+",source:"+source+",q:"+q
	wr.Write([]byte(data))
}