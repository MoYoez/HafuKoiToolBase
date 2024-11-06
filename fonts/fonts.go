package fonts

import (
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

/*
	在一些工具 （我说的就是那 GG），OpenType 渲染有问题，直接导入字体会出错，所以对于OTF字体采取此方法
*/

// LoadFontFace load font face once before running, to work it quickly and save memory.
func LoadFontFace(filePath string, size float64) font.Face {
	fontFile, _ := os.ReadFile(filePath)
	fontFileParse, _ := opentype.Parse(fontFile)
	fontFace, _ := opentype.NewFace(fontFileParse, &opentype.FaceOptions{Size: size, DPI: 72, Hinting: font.HintingFull})
	return fontFace
}
