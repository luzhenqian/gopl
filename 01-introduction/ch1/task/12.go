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
	"net/url"
	"strconv"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0x80, 0x00, 0x10},
	color.RGBA{0xff, 0xda, 0xb9, 0x10},
	color.RGBA{0xe6, 0x8a, 0xb8, 0x10},
	color.RGBA{0x87, 0xce, 0xfa, 0x10},
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	opt := defaultOption
	if err := r.ParseForm(); err == nil {
		assignOpt(&opt, r.Form)
	}
	lissajous(w, opt)
}

func assignOpt(old *LissajousOption, new url.Values) {
	for k, v := range new {
		if k == "cycles" {
			if v, err := strconv.ParseFloat(v[0], 64); err == nil {
				old.cycles = v
			}
			continue
		}
		if k == "res" {
			if v, err := strconv.ParseFloat(v[0], 64); err == nil {
				old.res = v
			}
			continue
		}
		if k == "size" {
			if v, err := strconv.ParseInt(v[0], 10, 64); err == nil {
				old.size = int(v)
			}
			continue
		}
		if k == "nframes" {
			if v, err := strconv.ParseInt(v[0], 10, 64); err == nil {
				old.nframes = int(v)
			}
			continue
		}
		if k == "delay" {
			if v, err := strconv.ParseInt(v[0], 10, 64); err == nil {
				old.delay = int(v)
			}
			continue
		}
	}
}

type LissajousOption struct {
	cycles  float64 // 完整的 x 振荡器变化的个数
	res     float64 // 角度分辨率
	size    int     // 图像画布包含尺寸大小 [-size - +size]
	nframes int     // 动画中的帧数
	delay   int     // 以 10ms 为单位的帧间延迟
}

var defaultOption = LissajousOption{
	cycles:  5,
	res:     0.001,
	size:    100,
	nframes: 64,
	delay:   8,
}

func lissajous(out io.Writer, opts ...LissajousOption) {
	opt := defaultOption
	if len(opts) > 0 {
		opt = opts[0]
	}
	freq := rand.Float64() * 3.0 // y 振荡器相对频率
	anim := gif.GIF{LoopCount: opt.nframes}
	phase := 0.0
	for i := 0; i < opt.nframes; i++ {
		rect := image.Rect(0, 0, 2*opt.size+1, 2*opt.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < opt.cycles*2*math.Pi; t += opt.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			idx := randInt(1, 4)
			img.SetColorIndex(opt.size+int(x*float64(opt.size)+0.5), opt.size+int(y*float64(opt.size)+0.5), uint8(idx))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, opt.delay)
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
