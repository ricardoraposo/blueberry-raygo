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

	circleRadius := 2 * math.Pi
	amplitude := 200.0
	speed := 0.01
	angle := 0.0

	var pathPoints []rl.Vector2
	for !rl.WindowShouldClose() {
		circleX := int32(1.5 * amplitude * math.Cos(angle))
		circleY := int32(amplitude * math.Cos(angle*3))
		angle += speed

		rl.BeginDrawing()
		rl.ClearBackground(rl.LightGray)

    // Draws the x and y axis
		rl.DrawLineEx(rl.Vector2{X: 400, Y: 0}, rl.Vector2{X: 400, Y: 600}, 2, rl.Black)
		rl.DrawLineEx(rl.Vector2{X: 0, Y: 300}, rl.Vector2{X: 800, Y: 300}, 2, rl.Black)

    // Draws the line that will follow the circle
		start := rl.Vector2{X: float32(400 + circleX), Y: 300}
		end := rl.Vector2{X: float32(400 + circleX), Y: float32(300 + circleY)}
		rl.DrawLineEx(start, end, 3, rl.Red)

		currentPosition := rl.Vector2{X: float32(400 + circleX), Y: float32(300 + circleY)}
    pathPoints = append(pathPoints, currentPosition)
		rl.DrawCircle(400+circleX, 300+circleY, float32(circleRadius), rl.Red)

    for i := 1; i < len(pathPoints); i++ {
			rl.DrawLineEx(pathPoints[i-1], pathPoints[i], 4, rl.Black)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
