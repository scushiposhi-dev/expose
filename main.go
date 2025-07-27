package main

import (
	"encoding/json"
	"expose/internal/bar"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("starting something")

	http.HandleFunc("/v1/bar/{bar_id}/test", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("bar_id")
		svc := bar.NewService()
		bar := svc.GetBar(id)
		b, _ := json.Marshal(bar)

		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(":8585", nil))
}
