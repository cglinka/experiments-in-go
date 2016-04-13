package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/guregu/kami"
	"golang.org/x/net/context"
)

func indexHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	k := ctx.Value("key")
	fmt.Println(k)
}

func main() {
	m := kami.New()

	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "val")
	m.Context = ctx

	m.Get("/", indexHandler)

	log.Fatalln(http.ListenAndServe(":8080", m))
}
