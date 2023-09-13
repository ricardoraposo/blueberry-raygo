package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	width := int32(800)
	heigth := int32(600)
	rl.InitWindow(width, heigth, "Blueberry Raygo")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	circleRadius := 3 * math.Pi
	amplitude := 200.0
	speed := 0.04
	angle := 0.0

	for !rl.WindowShouldClose() {
		circleX := int32(1.5 * amplitude * math.Cos(angle))
		circleY := int32(amplitude * math.Cos(angle*3))
		angle += speed
		rl.BeginDrawing()
		rl.ClearBackground(rl.LightGray)
		rl.DrawLineEx(rl.Vector2{X: 400, Y: 0}, rl.Vector2{X: 400, Y: 600}, 2, rl.Black)
		rl.DrawLineEx(rl.Vector2{X: 0, Y: 300}, rl.Vector2{X: 800, Y: 300}, 2, rl.Black)
		start := rl.Vector2{X: float32(400 + circleX), Y: 300}
		end := rl.Vector2{X: float32(400 + circleX), Y: float32(300 + circleY)}
		rl.DrawLineEx(start, end, 10, rl.Red)
		rl.DrawCircle(400+circleX, 300+circleY, float32(circleRadius), rl.Red)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
