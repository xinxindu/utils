package rand

import (
	"math/rand"
	"time"
)

/*
	RandNum 在[start,end)中随机num个不重复的元素
    适合小范围内随机少数的元素
*/

var mRand *rand.Rand
var mSlice = make([]int, 1000)

func init() {
	mRand = rand.New(rand.NewSource(time.Now().Unix()))
}

// RandomSlice: chose num items from [start, end)
// start: at least is 1 .
func RandomSlice(start, end, num int) []int {
	if end-start <= num {
		return nil
	}
	ret := make([]int, 0, num)
	cnt := end - start
	for i := start; i < end; i++ {
		r := mRand.Int() % cnt
		if mSlice[r+i] == 0 {
			ret = append(ret, r+i)
		} else {
			ret = append(ret, mSlice[r+i])
		}
		if len(ret) == num {
			break
		}
		mSlice[r+i] = i
		cnt--
	}

	for _, v := range ret {
		mSlice[v] = 0
	}
	return ret
}
