package main

import (
	"fmt"
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

type GameObject interface {
	update(dt float32)
	draw()
}

var window = &Window{
	width:   1280,
	height:  720,
	title:   "2D game",
	bgColor: rl.Black,
}

var player = &Player{
	pos:   rl.NewVector2(300, 300),
	size:  rl.NewVector2(50, 50),
	color: rl.Red,
	speed: 400,
	hp:    1,
	maxHP: 100,
}

var hpText = &HPText{
	Text: Text{
		pos:   rl.NewVector2(30, 30),
		color: rl.RayWhite,
		text:  fmt.Sprintf("hp: %.2f", player.hp),
		font:  24,
	},
}

var fruitSpawner = &FruitSpawner{
	fruits:             []Fruit{},
	fruitSize:          30,
	fruitSpawnTimer:    0,
	fruitSpawnInterval: 1 + rand.Intn(20),
	maxFruits:          20,
}

var entities = []GameObject{player, hpText, fruitSpawner}

func main() {
	rl.InitWindow(window.width, window.height, window.title)
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		for _, e := range entities {
			e.update(dt)
		}

		rl.BeginDrawing()
		rl.ClearBackground(window.bgColor)

		for _, e := range entities {
			e.draw()
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func checkCollisions(pos1, size1, pos2, size2 rl.Vector2) bool {
	overlapX := pos1.X < pos2.X+size2.X && pos1.X+size1.X > pos2.X
	overlapY := pos1.Y < pos2.Y+size2.Y && pos1.Y+size1.Y > pos2.Y

	return overlapX && overlapY
}
