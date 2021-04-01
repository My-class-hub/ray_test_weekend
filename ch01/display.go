package display

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Display struct {
	Width  int    //长
	Height int    //宽
	Title  string //标题
}

var displayImag Display

func PathURL() (result string) {
	nowTime := time.Now().Format("15_04_05")
	path := "D:\\Users\\MyWei\\test\\"
	var buffer = bytes.Buffer{}
	buffer.WriteString(path)
	buffer.WriteString("Chapter1_")
	buffer.WriteString(nowTime)
	buffer.WriteString(".ppm")
	result = buffer.String()
	CreateOrOpen(result)
	return
}
func CreateOrOpen(str string) {
	file, err := os.Open(str)
	if err != nil && os.IsNotExist(err) {
		file, _ = os.Create(str)
	}
	defer file.Close()
	height := displayImag.Height
	width := displayImag.Width
	fmt.Println(width)
	file.WriteString("P3\n" + strconv.Itoa(width) + " " + strconv.Itoa(height) + "\n255\n")
	for i := height - 1; i >= 0; i-- {
		for j := 0; j < width; j++ {
			r := float32(j) / float32(width)
			g := float32(i) / float32(height)
			b := float32(0.2)
			ir := int(255.59 * r)
			ig := int(255.59 * g)
			ib := int(255.59 * b)
			str := strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " + strconv.Itoa(ib) + "\n"
			//strBody := []byte(str)
			//buf := new(bytes.Buffer)
			//binary.Write(buf,binary.LittleEndian,strBody)
			file.WriteString(str)
		}
	}
}
func ReadDisplay(display *Display) {
	displayImag = *display
	fmt.Println(displayImag)
}
