package main

import (
	"image/color"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Window struct {
	width   int32
	height  int32
	title   string
	bgColor color.RGBA
}

type Text struct {
	text  string
	pos   rl.Vector2
	font  int32
	color color.RGBA
}

var window = &Window{
	width:   1280,
	height:  720,
	title:   "2d game",
	bgColor: rl.Black,
}

var title = &Text{
	text:  "2D game",
	pos:   rl.NewVector2(20, 20),
	font:  24,
	color: rl.RayWhite,
}

var player = &Player{
	pos:   rl.NewVector2(300, 300),
	size:  rl.NewVector2(30, 30),
	speed: 400,
	color: rl.Red,
}

var fruitSpawner = &FruitSpawner{
	fruits:             []Fruit{},
	fruitSize:          20,
	fruitSpawnTimer:    0,
	fruitSpawnInterval: 1 + rand.Intn(20),
	maxFruits:          20,
}

func main() {
	rl.InitWindow(window.width, window.height, window.title)
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		player.handlePlayerMovement(window, dt)
		fruitSpawner.spawnFruit(window, dt)

		rl.BeginDrawing()

		rl.ClearBackground(window.bgColor)

		rl.DrawRectangle(int32(player.pos.X), int32(player.pos.Y), int32(player.size.X), int32(player.size.Y), player.color)

		for _, fruit := range fruitSpawner.fruits {
			rl.DrawRectangle(int32(fruit.pos.X), int32(fruit.pos.Y), int32(fruit.size.X), int32(fruit.size.Y), fruit.color)
		}

		rl.DrawText(title.text, int32(title.pos.X), int32(title.pos.Y), title.font, title.color)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
