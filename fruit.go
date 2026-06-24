package main

import (
	"image/color"
	"log"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Fruit struct {
	pos   rl.Vector2
	size  rl.Vector2
	color color.RGBA
}

type FruitSpawner struct {
	fruits             []Fruit
	fruitSize          int
	fruitSpawnTimer    float32
	fruitSpawnInterval int
	maxFruits          int
}

func (fs *FruitSpawner) spawnFruit(window *Window, dt float32) {
	fs.fruitSpawnTimer += dt

	if fs.fruitSpawnTimer >= float32(fs.fruitSpawnInterval) && len(fs.fruits) < fs.maxFruits {
		fruitRandX := rand.Intn(int(window.width) - fs.fruitSize)
		fruitRandY := rand.Intn(int(window.height) - fs.fruitSize)

		fruit := Fruit{
			pos:   rl.NewVector2(float32(fruitRandX), float32(fruitRandY)),
			size:  rl.NewVector2(float32(fs.fruitSize), float32(fs.fruitSize)),
			color: rl.Green,
		}

		log.Printf("fruit spawned (x: %f, y: %f)", fruit.pos.X, fruit.pos.Y)

		fs.fruits = append(fs.fruits, fruit)

		fs.fruitSpawnTimer = 0
		fs.fruitSpawnInterval = 1 + rand.Intn(20)
	}
}
