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

func twoSum(nums []int, target int) []int {
    var result []int
    var kMax,cursor int
    kMax = len(nums)
    for k,v := range nums{
        if k == kMax{
            break
        }
        for cursor = k+1;cursor < kMax;cursor++{
            if v+nums[cursor] == target{
            result = append(result,k)
            result = append(result,cursor)
            }
        }
        
    }
    return result
}
