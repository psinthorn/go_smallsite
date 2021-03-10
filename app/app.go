package app

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Start() {

	fmt.Println(fmt.Sprintf("Server is started on port %s", portNumber))
	err := http.ListenAndServe(":8080", routes())
	if err != nil {
		fmt.Println(fmt.Sprintf("can't start server on port %s", portNumber))
	}

}
