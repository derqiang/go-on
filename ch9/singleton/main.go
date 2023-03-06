package main

import (
	"image"
	"sync"
)

var icons = make(map[string]image.Image)

func loadIcon() image.Image {

	return nil
}

var mu sync.RWMutex

func Icon(name string) image.Image {
	mu.RLock()
	if icons != nil {
		icon := icons[name]
		mu.RUnlock()
		return icon
	}
	mu.RUnlock()

	mu.Lock()
	if icons == nil {
		loadIcon()
	}
	icon := icons[name]
	mu.Unlock()
	return icon
}

var loadOnce sync.Once

func Icon1(name string) image.Image {
	loadOnce.Do(func() {
		loadIcon()
	})
	return icons[name]
}
