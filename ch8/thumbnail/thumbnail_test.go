package main

import "log"

//!+1
// makeThumbnails makes thumbnails of the specified files.
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

//!-1

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go ImageFile(f)
	}
}
