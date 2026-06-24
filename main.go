package main

import (
	"image/color"
	"log"
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

type Player struct {
	pos   rl.Vector2
	size  rl.Vector2
	speed float32
	color color.RGBA
}

type Fruit struct {
	pos   rl.Vector2
	size  rl.Vector2
	color color.RGBA
}

const maxFruits = 20

func main() {
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

	var fruits = []Fruit{}
	var fruitSpawnTimer float32
	var fruitSpawnInterval = 1 + rand.Intn(20)

	rl.InitWindow(window.width, window.height, window.title)
	rl.SetTargetFPS(60)

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

		fruitSpawnTimer += dt

		if fruitSpawnTimer >= float32(fruitSpawnInterval) && len(fruits) < maxFruits {
			fruitSize := 20
			fruitRandX := rand.Intn(int(window.width) - fruitSize)
			fruitRandY := rand.Intn(int(window.height) - fruitSize)

			fruit := &Fruit{
				pos:   rl.NewVector2(float32(fruitRandX), float32(fruitRandY)),
				size:  rl.NewVector2(float32(fruitSize), float32(fruitSize)),
				color: rl.Green,
			}

			log.Printf("fruit spawned: x %f y %f", fruit.pos.X, fruit.pos.Y)

			fruits = append(fruits, *fruit)

			fruitSpawnTimer = 0
			fruitSpawnInterval = 1 + rand.Intn(20)
		}

		rl.BeginDrawing()
		rl.ClearBackground(window.bgColor)
		rl.DrawRectangle(int32(player.pos.X), int32(player.pos.Y), int32(player.size.X), int32(player.size.Y), player.color)
		for _, fruit := range fruits {
			rl.DrawRectangle(int32(fruit.pos.X), int32(fruit.pos.Y), int32(fruit.size.X), int32(fruit.size.Y), fruit.color)
		}
		rl.DrawText(title.text, int32(title.pos.X), int32(title.pos.Y), title.font, title.color)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
