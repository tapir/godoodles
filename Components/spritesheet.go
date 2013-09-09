package main

import al "github.com/tapir/allegro5"

type PlayMode uint

const (
	Loop	PlayMode = iota
	Bounce
	Single
)

type AnimatedSprite struct {
	Image *al.Bitmap
	Images []*al.Bitmap
	CurrentImage *al.Bitmap
	Mode PlayMode
	Delay float32
}

func NewAnimatedSprite(file string, width, height int, offsetX, offsetY float32, mode PlayMode, delay float32) (*AnimatedSprite, err) {
	b := al.LoadBitmap(file)
	if b == nil {
		return nil, errors.New("AnimatedSprite: Can't load file ", file)
	}

	w := b.GetWidth() - offsetX
	h := b.GetHeight() - offsetY

	if w % width != 0 || h % height != 0 {
		return nil, errors.New("AnimatedSprite: Wrong parameters")
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
			return nil, errors.New("AnimatedSprite: Can't create sub bitmap")
		}
	}

	return &AnimatedSprite{b, images, images[0], mode, delay}, nil
}
