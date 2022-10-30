package container

import (
	"log"

	"dagger.io/dagger"
)

func CreateScratch(client *dagger.Client) *dagger.Container {
	log.Println("creating scratch image")

	return client.Container(dagger.ContainerOpts{ID: ""})
}
