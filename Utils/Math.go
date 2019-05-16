package Utils

func Abs(x float64) float64{
	if x < 0{
		return -x
	}

	return x
}

func EqualsFloat(f1 float64, f2 float64) bool{
	if Abs(f1 - f2) < 0.0001{
		return true
	}else{
		return false
	}
}