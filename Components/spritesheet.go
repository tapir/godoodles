package main

import al "github.com/tapir/allegro5"

type SpriteSheet struct {
	Image *al.Bitmap
	Images []*al.Bitmap
}

func NewSpriteSheet(file string, width, height int, offsetX, offsetY float32) (*SpriteSheer, err) {
	b := al.LoadBitmap(file)
	if b == nil {
		return nil, errors.New("SPRITESHEET: Can't load file ", file)
	}

	w := b.GetWidth() - offsetX
	h := b.GetHeight() - offsetY

	if w % width != 0 || h % height != 0 {
		return nil, errors.New("SPRITESHEET: Wrong parameters")
	}

	nx :=  w / width
	ny := h / height
	size := nx * ny
	images = make([]*al.Bitmap, size, size)

	for i := 0; i < size; i++ {
		alx := offsetX + (size % nx) * width
		aly := offsetY + (size / nx) * height
		images[index] = b.CreateSub(alx, aly, width, height)
		if images[index] == nil {
			return nil, errors.New("SPRITESHEET: Can't create sub bitmap")
		}
	}

	return &SpriteSheet{b, images}, nil
}