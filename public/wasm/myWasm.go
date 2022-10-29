package main

import (  
    "fmt"
	"syscall/js"
	"time"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	//"github.com/markfarnan/go-canvas/canvas"
	//"github.com/wawesomeNOGUI/mandlebrot/public/canvas"
)

//var cnvs js.Value = js.Global().Get("cnvs")
//var width float64 = cnvs.Get("width").Float()
//var height float64 = cnvs.Get("height").Float() 
//var ctx js.Value = js.Global().Get("ctx") 

var cvs *Canvas2d
var ctx *draw2dimg.GraphicContext
var width float64
var height float64
var renderDelay time.Duration = 20 * time.Millisecond

var realStart float64 = js.Global().Get("realStart").Float()
var realEnd float64 = js.Global().Get("realEnd").Float()
var imaginaryStart float64 = js.Global().Get("imaginaryStart").Float()
var imaginaryEnd float64 = js.Global().Get("imaginaryEnd").Float()
var res float64 = js.Global().Get("res").Float()
var maxIteration float64 = js.Global().Get("maxIteration").Float()

func drawMandlebrot(i float64, j float64) {
	x := float64(0)
	y := float64(0)
	xScaled := realStart + (i/width) * (realEnd - realStart)
	yScaled := imaginaryStart + (j/height) * (imaginaryEnd - imaginaryStart)

	iteration := float64(0);
	for x*x + y*y <= 2*2 && iteration < maxIteration {
		xTemp := x*x - y*y + xScaled
		y = 2*x*y + yScaled
		x = xTemp
		iteration++
	}

	if x*x + y*y >= 4 {
		//ctx.Set("fillStyle", "#FFFFFF")
		//ctx.Call("fillRect", i, j, res, res)
		ctx.SetFillColor(color.RGBA{0x00, 0x00, 0xFF, 0xFF})
		ctx.BeginPath()
		draw2dkit.Rectangle(ctx, i, j, i+res, j+res)
		ctx.Fill()
	} 
}

func Render(this js.Value, args []js.Value) any {
	realStart = js.Global().Get("realStart").Float()
	realEnd = js.Global().Get("realEnd").Float()
	imaginaryStart = js.Global().Get("imaginaryStart").Float()
	imaginaryEnd = js.Global().Get("imaginaryEnd").Float()
	res = js.Global().Get("res").Float()
	maxIteration = js.Global().Get("maxIteration").Float()

	ctx.SetFillColor(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})
	ctx.Clear()

	for i := float64(0); i < width; i+=res {
		for j := float64(0); j < height; j+=res {
			drawMandlebrot(i, j)
		}
	}
	cvs.imgCopy()

	fmt.Println("done")

	return true
}

func main() {  
    fmt.Println("Go Web Assembly")

	cvs, _ = NewCanvas2d(true)
	//cvs.Create(int(js.Global().Get("innerWidth").Float()), int(js.Global().Get("innerHeight").Float())) 

	ctx = cvs.Gc()

	height = float64(cvs.Height())
	width = float64(cvs.Width())

	js.Global().Call("requestAnimationFrame", js.FuncOf(Render))
	js.Global().Set("Render", js.FuncOf(Render))

	js.Global().Set("cvs", cvs.canvas)

	/*
	go func() {
		for {
			time.Sleep(10000)
			cvs.imgCopy()
		}
	}()
*/
	/*
	var animate js.Func
	//var i float64
	animate = js.FuncOf(func(this js.Value, args []js.Value) any {
		// for z := float64(0); z < 60; z++ {
		// 	i++
		// 	draw(i, 0)
		// }

		// if i < width {
		// 	js.Global().Call("requestAnimationFrame", animate)
		// } else {
		// 	i = 0
		// 	//j = 0
		// 	fmt.Println("yes")
		// }

		for z := float64(0); z < width; z += width / 5 {
			go draw(float64(int(z)), 0, z + width / 5)
			fmt.Println(float64(int(z)))
		}

		return true
	})

	go func() {
		for {
			time.Sleep(50)
			js.Global().Call("requestAnimationFrame", animate)
		}
	}()
	js.Global().Set("animate", animate)
	*/

	select {}  //so go program doesn't exit
}
