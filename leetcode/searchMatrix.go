package leetcode
func SearchMatrix(matrix [][]int, target int) bool {
    var(
		matrixMax int
		matrixMin int  = 0
	)
	if matrix == nil || len(matrix) == 0 {
        return false
	}
	
	for _,array := range matrix{
		matrixMax = len(array)-1
		if array[matrixMin] > target || array[matrixMax] < target {
			continue
		}else{
			for _,value := range array{
				if value > target{
					continue
				}
				if value == target{
					return true
				}

			}
		}
		
	}
	return false

}
