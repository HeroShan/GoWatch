package seckill
import (
  "fmt"
  "strings"
  "time"
  "os"
)
func main() {
  for i := 0; i < 50; i++ {
    time.Sleep(100 * time.Millisecond)
    h := strings.Repeat("=", i) + strings.Repeat(" ", 49-i)
    fmt.Printf("\r%.0f%%[%s]", float64(i)/49*100, h)
    os.Stdout.Sync()
  }
  fmt.Println("\nAll System Go!")
}