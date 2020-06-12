package cmd

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r * http.Request){
	fmt.Fprintf(w,"Merbaha %s",r.URL.Path[1])
}
func main(){
	http.HandleFunc("/",handler)
	http.ListenAndServe("8000",nil)
	fmt.Println("Web server is started!")
}