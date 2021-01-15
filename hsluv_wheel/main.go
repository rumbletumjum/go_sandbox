package main

import (
	"fmt"
	"github.com/hsluv/hsluv-go"
	"math"
	"strconv"
	"strings"
)

const (
	SATURATION = 100
	LIGHTNESS  = 75
	OCTOTHORPE = "#"
	reset      = "\033[0m"
)

var (
	fgRed = fgColor(1)
)

func fgColor(color int) string {
	return fmt.Sprintf("\033[3%dm", color)
}
func fgRgb(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func cleanHex(s string) string {
	return strings.Trim(s, OCTOTHORPE)
}

func main() {
	//start := 0.0
	//count := 6.0
	//del := 360.0 / count

	//stops := []string{"red", "yellow", "green", "blue", "magenta", "cyan"}
	//values := make(map[string]float64)
	//
	//for i, stop := range stops {
	//	values[stop] = float64(i) * 60.0
	//}
	//
	//for _, hue := range values {
	//	hex := cleanHex(hsluv.HsluvToHex(hue, SATURATION, LIGHTNESS))
	//	fmt.Println(hex)
	//}
	//
	//fmt.Printf("%dm\n", 7)

	//for hue := start; hue < 360.0; hue += del {
	//	light := hsluv.HpluvToHex(hue, SATURATION, LIGHTNESS)
	//	dark := hsluv.HpluvToHex(hue, SATURATION, LIGHTNESS - 20)
	//	//fmt.Printf("Hue %f: %s\n", hue, strings.Trim(light, "#"))
	//	fmt.Println(strings.Trim(light, "#"))
	//	fmt.Println(strings.Trim(dark, "#"))
	//}

	//fmt.Println(fgRed + "██████████" + reset)
	//fmt.Println(fgRgb(40, 100, 100) + "██████████" + reset)
	h := 243.7
	s := 61.4
	l := 61.6
	r, g, b := hsluv.HsluvToRGB(h, s, l)
	fmt.Printf("(%f, %f, %f)\n", h, s, l)
	fmt.Printf("(%f, %f, %f)\n", r, g, b)

	rInt := math.Round(r * 255.0)
	gInt := math.Round(g * 255.0)
	bInt := math.Round(b * 255.0)

	fmt.Printf("(%f, %f, %f)\n", rInt, gInt, bInt)
	fmt.Printf("(%d, %d, %d)\n", int(rInt), int(gInt), int(bInt))

	rr := strconv.FormatInt(int64(math.Round(rInt)), 16)
	gg := strconv.FormatInt(int64(math.Round(gInt)), 16)
	bb := strconv.FormatInt(int64(math.Round(bInt)), 16)

	fmt.Println(rr + gg + bb)

}
