package rand

import (
	"fmt"
	"testing"
	"time"
)

func TestRandLitterNum(t *testing.T) {
	t1 := time.Now()
	for i := 0; i < 1000000; i++ {
		RandomSlice(1, 6, 3)
	}
	fmt.Println(time.Now().Sub(t1).Nanoseconds())

	t1 = time.Now()
	for i := 0; i < 100000; i++ {
		RandomSlice(1, 9, 3)
	}
	fmt.Println(time.Now().Sub(t1).Nanoseconds())
}
