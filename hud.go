package main

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
	pos   rl.Vector2
	color color.RGBA

	text string
	font int32
}

func (t *Text) update(dt float32) {}

func (t *Text) draw() {
	rl.DrawText(t.text, int32(t.pos.X), int32(t.pos.Y), t.font, t.color)
}

type HPText struct {
	Text
	label    string
	hpSource func() float32
}

func (h *HPText) update(dt float32) {
	h.text = fmt.Sprintf("%s %.2f", h.label, h.hpSource())
}
