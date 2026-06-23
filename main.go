package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Window struct {
	width  int32
	height int32
	title  string
}

type Player struct {
	pos   rl.Vector2
	size  rl.Vector2
	speed float32
}

func main() {
	window := &Window{
		width:  1280,
		height: 720,
		title:  "2d game",
	}

	rl.InitWindow(window.width, window.height, window.title)

	player := &Player{
		pos:   rl.NewVector2(300, 300),
		size:  rl.NewVector2(50, 50),
		speed: 400,
	}

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		rl.SetTargetFPS(60)

		move := rl.Vector2{}

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

		speed := player.speed
		if rl.IsKeyDown(rl.KeyLeftShift) {
			speed *= 2
		}

		if move.X != 0 || move.Y != 0 {
			move = rl.Vector2Normalize(move)

			player.pos.X += move.X * speed * dt
			player.pos.Y += move.Y * speed * dt
		}

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

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawRectangle(int32(player.pos.X), int32(player.pos.Y), int32(player.size.X), int32(player.size.Y), rl.Red)
		rl.DrawText("2D game", 20, 20, 24, rl.RayWhite)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
