package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	pos   rl.Vector2
	size  rl.Vector2
	color color.RGBA

	speed float32
	hp    float32
	maxHP float32
}

const playerHpIncrement = 0.25

func (p *Player) update(dt float32) {
	p.movePlayer(window, dt)
	p.handlePlayerFruitCollision(fruitSpawner)
}

func (p *Player) draw() {
	rl.DrawRectangle(int32(p.pos.X), int32(p.pos.Y), int32(p.size.X), int32(p.size.Y), p.color)
}

func (p *Player) movePlayer(window *Window, dt float32) {
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

	// clamp to game window edges
	clamp(0, &p.pos.X, &p.size.X, float32(window.width))
	clamp(0, &p.pos.Y, &p.size.Y, float32(window.height))
}

func (p *Player) handlePlayerFruitCollision(fs *FruitSpawner) {
	for i := len(fs.fruits) - 1; i >= 0; i-- {
		hasPlayerCollidedWithFruit := checkCollisions(p.pos, p.size, fs.fruits[i].pos, fs.fruits[i].size)

		if hasPlayerCollidedWithFruit && p.hp < p.maxHP {
			p.hp += playerHpIncrement
			fs.despawnFruit(i)
		}
	}
}
