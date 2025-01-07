// basic hello world  in terminal app

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}


// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, World!")
// 	})
// 	http.ListenAndServe(":8080", nil)
// }