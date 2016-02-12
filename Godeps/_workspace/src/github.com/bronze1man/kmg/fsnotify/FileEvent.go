package fsnotify

import (
	"time"

	originFsnotify "github.com/howeyc/fsnotify"
)

//it is a wrapper of github.com/howeyc/fsnotify.FileEvent
type FileEvent struct {
	*originFsnotify.FileEvent
	Name string
	Time time.Time
}

func newFileEvent(originEvent *originFsnotify.FileEvent) *FileEvent {
	return &FileEvent{FileEvent: originEvent, Name: originEvent.Name, Time: time.Now()}
}
