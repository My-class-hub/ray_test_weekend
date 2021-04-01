package main

import vec3 "ray_test_weekend/ch02"

func main() {
	//d := display.Display{200, 100, "Title"}
	//display.ReadDisplay(&d)
	//pathURL := display.PathURL()
	//fmt.Println(pathURL)

	//test Vec3
	//newVec3 := vec3.NewVec3(1, 2, 3)
	//fmt.Println(newVec3.X())
	//fmt.Println(newVec3.Y())
	//fmt.Println(newVec3.Z())
	//fmt.Println("--------------")
	//newVec3 = newVec3.Scale(2)
	//fmt.Println(*newVec3)
	//fmt.Println(newVec3.X())
	//fmt.Println(newVec3.Y())
	//fmt.Println(newVec3.Z())
	display := vec3.Display{
		Width:      1980,
		Height:     1080,
		Title:      "sb",
		LowerLeft:  *vec3.NewVec3(-2.0, -1.0, -1.0),
		Horizontal: *vec3.NewVec3(4.0, 0.0, 0.0),
		Vertical:   *vec3.NewVec3(0.0, 2.0, 0.0),
		Origin:     *vec3.NewVec3(0.0, 0.0, 0.0),
	}
	display.CreateOrOpen()

}
