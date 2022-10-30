package container

import (
	"dagger.io/dagger"
)

func Workdir(container *dagger.Container, into string) *dagger.Container {
	return container.WithWorkdir(into)
}
