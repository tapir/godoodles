package main

import al "github.com/tapir/allegro5"

type Sprite struct {
	Image *al.Bitmap
}

func NewSprite(file string) (*Sprite, err) {
	b := al.LoadBitmap(file)
	if b == nil {
		return nil, errors.New("SPRITE: Can't load file ", file)
	}
	return &Sprite{b}, nil
}