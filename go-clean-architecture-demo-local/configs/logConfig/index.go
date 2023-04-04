package logConfig

import "log"

func Init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
