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
	size:  rl.NewVector2(30, 30),
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
	fruitSize:          20,
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
