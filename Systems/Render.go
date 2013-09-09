package main

type RenderSystem struct {

}

func (r *RenderSystem) Update(e *Entity, userData interface{}) {
	sprite := e.GetComponent("Sprite")
	animatedSprite := e.GetComponent("AnimatedSprite")

	if sprite != nil {

	}

	if animatedSprite != nil {
		
	}
}
