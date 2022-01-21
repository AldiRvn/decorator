package main

import (
	"decorator/src/model"
	"fmt"
)

func main() {
	fmt.Println("> Base")
	requester := model.NewRequester()
	requester.Fetch()

	fmt.Println("\n> Base->Decorator 1")
	requesterWithSecurity := model.NewSecurityRequester(requester)
	requesterWithSecurity.Fetch()

	fmt.Println("\n> Base->Decorator 2")
	model.NewCompressResponse(requester).Fetch()

	fmt.Println("\n> Base->Decorator 1->Decorator 2")
	model.NewCompressResponse(requesterWithSecurity).Fetch()
}
