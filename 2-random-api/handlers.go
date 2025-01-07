package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func RandomHandler(w http.ResponseWriter, r *http.Request) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := random.Intn(6) + 1
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("%d", randomNum)))
}
