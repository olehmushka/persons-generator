package orchestrator

import (
	"fmt"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"
	"persons_generator/engine/entities/world"
)

func (o *Orchestrator) CreateWorld(
	size int,
	cultures []*culture.Culture,
	religions []*religion.Religion,
	refs []*religion.CultureReference,
) (*world.World, error) {
	return world.New(world.Config{StorageFolderName: o.storageFolderName}, size, cultures, religions, refs)
}

func (o *Orchestrator) RunAndSaveWorld(w *world.World, stopYear int) error {
	if w == nil {
		return nil
	}

	progressCh := make(chan world.ProgressRunWorld)
	errCh := make(chan error)
	var isFinished bool
	go w.RunWorld(stopYear, progressCh, errCh)
	for {
		select {
		case err := <-errCh:
			if err != nil {
				return err
			}
			isFinished = true
			close(errCh)
			break
		case progress := <-progressCh:
			fmt.Printf("year: %d ; population: %d ; progress: %f\n", progress.Year, progress.Population, progress.Progress)
			if progress.Progress == 100 {
				isFinished = true
				close(progressCh)
				break
			}
		}
		if isFinished {
			break
		}
	}

	return w.Save()
}
