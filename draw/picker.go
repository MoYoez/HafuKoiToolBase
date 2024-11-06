package draw

import (
	"image"
	"math"
)

/*
	图片取色算法
	这个算法是用来获取图片的平均颜色，并且进行调整的。

	当时这个算法是为了给背景图取色，所以说只采取了一部分（右上角），因为当时时间显示就在那边
	可以根据需要修改

*/

// GetAverageBGColor different from k-means algorithm,it uses origin plugin's algorithm.(Reduce the cost of averge color usage.)
func GetAverageBGColor(image image.Image) (int, int, int) {
	var RList []int
	var GList []int
	var BList []int
	width, height := image.Bounds().Size().X, image.Bounds().Size().Y
	// use the center of the bg, to make it more quickly and save memory and usage.
	for x := int(math.Round(float64(width) / 1.5)); x < int(math.Round(float64(width))); x++ {
		for y := height / 10; y < height/2; y++ {
			r, g, b, _ := image.At(x, y).RGBA()
			RList = append(RList, int(r>>8))
			GList = append(GList, int(g>>8))
			BList = append(BList, int(b>>8))
		}
	}
	RAverage := int(Average(RList))
	GAverage := int(Average(GList))
	BAverage := int(Average(BList))
	return RAverage, GAverage, BAverage
}

// Average sum all the numbers and divide by the length of the list.
func Average(numbers []int) float64 {
	var sum float64
	for _, num := range numbers {
		sum += float64(num)
	}
	return math.Round(sum / float64(len(numbers)))
}
