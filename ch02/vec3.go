package vec3

import "math"

type Vec3 struct {
	e [3]float32
}

//模拟构造函数
func NewVec3(x, y, z float32) *Vec3 {
	f := [3]float32{x, y, z}
	vec3 := Vec3{e: f}
	return &vec3
}

//获取三维坐标任一维度索引,返回对应坐标的值
func (v1 Vec3) X() float32 {
	return v1.e[0]
}
func (v1 Vec3) Y() float32 {
	return v1.e[1]
}
func (v1 Vec3) Z() float32 {
	return v1.e[2]
}

//向量求和函数的形式举个例子
//func Add(e1,e2 Vec3) Vec3{
//	return Vec3{
//		e: [3]float32{e1.X()+e2.X(),e1.Y()+e2.Y(),e1.Z()+e2.Z()},
//	}
//}
//方法的形式

//因为我在这里并没有添加指针对象所以使用还是要产生一个新对象
//Add adds the 2 vectors (return a new vector)
func (v1 Vec3) Add(v2 Vec3) *Vec3 {
	return &Vec3{e: [3]float32{v1.X() + v2.X(), v1.Y() + v2.Y(), v1.Z() + v2.Z()}}
}

//Sub substracts the 2 vector (return a new vector)
func (v1 Vec3) Sub(v2 Vec3) *Vec3 {
	return &Vec3{e: [3]float32{v1.X() - v2.X(), v1.Y() - v2.Y(), v1.Z() - v2.Z()}}
}

//Mult multiplies the 2 vectors （return a new vector）
func (v1 Vec3) Mult(v2 Vec3) *Vec3 {
	return &Vec3{e: [3]float32{v1.X() * v2.X(), v1.Y() * v2.Y(), v1.Z() * v2.Z()}}
}

// Length returns the size of the vector,向量的模长
func (v1 Vec3) Length() float32 {
	return float32(math.Sqrt(float64(v1.X()*v1.X() + v1.Y()*v1.Y() + v1.Z()*v1.Z())))
}

//SqrLength return the sqr size of the vector
func (v1 Vec3) SqrLength() float32 {
	return float32(math.Pow(float64(v1.X()), 2) + math.Pow(float64(v1.Y()), 2) + math.Pow(float64(v1.Z()), 2))
}

//Scale scales the vector by the value (return a new vector)
//向量数乘
func (v1 Vec3) Scale(t float32) *Vec3 {
	return &Vec3{e: [3]float32{v1.X() * t, v1.Y() * t, v1.Z() * t}}
}

//Unit returns a new vector with same direction and length 1
//向量归一化
func (v1 Vec3) Unit() *Vec3 {
	return v1.Scale(1 / v1.Length())
}

//Negate returns a new vector with X/Y/Z negated
//向量取反
func (v1 Vec3) Negate() *Vec3 {
	return &Vec3{[3]float32{-v1.X(), -v1.Y(), -v1.Z()}}
}

// Dot returns the dot product (a scalar) of 2 vectors
//点乘
func Dot(v1, v2 Vec3) float32 {
	return v1.X()*v2.X() + v1.Y() + v2.Y() + v1.Z() + v2.Z()
}

//Cross returns the cross product of 2 vectors (another vector)
//叉乘
func Cross(v1, v2 Vec3) *Vec3 {
	return &Vec3{e: [3]float32{
		v1.Y()*v2.Z() - v1.Z()*v2.Y(),
		v1.Z()*v2.X() - v1.X()*v2.Z(),
		v1.X()*v2.Y() - v1.Y()*v2.X(),
	}}
}

//Reflect simply reflects the vector based on the normal n
//反射
func (v1 Vec3) Reflect(v2 Vec3) *Vec3 {
	return v1.Sub(*v2.Scale(2.0 * Dot(v1, v2)))
}
