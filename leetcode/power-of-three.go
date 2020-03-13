package leetcode



func IsPowerOfThree(n int) (result bool) {
	if n == 1{
		return true
	}
	cn := float64(n)
	for{
		if cn == 3{
			result = true
			break
		}else{
			if cn<3{
				result = false
				break
			}
		}
		cn = cn/3
	}
	return result

}