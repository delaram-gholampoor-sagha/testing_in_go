package httpservice

import (
	"fmt"
	"log"
	"net/http"
)


func HomePage(w http.ResponseWriter , r *http.Request) {
     fmt.Fprintf(w , "Welcome to the home page")
	 fmt.Println("Endpoint Hit: HomePage")
}


func SetUp() {
   http.HandleFunc("/" , HomePage)
}

func main() {
	SetUp()
	log.Fatal(http.ListenAndServe(":10000" , nil))
}