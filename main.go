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

var fruitSpawner = &FruitSpawner{
	fruits:             []Fruit{},
	fruitSize:          30,
	fruitColor:         rl.Yellow,
	fruitSpawnTimer:    0,
	fruitSpawnInterval: 1 + rand.Intn(fruitSpawnIntervalMaxSeconds),
	maxFruits:          20,
}

var snake = &Snake{
	pos:   rl.NewVector2(float32(window.width)/2, float32(window.height)/2),
	size:  rl.NewVector2(50, 50),
	color: rl.Green,
	speed: 400,
	hp:    1,
	maxHP: 100,
}

var playerHpText = &HPText{
	Text: Text{
		pos:   rl.NewVector2(30, 30),
		color: rl.RayWhite,
		font:  24,
	},
	label:    "hp:",
	hpSource: func() float32 { return player.hp },
}

var snakeHpText = &HPText{
	Text: Text{
		pos:   rl.NewVector2(float32(window.width)-190, 30),
		color: rl.RayWhite,
		font:  24,
	},
	label:    "snake hp:",
	hpSource: func() float32 { return snake.hp },
}

var entities = []GameObject{player, fruitSpawner, snake}
var hud = []GameObject{playerHpText, snakeHpText}

func main() {
	rl.InitWindow(window.width, window.height, window.title)
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		for _, e := range entities {
			e.update(dt)
		}

		for _, e := range hud {
			e.update(dt)
		}

		rl.BeginDrawing()
		rl.ClearBackground(window.bgColor)

		for _, e := range entities {
			e.draw()
		}

		for _, e := range hud {
			e.draw()
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
