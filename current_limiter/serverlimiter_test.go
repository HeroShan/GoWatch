package current_limiter

import(
	"testing"
	"strconv"
)

func BenchmarkSerlock(b *testing.B){
	for i := 0; i < b.N; i++ {
        for j := 0; j<100;j++{
			Serlock("qwerr"+strconv.Itoa(j))
		}
    }

}