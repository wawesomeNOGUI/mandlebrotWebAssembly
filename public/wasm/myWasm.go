package main

import (  
    "fmt"
	"syscall/js"
)

var cnvs js.Value = js.Global().Get("cnvs")
var width float64 = cnvs.Get("width")
var height float64 = cnvs.Get("height") 
var ctx js.Value = js.Global().Get("ctx") 

var realStart float64 = js.Global().Get("realStart").Float()
var realEnd float64 = js.Global().Get("realEnd").Float()
var imaginaryStart float64 = js.Global().Get("imaginaryStart").Float()
var imaginaryEnd float64 = js.Global().Get("imaginaryEnd").Float()
var res float64 = js.Global().Get("res").Float()
var maxIteration float64 = js.Global().Get("maxIteration").Float()

func drawMandlebrot(this js.Value, args []js.Value) {
	for i := 0; i < width; i += res {
		for j := 0; j < height; j += res {
			x := 0
			y := 0
			xScaled := realStart + (i/width) * (realEnd - realStart)
			yScaled := imaginaryStart + (j/height) * (imaginaryEnd - imaginaryStart)
	
			iteration := 0;
			for x*x + y*y <= 2*2 && iteration < maxIteration {
				xTemp := x*x - y*y + xScaled
				y = 2*x*y + yScaled
				x = xTemp
				iteration++
			}
	
			if x*x + y*y > 4.5 {
				ctx.Set("fillStyle", "#FFFFFF")
				ctx.Call("fillRect", i, j, res, res)
			} 
		}
	}
}

func main() {  
    fmt.Println("Go Web Assembly")

	js.Global().Set("drawMandlebrot", js.FuncOf(drawMandlebrot))

	select {}  //so go program doesn't exit
}
