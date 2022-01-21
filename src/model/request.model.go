package model

import "fmt"

type Request interface {
	Fetch()
}

/* ---------------------------- Base ---------------------------- */
type Requester struct{}

func NewRequester() *Requester {
	return &Requester{}
}

func (this *Requester) Fetch() {
	fmt.Println("request data")
}

/* ------------------------------- Decorator 1 ------------------------------ */
type SecurityRequester struct{ Request }

func NewSecurityRequester(request Request) *SecurityRequester {
	return &SecurityRequester{Request: request}
}

func (this *SecurityRequester) Fetch() {
	fmt.Println("generating security key")
	this.Request.Fetch()
}

/* ------------------------------- Decorator 2 ------------------------------ */
type CompressResponse struct{ Request }

func NewCompressResponse(request Request) *CompressResponse {
	return &CompressResponse{Request: request}
}

func (this *CompressResponse) Fetch() {
	this.Request.Fetch()
	fmt.Println("compressing response")
}
