package fsnotify

import "github.com/fsnotify/fsnotify"

type (
	// Event represents a file system notification.
	// Name Path to the file or directory.
	Event = fsnotify.Event
)
