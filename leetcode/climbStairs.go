package leetcode
//泰波那契数
func ClimbStairs(n int) int {
    if n < 2{
        return 1
    }
    
    a := [][]int{
        []int{1, 1}, // T(0), T(1)
        []int{1, 2},  // T(1), T(2)
    }

    nx := [][]int{
        []int{0, 1},
        []int{1, 1},
    }

    res := matriceMultiple(a, matircePow(nx, n-2))
    return res[1][1]
}

// 矩阵快速幂
func matircePow(a [][]int, n int)[][]int{
    tag := a
    res := [][]int{
        []int{1, 0},
        []int{0, 1},
    }

    for n!=0{
        if n&1 != 0{
            res = matriceMultiple(res, tag)
        }
        tag = matriceMultiple(tag, tag)
        n >>= 1
    }
    return res
}

// 矩阵相乘
func matriceMultiple(a [][]int, b [][]int)[][]int{
    return [][]int{
        []int{ a[0][0] * b[0][0] + a[0][1]*b[1][0], a[0][0] * b[0][1] + a[0][1]*b[1][1]},
        []int{ a[1][0] * b[0][0] + a[1][1]*b[1][0], a[1][0] * b[0][1] + a[1][1]*b[1][1]},
    }
}