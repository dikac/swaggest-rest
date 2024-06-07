package service

import "github.com/dikac/swaggest-rest/_examples/task-api/pkg/graceful"

// Locator defines application services.
type Locator struct {
	graceful.Shutdown

	TaskCreatorProvider
	TaskUpdaterProvider
	TaskFinderProvider
	TaskFinisherProvider
}
