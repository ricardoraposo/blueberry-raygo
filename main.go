package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Setup window
	width := int32(1200)
	heigth := int32(800)
	rl.InitWindow(width, heigth, "Blueberry Raygo")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// Camera stuff
	initP := rl.Vector2{X: float32(width) / 2, Y: float32(heigth) / 2}
	camera := rl.NewCamera2D(initP, initP, 0, 1.0)

	// Music stuff
	rl.InitAudioDevice()
	var music rl.Music

	// Wave stuff
	var wave rl.Wave
	var allSamples []float32

	circleRadius := 2 * math.Pi
	amplitude := 200.0
	acc := 0.01
	speed := acc
	angle := 0.0

	var pathPoints []rl.Vector2
	// var pathReversedPoints []rl.Vector2
	font := rl.LoadFont("pragmasevka-regular.ttf")

	average := float32(0)
	step := 0
	i := 0
	// Event loop start
	for !rl.WindowShouldClose() {
		if rl.IsFileDropped() {
			droppedFile := rl.LoadDroppedFiles()
			music = rl.LoadMusicStream(droppedFile[0])
			rl.PlayMusicStream(music)
			wave = rl.LoadWave(droppedFile[0])
			allSamples = rl.LoadWaveSamples(wave)
			step = len(allSamples) / 1024
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			if rl.IsMusicStreamPlaying(music) {
				rl.PauseMusicStream(music)
				speed = 0
				step = 0
			} else {
				rl.ResumeMusicStream(music)
				step = len(allSamples) / 1024
				speed = acc
			}
		}

		if rl.IsMusicStreamPlaying(music) {
			rl.UpdateMusicStream(music)

			slice := allSamples[i : i+step]

			sum := float32(0.0)
			for k := 0; k < len(slice); k++ {
				sum = sum + slice[k]
			}
			average = sum / float32(len(slice))
			i = i + step

			// Drawing stuff
			rl.BeginDrawing()
			rl.ClearBackground(rl.LightGray)
			circleX := int32(amplitude * angle)
			circleY := average
			angle += speed
			// circleX := int32(2 * amplitude * math.Sin(angle))
			// circleY := int32(-amplitude * math.Sin(angle*10))
			// rls := rl.Vector2{X: float32(width / 2), Y: float32(heigth / 2)}
			// rle := rl.Vector2{X: float32(width / 2), Y: float32(heigth / 2)}
			// rl.DrawLineEx(rls, rle, 3, rl.Red)

			cp := rl.Vector2{X: float32(width/2 + circleX), Y: float32(heigth/2) * circleY}
			// crp := rl.Vector2{X: float32(width/2 - circleX), Y: float32(heigth/2) - circleY}

			rl.BeginMode2D(camera)
			camera.Target = cp

			pathPoints = append(pathPoints, cp)
			// pathReversedPoints = append(pathReversedPoints, crp)

			rl.DrawCircle(int32(cp.X), int32(cp.Y), float32(circleRadius), rl.Red)
			// rl.DrawCircle(int32(crp.X), int32(crp.Y), float32(circleRadius), rl.Red)

			for i := 1; i < len(pathPoints); i++ {
				rl.DrawLineEx(pathPoints[i-1], pathPoints[i], 4, rl.Black)
				// rl.DrawLineEx(pathReversedPoints[i-1], pathReversedPoints[i], 4, rl.Black)
			}
		} else {
			rl.ClearBackground(rl.LightGray)
			text := "Drag and drop files to play"
			rl.DrawTextEx(
				font,
				text,
				rl.Vector2{X: float32(width/2 - int32(len(text)/2)*16), Y: float32(heigth/2) - 16},
				32.0,
				1.0,
				rl.Black)
		}

		// Draws the x and y axis
		// rl.DrawLineEx(rl.Vector2{X: float32(width / 2), Y: 0}, rl.Vector2{X: float32(width / 2), Y: float32(heigth)}, 2, rl.Black)
		// rl.DrawLineEx(rl.Vector2{X: 0, Y: float32(heigth / 2)}, rl.Vector2{X: float32(width), Y: float32(heigth / 2)}, 2, rl.Black)
		// Draws the line that will follow the circle
		// lls := rl.Vector2{X: float32(width/2 - circleX), Y: float32(heigth / 2)}
		// lle := rl.Vector2{X: float32(width/2 - circleX), Y: float32(heigth/2 - circleY)}
		// rl.DrawLineEx(lls, lle, 3, rl.Red)

		rl.EndDrawing()
	}

	rl.CloseAudioDevice()
	rl.CloseWindow()
}
