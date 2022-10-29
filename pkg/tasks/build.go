package tasks

import "log"

func Build(imageTag string) error {

	log.Printf("building image: %s", imageTag)

	return nil
}
