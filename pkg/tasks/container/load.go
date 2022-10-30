package container

import (
	"log"

	"dagger.io/dagger"
)

func LoadImage(client *dagger.Client, image string) *dagger.Container {
	log.Printf("loading image: %s", image)

	return client.Container().From(image)
}
