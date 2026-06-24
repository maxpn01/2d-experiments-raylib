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

const playerHPIncrement = 0.25

func (p *Player) update(dt float32) {
	handlePlayerMovement(window, p, dt)
	handlePlayerFruitCollision(p, fruitSpawner)
}

func (p *Player) draw() {
	rl.DrawRectangle(int32(p.pos.X), int32(p.pos.Y), int32(p.size.X), int32(p.size.Y), p.color)
}

func handlePlayerMovement(window *Window, player *Player, dt float32) {
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
	speed := player.speed
	if rl.IsKeyDown(rl.KeyLeftShift) {
		speed *= 2
	}

	// normalize the diagonal speed
	if move.X != 0 || move.Y != 0 {
		move = rl.Vector2Normalize(move)

		player.pos.X += move.X * speed * dt
		player.pos.Y += move.Y * speed * dt
	}

	// clamp to game window edges
	if player.pos.X < 0 {
		player.pos.X = 0
	}
	if player.pos.Y < 0 {
		player.pos.Y = 0
	}
	if player.pos.X+player.size.X > float32(window.width) {
		player.pos.X = float32(window.width) - player.size.X
	}
	if player.pos.Y+player.size.Y > float32(window.height) {
		player.pos.Y = float32(window.height) - player.size.Y
	}
}

func handlePlayerFruitCollision(player *Player, fs *FruitSpawner) {
	for i := len(fs.fruits) - 1; i >= 0; i-- {
		hasPlayerCollidedWithFruit := checkCollisions(player.pos, player.size, fs.fruits[i].pos, fs.fruits[i].size)

		if hasPlayerCollidedWithFruit && player.hp < player.maxHP {
			player.hp += playerHPIncrement
			fs.despawnFruit(i)
		}
	}
}
