package build

import (
	"errors"
	"sync"
)

type goncurrency struct {
	workers     int
	callback    func() error
	flowControl bool
	// TODO: Include an attribute to set unlimited workers
}

func (g *goncurrency) Run() {
  if g.flowControl {
 			g.controlledExecution()   
  }
}
func (g *goncurrency) controlledExecution() {
	ch := make(chan error)
	wg := sync.WaitGroup{}

	for i := 0; i < g.workers; i++ {
		wg.Add(1)
		go func() {
			err := g.callback()
			ch <- err
		}()
	}
	defer wg.Done()

}

func (g *goncurrency) validateConfiguration() error {
	if g.workers < 1 {
		return errors.New("workers must be greater than 0")
	}

	if g.callback == nil {
    return errors.New("callback must be set")
	}
	
	return nil
}

func Build(
	workers int,
	callback func() error,
	flowControl bool,
) (error, *goncurrency) {
	builder := &goncurrency{
		workers:     workers,
		callback:    callback,
		flowControl: flowControl,
	}
	err := builder.validateConfiguration()

	if err != nil {
	  return err, nil
	}

	return nil, builder
}
