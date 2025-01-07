package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func RandomHandler(w http.ResponseWriter, r *http.Request) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Fprint(w, random.Intn(6)+1)
}
