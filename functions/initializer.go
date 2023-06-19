package functions

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

var (
	client *firestore.Client
	ctx    context.Context
)

func init() {
	var err error
	ctx = context.Background()
	client, err = firestore.NewClient(ctx, "project-name-here")

	if err != nil {
		log.Fatalf("firestore new error: %s\n", err)
	}
}
