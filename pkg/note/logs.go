package note

import (
	"io"
	"log"
)

// Change default log flags
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
	log.SetPrefix("\t")
}

// Option to suppress logs for command line usage
func SuppressLogs() {
	// Suppress logs for command line usage
	log.SetOutput(io.Discard)
}
