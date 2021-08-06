package folder_watcher

import (
	"bitbucket.org/Monitoring/gaemi/logging"
	"os"
	"path/filepath"
)

type FolderWatcher struct {
	inserter *FolderWatcherInserter
	path     string
}

func NewFolderWatcher(inserter *FolderWatcherInserter, dataPath string) *FolderWatcher {
	return &FolderWatcher{
		inserter: inserter,
		path:     dataPath,
	}
}

func (fw *FolderWatcher) Probe() {
	dirSize, err := getCurrentDirSize(fw.path)
	check(err)

	fw.inserter.WriteFolderStatus("tag", dirSize)
	logging.GetLogger().Tracef("current folder size is : %d", dirSize)
}

func check(e error) {
	if e != nil {
		logging.GetLogger().Panicf("Panic!! %s", e.Error())
		panic(e)
	}
}

func getCurrentDirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
