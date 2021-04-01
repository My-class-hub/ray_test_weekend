package vec3

import (
	"bytes"
	"os"
	"strconv"
	"time"
)

type Display struct {
	Width      int
	Height     int
	Title      string
	LowerLeft  Vec3
	Horizontal Vec3
	Vertical   Vec3
	Origin     Vec3
}

func initPath() string {
	nowTime := time.Now().Format("15_04_05")
	path := "D:\\Users\\MyWei\\test\\"
	//拼接字符串
	buffer := bytes.Buffer{}
	buffer.WriteString(path)
	buffer.WriteString("Chapter2_")
	buffer.WriteString(nowTime)
	buffer.WriteString(".ppm")
	return buffer.String()
}
func (display Display) CreateOrOpen() {
	path := initPath()
	file, err := os.Open(path)
	if err != nil && os.IsNotExist(err) {
		file, _ = os.Create(path)
	}
	defer file.Close()

	width := display.Width
	height := display.Height

	//写文件
	file.WriteString("P3\n" + strconv.Itoa(width) + " " + strconv.Itoa(height) + "\n255\n")
	for i := height - 1; i >= 0; i-- {
		for j := 0; j < width; j++ {
			u := float32(j) / float32(width)
			v := float32(i) / float32(height)
			var ray Ray = Ray{Origin: display.Origin,
				Direction: *display.LowerLeft.Add(*display.Horizontal.Scale(u).Add(*display.Vertical.Scale(v)))} //每一条光线
			color := display.Color(ray)
			ir := int(255.59 * color.X())
			ig := int(255.59 * color.Y())
			ib := int(255.59 * color.Z())
			file.WriteString(strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " + strconv.Itoa(ib) + "\n")
			//file.Write([]byte{byte(ir),byte(ig),byte(ib),byte("\n")})
		}
	}
}
func (display Display) Color(r Ray) Vec3 {
	unit := r.Direction.Unit()
	var t float32 = 0.5 * (unit.Y() + 1.0)
	return *Vec3{[3]float32{1.0, 1.0, 1.0}}.Scale(1.0 - t).Add(*Vec3{[3]float32{0.5, 0.7, 1.0}}.Scale(t))
}
