package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1280, 720, "2d game")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawRectangle(300, 300, 50, 50, rl.Red)

		rl.DrawText("2D game", 20, 20, 24, rl.RayWhite)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
