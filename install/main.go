package main

import (
	authorv1 "buf.build/gen/go/sentiger/test/protocolbuffers/go/author/v1"
	bookv1 "buf.build/gen/go/sentiger/test/protocolbuffers/go/book/v1"
	"fmt"
)

func main() {
	book := &bookv1.Book{
		Id:    0,
		Title: "熟人",
		Author: &authorv1.Author{
			Id:   0,
			Name: "",
		},
	}
	fmt.Println(book)
}
