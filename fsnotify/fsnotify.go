package tyfsnotify

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
)

// 文件系统监控
type FileWatcher struct {
	watch *fsnotify.Watcher
	fn    func(file string)
	done  chan struct{}
	wg    sync.WaitGroup
}

func init() {

}

func New() (*FileWatcher, error) {
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, errors.New("create watcher failed")
	}

	fileWatcher := &FileWatcher{
		watch: watch,
		done:  make(chan struct{}, 1),
	}

	return fileWatcher, nil
}

func (fw *FileWatcher) watchDir(dir string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			path, err = filepath.Abs(path)
			if err != nil {
				return err
			}
			if err = fw.watch.Add(path); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	fw.wg.Add(1)
	go fw.run()

	return nil
}

func (fw *FileWatcher) Start(dir string, fn func(file string)) error {
	fw.fn = fn

	if err := fw.watchDir(dir); err != nil {
		return err
	}

	return nil
}

func (fw *FileWatcher) Stop() {
	if fw == nil {
		return
	}
	close(fw.done)
	fw.wg.Wait()
}

func (fw *FileWatcher) run() {
	defer func() {
		fw.wg.Done()
	}()

	for {
		select {
		case evs := <-fw.watch.Events:
			if evs.Op&fsnotify.Create == fsnotify.Create {
				finfo, err := os.Stat(evs.Name)
				if err == nil && finfo.IsDir() {
					err = fw.watch.Add(evs.Name)
					if err != nil {
					}
				}
			}

			if evs.Op&fsnotify.Write == fsnotify.Write {
				fw.fn(evs.Name)
			}

			if evs.Op&fsnotify.Remove == fsnotify.Remove {
				finfo, err := os.Stat(evs.Name)
				if err == nil && finfo.IsDir() {
					err = fw.watch.Remove(evs.Name)
					if err != nil {
					}
				}
			}

			if evs.Op&fsnotify.Rename == fsnotify.Rename {
				err := fw.watch.Remove(evs.Name)
				if err != nil {
				}
			}
		case err := <-fw.watch.Errors:
			if err != nil {

			}
			return
		case <-fw.done:
			_ = fw.watch.Close()
		}
	}
}
