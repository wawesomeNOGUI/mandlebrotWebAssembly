package main

import (  
    "fmt"
	"syscall/js"
)

var cnvs js.Value = js.Global().Get("cnvs")
var width float64 = cnvs.Get("width").Float()
var height float64 = cnvs.Get("height").Float() 
var ctx js.Value = js.Global().Get("ctx") 

var realStart float64 = js.Global().Get("realStart").Float()
var realEnd float64 = js.Global().Get("realEnd").Float()
var imaginaryStart float64 = js.Global().Get("imaginaryStart").Float()
var imaginaryEnd float64 = js.Global().Get("imaginaryEnd").Float()
var res float64 = js.Global().Get("res").Float()
var maxIteration float64 = js.Global().Get("maxIteration").Float()

var i float64 = js.Global().Get("i").Float()
var j float64
func draw() {
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
		ctx.Set("fillStyle", "#FFFFFF")
		ctx.Call("fillRect", i, j, res, res)
	} 

	j += res;
    if(j > height) {
        j = 0;
        i += res;
		//js.Global().Set("i", i)
    }  
}

/*
func drawMandlebrot(this js.Value, args []js.Value) any {
	animate(this, args)

	return true
}
*/

func main() {  
    fmt.Println("Go Web Assembly")

	//js.Global().Set("drawMandlebrot", js.FuncOf(drawMandlebrot))
	var animate js.Func

	animate = js.FuncOf(func(this js.Value, args []js.Value) any {
		for z := float64(0); z < height/res*10; z++ {
			draw()
		}

		if i < width {
			js.Global().Call("requestAnimationFrame", animate)
		}

		return nil
	})

	js.Global().Call("requestAnimationFrame", animate)

	select {}  //so go program doesn't exit
}
