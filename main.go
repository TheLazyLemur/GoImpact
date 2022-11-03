package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
	Speed  float32
}

func (p *Player) Update() {
	if rl.IsKeyDown(rl.KeyRight) {
		p.X += p.Speed * rl.GetFrameTime()
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		p.X -= p.Speed * rl.GetFrameTime()
	}

	if rl.IsKeyDown(rl.KeyUp) {
		p.Y -= p.Speed * rl.GetFrameTime()
	}

	if rl.IsKeyDown(rl.KeyDown) {
		p.Y += p.Speed * rl.GetFrameTime()
	}

	if p.X < 0 {
		p.X = 0
	}

	if p.X+p.Width > float32(rl.GetScreenWidth()) {
		p.X = float32(rl.GetScreenWidth()) - p.Width
	}

	if p.Y < 0 {
		p.Y = 0
	}

	if p.Y+p.Height > float32(rl.GetScreenHeight()) {
		p.Y = float32(rl.GetScreenHeight()) - p.Height
	}

	playerRec := rl.Rectangle{
		X:      p.X,
		Y:      p.Y,
		Width:  p.Width,
		Height: p.Height,
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		NewBullet(playerRec.X, playerRec.Y+p.Height/2-5)
	}

	rl.DrawRectangleRec(playerRec, rl.Blue)
}

type Enemy struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type Enemies []Enemy

func (b Enemies) Update() {
	for i := range b {
		ele := &b[i]

		rec := rl.Rectangle{
			X:      ele.X,
			Y:      ele.Y,
			Width:  ele.Width,
			Height: ele.Height,
		}
		rl.DrawRectangleRec(rec, rl.Green)
	}
}

type Bullet struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	Speed    float32
	Lifetime float32
}

func NewBullet(x float32, y float32) {
	b := Bullet{
		X:        x,
		Y:        y,
		Width:    5,
		Height:   5,
		Speed:    300,
		Lifetime: 5,
	}
	bullets = append(bullets, b)
}

type Bullets []Bullet

func (b Bullets) Update() {
	for i := range b {
		ele := &b[i]

		ele.X += ele.Speed * rl.GetFrameTime()

		rec := rl.Rectangle{
			X:      ele.X,
			Y:      ele.Y,
			Width:  ele.Width,
			Height: ele.Height,
		}

		rl.DrawRectangleRec(rec, rl.Red)

		ele.Lifetime -= rl.GetFrameTime()

		if ele.Lifetime <= 0 {
			bullets = remove(bullets, i)
		}
	}
}

var bullets Bullets
var enemies Enemies

func main() {
	rl.InitWindow(800, 450, "Go Impact - Space Impact but in Go")
	rl.SetTargetFPS(60)
	enemies = append(enemies, Enemy{
		X:      50,
		Y:      0,
		Height: 50,
		Width:  50,
	})
	player := Player{
		X:      0,
		Y:      0,
		Width:  20,
		Height: 20,
		Speed:  300,
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.DrawFPS(10, 10)

		player.Update()
		bullets.Update()
		enemies.Update()
		UpdateCollisions()

		rl.ClearBackground(rl.Black)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func UpdateCollisions() {
	for _, enemy := range enemies {
		for _, bullet := range bullets {
			enemyRec := rl.NewRectangle(enemy.X, enemy.Y, enemy.Height, enemy.Width)
			bulletRec := rl.NewRectangle(bullet.X, bullet.Y, bullet.Height, bullet.Width)

			if rl.CheckCollisionRecs(enemyRec, bulletRec) {
				fmt.Println("BANG!!!!")
			}
		}
	}
}

func remove(slice []Bullet, s int) []Bullet {
	return append(slice[:s], slice[s+1:]...)
}
