package main

//minute 10:19
import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"os"
	"log"
)

var INDEX_HTML []byte

func main() {
	fmt.Println("starting server on http://localhost:3000/")
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/redirect", CreateRedirectHandler)
	http.ListenAndServe(":3000", nil)

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello"))
	log.Println("GET /")
	w.Write(INDEX_HTML)
}

func CreateRedirectHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("creating redirect handler", r.Form)
}

//does this function get called automatically by main?
func init() {
	//f, err := os.Open("./tmpl/index.html")
	INDEX_HTML, _ = ioutil.ReadFile("./tmpl/index.html")
}
