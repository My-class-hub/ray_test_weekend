package vec3

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

//p(t) = A+t*B 即返回t时刻光线的位置
func (ray Ray) PointAtParameter(t float32) *Vec3 {
	return ray.Origin.Add(*ray.Direction.Scale(t))
}
