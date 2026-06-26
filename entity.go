package main

type GameObject interface {
	update(dt float32)
	draw()
}
