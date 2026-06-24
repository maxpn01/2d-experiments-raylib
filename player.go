package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	pos   rl.Vector2
	size  rl.Vector2
	speed float32
	color color.RGBA
}

func (p *Player) handlePlayerMovement(window *Window, dt float32) {
	move := rl.Vector2{}

	// wasd movement
	if rl.IsKeyDown(rl.KeyW) {
		move.Y -= 1
	}
	if rl.IsKeyDown(rl.KeyS) {
		move.Y += 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		move.X -= 1
	}
	if rl.IsKeyDown(rl.KeyD) {
		move.X += 1
	}

	// shift speed
	speed := p.speed
	if rl.IsKeyDown(rl.KeyLeftShift) {
		speed *= 2
	}

	// normalize the diagonal speed
	if move.X != 0 || move.Y != 0 {
		move = rl.Vector2Normalize(move)

		p.pos.X += move.X * speed * dt
		p.pos.Y += move.Y * speed * dt
	}

	// check for edges
	if p.pos.X < 0 {
		p.pos.X = 0
	}
	if p.pos.Y < 0 {
		p.pos.Y = 0
	}
	if p.pos.X+p.size.X > float32(window.width) {
		p.pos.X = float32(window.width) - p.size.X
	}
	if p.pos.Y+p.size.Y > float32(window.height) {
		p.pos.Y = float32(window.height) - p.size.Y
	}
}
