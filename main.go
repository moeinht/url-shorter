package main

import (
	"fmt"
	urlstore "url-shorter/url-store"
)

func main() {
	store := urlstore.CreateStore()
	hashed := store.CreateURL("moein")
	hashed2 := store.CreateURL("moein")
	fmt.Println(hashed)
	fmt.Println(hashed2)
}
