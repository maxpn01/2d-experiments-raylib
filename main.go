package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	pos   rl.Vector2
	speed float32
}

func main() {
	rl.InitWindow(1280, 720, "2d game")

	player := &Player{
		pos:   rl.NewVector2(300, 300),
		speed: 400,
	}

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

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

		if move.X != 0 || move.Y != 0 {
			move = rl.Vector2Normalize(move)
			player.pos.X += move.X * player.speed * dt
			player.pos.Y += move.Y * player.speed * dt
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawRectangle(int32(player.pos.X), int32(player.pos.Y), 50, 50, rl.Red)

		rl.DrawText("2D game", 20, 20, 24, rl.RayWhite)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
