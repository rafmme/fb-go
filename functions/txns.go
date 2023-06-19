package functions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"google.golang.org/api/iterator"
)

func init() {
	functions.HTTP("FetchTxns", fetchTxns)
}

func fetchTxns(w http.ResponseWriter, r *http.Request) {
	iter := client.Collection("tests").Documents(ctx)
	defer iter.Stop()
	defer client.Close()

	data := []map[string]interface{}{}

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Error fetching Doc: %s", err)
		}

		data = append(data, doc.Data())
	}

	result, err := json.Marshal(data)

	if err != nil {
		log.Fatalf("Error fetching Doc: %s", err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(w, string(result))
}
