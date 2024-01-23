package fsx

import (
	"fmt"
	"io/fs"
	"log"
	"sync"
	"time"

	"github.com/Rockup-Consulting/std/x/randx"
)

type Listener func(name string)

type watch struct {
	f    fs.FS
	glob string
}

type Watcher struct {
	mu        sync.Mutex
	l         *log.Logger
	watchers  []watch
	t         time.Ticker
	running   bool
	stop      chan struct{}
	files     map[string]time.Time
	listeners map[string]Listener
	// lastUpdatedFile string
}

func NewWatcher(l *log.Logger) *Watcher {
	return &Watcher{
		l:         l,
		watchers:  []watch{},
		files:     map[string]time.Time{},
		listeners: map[string]Listener{},
	}
}

func (w *Watcher) AddFS(f fs.FS, glob string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.l.Println("fswatcher: adding filesystem")

	w.watchers = append(w.watchers, watch{
		f:    f,
		glob: glob,
	})

}

func (w *Watcher) getUniqueID() string {
	var randID string

	for {
		randID := randx.UID()

		// if a listener with this ID is already registered, we just pick a new one
		if _, ok := w.listeners[randID]; ok {
			continue
		}

		break
	}

	return randID
}

func (w *Watcher) Running() bool {
	return w.running
}

// Listen registers a listener. A cleanup function is returned.
func (w *Watcher) Listen(f Listener) func() {
	w.mu.Lock()
	defer w.mu.Unlock()

	uid := w.getUniqueID()
	w.listeners[uid] = f

	w.l.Printf("fswatcher: registering listener -> %s", uid)

	return func() {
		w.mu.Lock()
		defer w.mu.Unlock()

		delete(w.listeners, uid)
	}
}

// Watch starts watching the filesystem for changes
func (w *Watcher) Watch() error {
	if w.running {
		return nil
	}

	w.running = true

	ticker := time.NewTicker(100 * time.Millisecond)
	stop := make(chan struct{})
	errChan := make(chan error, 1)

	w.stop = stop

	fileTracker := map[string]time.Time{}

	// first run through, just load the files into fileTracker, do this before the go routine kicks
	// off so that it is blocking
	for _, watch := range w.watchers {
		files, err := fs.Glob(watch.f, watch.glob)
		if err != nil {
			panic(fmt.Errorf("starting file watcher: %s", err))
		}

		for _, name := range files {
			file, err := watch.f.Open(name)
			if err != nil {
				fmt.Printf("error opening file: %s\n", name)
			}

			stat, err := file.Stat()
			if err != nil {
				panic(err)
			}

			_, trackingFile := fileTracker[name]
			if trackingFile {
				continue
			} else {
				fileTracker[name] = stat.ModTime()
			}
		}
	}

	go func(fileTracker map[string]time.Time) {
		for {
			select {
			case <-stop:
				return
			case <-ticker.C:
				for _, watch := range w.watchers {
					// if w.lastUpdatedFile != "" {
					// 	// first check this file then proceed to the rest
					// }

					files, err := fs.Glob(watch.f, watch.glob)
					if err != nil {
						errChan <- fmt.Errorf("starting file watcher: %s", err)
					}

					for _, name := range files {
						file, err := watch.f.Open(name)

						if err != nil {
							fmt.Printf("error opening file: %s\n", name)
						}

						stat, err := file.Stat()
						if err != nil {
							errChan <- err
						}

						oldModTime, trackingFile := fileTracker[name]
						if trackingFile {
							modTime := stat.ModTime()
							if !oldModTime.Equal(modTime) {
								fileTracker[name] = modTime

								// it takes tailwind roughly "12ms" (more like 30) to recalculate the styles
								// but I don't think they include writing to the file in their time estimate
								time.Sleep(time.Millisecond * 50)
								w.signalListeners(name)

								break
							}
						} else {
							fileTracker[name] = stat.ModTime()
							// we have a new file, so we should let our listeners know about that
							w.signalListeners(name)
							break
						}
					}
				}

			}
		}
	}(fileTracker)

	return <-errChan
}

func (w *Watcher) signalListeners(name string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	for _, l := range w.listeners {
		l(name)
	}
}

// Stop stops the filewatcher if it is currently running, if not no action is taken
func (w *Watcher) Stop() {
	if !w.running {
		return
	}

	w.t.Stop()
	w.stop <- struct{}{}
	w.running = false
}
