// 利萨茹曲线

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0x80, 0x00, 0x10},
	color.RGBA{0xff, 0xda, 0xb9, 0x10},
	color.RGBA{0xe6, 0x8a, 0xb8, 0x10},
	color.RGBA{0x87, 0xce, 0xfa, 0x10},
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 完整的 x 振荡器变化的个数
		res     = 0.001 // 角度分辨率
		size    = 100   // 图像画布包含尺寸大小 [-size - +size]
		nframes = 64    // 动画中的帧数
		delay   = 8     // 以 10ms 为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0 // y 振荡器相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			idx := randInt(1, 4)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(idx))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func randInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return (rand.Intn(int(max-min)) + min)
}
