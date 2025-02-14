package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"math/rand"
)

var (
	Black = color.RGBA{0, 0, 0, 255}
	White = color.RGBA{255, 255, 255, 255}
)

// GUI
type Pixels struct {
	Pix   []uint8
	Width int
}

func NewPixels(width, height int) *Pixels { //NewPixels is a constructor
	return &Pixels{Width: width, Pix: make([]uint8, width*height*4)}
}

func (p *Pixels) DrawRect(x, y, width, height int, rgba color.RGBA) { //作用是在屏幕上画一个矩形
	for idx := 0; idx < width; idx++ {
		for idy := 0; idy < height; idy++ {
			p.SetColor(x+idx, y+idy, rgba)
		}
	}
}

func (p *Pixels) SetColor(x, y int, rgba color.RGBA) { // 作用是设置像素点的颜色
	r, g, b, a := rgba.RGBA()
	index := (y*p.Width + x) * 4
	p.Pix[index] = uint8(r)
	p.Pix[index+1] = uint8(g)
	p.Pix[index+2] = uint8(b)
	p.Pix[index+3] = uint8(a)
}

// Game

// 定义一个名为 CountNeighbors 的函数，接收一个二维整数数组 matrix，返回一个二维整数数组。
// 创建一个二维整数数组 neighbors，其行数与 matrix 相同
// 遍历 matrix 的每一行，为 neighbors 的每���行分配与 matrix 对应行相同长度的切片
func CountNeighbors(matrix [][]int) [][]int {
	neighbors := make([][]int, len(matrix))
	for idx, val := range matrix {
		neighbors[idx] = make([]int, len(val))
	}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			for rowMod := -1; rowMod < 2; rowMod++ {
				newRow := row + rowMod
				if newRow < 0 || newRow >= len(matrix) {
					continue
				}
				for colMod := -1; colMod < 2; colMod++ {
					if rowMod == 0 && colMod == 0 {
						continue
					}
					newCol := col + colMod
					if newCol < 0 || newCol >= len(matrix[row]) {
						continue
					}
					neighbors[row][col] += matrix[newRow][newCol]
				}
			}

		}
	}
	return neighbors
}

type GameOfLife struct {
	gameBoard [][]int
	pixels    *Pixels
	size      int
}

// 创建新的Game of life，空游戏版，以及一个用随机细胞填充游戏版的函数

func NewGameOfLife(width, height, size int) *GameOfLife {
	gameBoard := make([][]int, height)
	for idx := range gameBoard {
		gameBoard[idx] = make([]int, width)
	}
	pixels := NewPixels(width*size, height*size)
	return &GameOfLife{gameBoard: gameBoard, pixels: pixels, size: size}
}

func (gol *GameOfLife) Random() {
	for idy := range gol.gameBoard {
		for idx := range gol.gameBoard[idy] {
			gol.gameBoard[idy][idx] = rand.Intn(2)
		}
	}
}

func (gol *GameOfLife) PlayRound() {
	neighbors := CountNeighbors(gol.gameBoard)
	for idy := range gol.gameBoard {
		for idx, value := range gol.gameBoard[idy] {
			n := neighbors[idy][idx]

			if value == 1 && (n == 2 || n == 3) {
				continue
			} else if n == 3 {
				gol.gameBoard[idy][idx] = 1
				gol.pixels.DrawRect(idx*gol.size, idy*gol.size, gol.size, gol.size, Black)
			} else {
				gol.gameBoard[idy][idx] = 0
				gol.pixels.DrawRect(idx*gol.size, idy*gol.size, gol.size, gol.size, White)
			}
		}
	}
}

func run() {
	size := float64(2)
	width := 400
	height := 400
	cfg := pixelgl.WindowConfig{
		Title:  "Conway's Game of Life",
		Bounds: pixel.R(0, 0, float64(width)*size, float64(height)*size),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	gol := NewGameOfLife(width, height, int(size))
	gol.Random()
	for !win.Closed() {
		gol.PlayRound()
		win.Canvas().SetPixels(gol.pixels.Pix)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
