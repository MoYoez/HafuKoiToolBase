package draw

type RenderBorderColorSize struct {
	s                        string
	renderX, renderY, ax, ay float64
}

/*

	对字符串使用描边

	返回值为描边后的字符串坐标，用于后续的绘制。


*/

// DrawBorderString Draw the border of the string.
func DrawBorderString(s string, size int, x float64, y float64, ax float64, ay float64) []RenderBorderColorSize {
	PageContentBorder := make([]RenderBorderColorSize, 0)
	for dy := -size; dy <= size; dy++ {
		for dx := -size; dx <= size; dx++ {
			if dx*dx+dy*dy >= size*size {
				continue
			}
			renderX := x + float64(dx)
			renderY := y + float64(dy)
			PageContentBorder = append(PageContentBorder, RenderBorderColorSize{s, renderX, renderY, ax, ay})
		}
	}
	return PageContentBorder
}
