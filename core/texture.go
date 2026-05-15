package core

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func LoadImage(path string, sampling uint32) uint32 {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(PrefixWarn + "Couldn't find or open image: " + path)
	}

	defer file.Close()

	imgInfo, format, err := image.DecodeConfig(file)
	if err != nil {
		log.Fatal(PrefixWarn + "Couldn't decode image!")
	}

	file.Seek(0, 0)

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(PrefixWarn + "Couldn't decode image!")
	}

	if format != "png" && format != "jpeg" {
		log.Fatal(PrefixWarn + "Unsupported image format must be PNG or JPEG!")
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba == nil {
		log.Fatal(PrefixWarn + "Couldn't create RGBA image!")
	}

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	bounds := rgba.Bounds()
	height := bounds.Dy()
	stride := rgba.Stride

	tempRow := make([]uint8, stride)

	for y := 0; y < height/2; y++ {
		topOffset := y * stride
		bottomOffset := (height - 1 - y) * stride

		copy(tempRow, rgba.Pix[topOffset:topOffset+stride])
		copy(rgba.Pix[topOffset:topOffset+stride], rgba.Pix[bottomOffset:bottomOffset+stride])
		copy(rgba.Pix[bottomOffset:bottomOffset+stride], tempRow)
	}

	var texture uint32

	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, int32(sampling))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, int32(sampling))

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(imgInfo.Width), int32(imgInfo.Height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	return texture
}

func flipV(rgba *image.RGBA) {
	bounds := rgba.Bounds()
	height := bounds.Dy()
	stride := rgba.Stride // Her bir yatay satırın byte cinsinden uzunluğu

	// Bir satırı geçici olarak tutmak için buffer oluştur
	tempRow := make([]uint8, stride)

	// Görselin yüksekliğinin yarısına kadar dönüp alt ve üst satırları yer değiştir
	for y := 0; y < height/2; y++ {
		topOffset := y * stride
		bottomOffset := (height - 1 - y) * stride

		// 1. Üst satırı geçici buffer'a kopyala
		copy(tempRow, rgba.Pix[topOffset:topOffset+stride])
		// 2. Alt satırı üst satırın yerine kopyala
		copy(rgba.Pix[topOffset:topOffset+stride], rgba.Pix[bottomOffset:bottomOffset+stride])
		// 3. Geçici buffer'daki (eski üst satır) veriyi alt satıra kopyala
		copy(rgba.Pix[bottomOffset:bottomOffset+stride], tempRow)
	}
}
