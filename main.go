package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	//原始图片是sam.jpg
	imgb, _ := os.Open("sam.jpg")
	img, _ := jpeg.Decode(imgb)
	defer imgb.Close()

	wmb, _ := os.Open("text.png")
	watermark, _ := png.Decode(wmb)
	defer wmb.Close()

	//把水印写到右下角，并向0坐标各偏移10个像素
	offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-10, img.Bounds().Dy()-watermark.Bounds().Dy()-10)
	b := img.Bounds()
	m := image.NewNRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	//生成新图片new.jpg，并设置图片质量..
	imgw, _ := os.Create("new.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{100})

	defer imgw.Close()

	fmt.Println("水印添加结束,请查看new.jpg图片...")
}

// package main

// import (
//     "fmt"
//     "image"
//     "image/color"
//     "image/jpeg"
//     "io/ioutil"
//     "log"
//     "os"

//     "github.com/golang/freetype"
// )

// func main() {
//     //需要加水印的图片
//     imgfile, _ := os.Open("u=488179422,3251067872&fm=200&gp=0.jpg")
//     defer imgfile.Close()

//     jpgimg, _ := jpeg.Decode(imgfile)

//     img := image.NewNRGBA(jpgimg.Bounds())

//     for y := 0; y < img.Bounds().Dy(); y++ {
//         for x := 0; x < img.Bounds().Dx(); x++ {
//             img.Set(x, y, jpgimg.At(x, y))
//         }
//     }
//     //拷贝一个字体文件到运行目录
//     fontBytes, err := ioutil.ReadFile("simsun.ttc")
//     if err != nil {
//         log.Println(err)
//     }

//     font, err := freetype.ParseFont(fontBytes)
//     if err != nil {
//         log.Println(err)
//     }

//     f := freetype.NewContext()
//     f.SetDPI(72)
//     f.SetFont(font)
//     f.SetFontSize(12)
//     f.SetClip(jpgimg.Bounds())
//     f.SetDst(img)
//     f.SetSrc(image.NewUniform(color.RGBA{R: 255, G: 0, B: 0, A: 255}))

//     pt := freetype.Pt(img.Bounds().Dx()-200, img.Bounds().Dy()-12)
//     _, err = f.DrawString("中文 string 255.43,232.12312 老纪", pt)

//     //draw.Draw(img,jpgimg.Bounds(),jpgimg,image.ZP,draw.Over)

//     //保存到新文件中
//     newfile, _ := os.Create("aaa.jpg")
//     defer newfile.Close()

//     err = jpeg.Encode(newfile, img, &jpeg.Options{100})
//     if err != nil {
//         fmt.Println(err)
//     }
// }
