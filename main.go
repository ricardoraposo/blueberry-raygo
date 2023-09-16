package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	width := int32(1200)
	heigth := int32(800)
	rl.InitWindow(width, heigth, "Blueberry Raygo")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	circleRadius := 2 * math.Pi
	amplitude := 200.0
	speed := 0.01
	angle := 0.0

	var pathPoints []rl.Vector2
	var pathReversedPoints []rl.Vector2
	for !rl.WindowShouldClose() {
		// circleX := int32(-2 * amplitude * math.Sin(angle))
		circleX := int32(amplitude * angle)
		circleY := int32(-amplitude * math.Sin(angle*7))
		// circleY := int32(math.Sin(float64(circleX)) * amplitude / 3)
		angle += speed

		rl.BeginDrawing()
		rl.ClearBackground(rl.LightGray)

		// Draws the x and y axis
		rl.DrawLineEx(rl.Vector2{X: float32(width / 2), Y: 0}, rl.Vector2{X: float32(width / 2), Y: float32(heigth)}, 2, rl.Black)
		rl.DrawLineEx(rl.Vector2{X: 0, Y: float32(heigth / 2)}, rl.Vector2{X: float32(width), Y: float32(heigth / 2)}, 2, rl.Black)

		// Draws the line that will follow the circle
		rls := rl.Vector2{X: float32(width/2 + circleX), Y: float32(heigth / 2)}
		rle := rl.Vector2{X: float32(width/2 + circleX), Y: float32(heigth/2 + circleY)}
		lls := rl.Vector2{X: float32(width/2 - circleX), Y: float32(heigth / 2)}
		lle := rl.Vector2{X: float32(width/2 - circleX), Y: float32(heigth/2 - circleY)}
		rl.DrawLineEx(rls, rle, 3, rl.Red)
		rl.DrawLineEx(lls, lle, 3, rl.Red)

    // Current position of the balls
		cp := rl.Vector2{X: float32(width/2 + circleX), Y: float32(heigth/2 + circleY)}
		crp := rl.Vector2{X: float32(width/2 - circleX), Y: float32(heigth/2 - circleY)}

		pathPoints = append(pathPoints, cp)
		pathReversedPoints = append(pathReversedPoints, crp)

		rl.DrawCircle(int32(cp.X), int32(cp.Y), float32(circleRadius), rl.Red)
		rl.DrawCircle(int32(crp.X), int32(crp.Y), float32(circleRadius), rl.Red)

		for i := 1; i < len(pathPoints); i++ {
			rl.DrawLineEx(pathPoints[i-1], pathPoints[i], 4, rl.Black)
			rl.DrawLineEx(pathReversedPoints[i-1], pathReversedPoints[i], 4, rl.Black)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
