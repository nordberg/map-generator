package main

import (
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "math/rand"
    "os"
    "bufio"
)

func main() {
    var sizeX, sizeY int = 1000, 1000
    var islands int = 100

    matrix := make([][]int, sizeY)

    for i := range matrix {
        matrix[i] = make([]int, sizeX)
    }

    for i := range matrix {
        for j := range matrix[i] {
            matrix[i][j] = -15
        }
    }

    for i := 0; i < islands; i++ {
        posX := rand.Intn(sizeX)
        posY := rand.Intn(sizeY)
        r := rand.Intn(300)

        for j := posX - r; j < posX + r; j++ {
            if j < 0 {
                continue
            }

            if j > sizeX - 1 {
                continue
            }

            for k := posY - r; k < posY + r; k++ {
                if k < 0 {
                    continue
                }

                if k > sizeY - 1 {
                    continue
                }

                if ((j - posX) * (j - posX) + (k - posY) * (k - posY)) < r * r {
                    matrix[j][k] += rand.Intn(40)
                    if matrix[j][k] > 255 {
                        matrix[j][k] = 255
                    }
                }
            }
        }
    }

    for i := range matrix {
        if i == 0 || i == len(matrix) - 1 {
            continue
        }
        for j := range matrix[i] {
            if j == 0 || j == len(matrix[i]) - 1 {
                continue
            }

            sum := 0
            sum += matrix[i - 1][j]
            sum += matrix[i + 1][j]
            sum += matrix[i][j - 1]
            sum += matrix[i][j + 1]
            sum += matrix[i][j]
            avg := sum/5

            matrix[i][j] = int(avg)
        }
    }

    imgRect := image.Rect(0, 0, sizeX, sizeY)
    img := image.NewGray(imgRect)
    draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

    for y := 0; y < sizeY; y++ {
        for x := 0; x < sizeX; x++ {
            r := uint8(25)
            g := uint8(125)
            b := uint8(255)
            if matrix[x][y] > 127 {
                b += uint8(255 - matrix[x][y])
            }
            
            a := uint8(255)
            img.Set(x, y, color.RGBA{r, g, b, a})
        }
    }

    o, _ := os.Create("out.png")
    defer o.Close()
    writer := bufio.NewWriter(o)
    png.Encode(writer, img)
    writer.Flush()
    fmt.Println("DONE")
}