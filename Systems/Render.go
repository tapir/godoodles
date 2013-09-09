package main

type RenderSystem struct {

}

func (r *RenderSystem) Update(e *Entity, userData interface{}) {
	sprite := e.GetComponent("Sprite")
	spriteSheet := e.GetComponent("SpriteSheet")

	if sprite != nil {

	}

	if spriteSheet != nil {
		
	}
}