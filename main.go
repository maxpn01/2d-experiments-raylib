package main

import (
	"image/color"

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
	title:   "2D sandbox",
	bgColor: rl.Black,
}

var windowCenter = rl.NewVector2(float32(window.width)/2, float32(window.height)/2)

/* Game entities */
var player = NewPlayer(rl.NewVector2(300, 300), rl.NewVector2(50, 50), rl.Red, 400, 100, 100)
var fruitSpawner = NewFruitSpawner(rl.NewVector2(30, 30), rl.Yellow, 1, 20)
var snake = NewSnake(rl.NewVector2(windowCenter.X, windowCenter.Y), rl.NewVector2(50, 50), rl.Green, 200, 100, 100)

var entities = []GameObject{player, fruitSpawner, snake}

/* HUD entities */
var playerHpText = NewHUDText(rl.NewVector2(30, 30), 24, rl.RayWhite, "hp:", func() float32 { return player.hp })
var playerLvlText = NewHUDText(rl.NewVector2(130, 30), 24, rl.RayWhite, "lvl:", func() float32 { return float32(player.lvl) })
var playerExpText = NewHUDText(rl.NewVector2(230, 30), 24, rl.RayWhite, "exp:", func() float32 { return player.exp })
var playerSpeedText = NewHUDText(rl.NewVector2(30, 60), 22, rl.RayWhite, "speed:", func() float32 { return player.speed })

var snakeHpText = NewHUDText(rl.NewVector2(float32(window.width-190), 30), 24, rl.RayWhite, "snake hp:", func() float32 { return snake.hp })
var snakeLvlText = NewHUDText(rl.NewVector2(float32(window.width-380), 30), 24, rl.RayWhite, "snake lvl:", func() float32 { return float32(snake.lvl) })
var snakeExpText = NewHUDText(rl.NewVector2(float32(window.width-590), 30), 24, rl.RayWhite, "snake exp:", func() float32 { return snake.exp })

var hud = []GameObject{
	playerHpText,
	playerLvlText,
	playerExpText,
	playerSpeedText,
	snakeHpText,
	snakeLvlText,
	snakeExpText,
}

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
