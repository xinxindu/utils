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
var step = 6
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

func RandomLowercaseStr(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := mRand.Intn(26) + 97
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func RandomUppercaseStr(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := mRand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func randomLowercaseBytes(len int) []byte {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := mRand.Intn(26) + 97
		bytes[i] = byte(b)
	}
	return bytes
}

func randomUppercaseBytes(len int) []byte {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := mRand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return bytes
}

func RandomLowerUppercaseStr(len int) string {
	bytes := make([]byte, 0, len)
	var num int
	for i := 0; i < len; {
		if len-i < step {
			num = mRand.Intn(len-i) + 1
		} else {
			num = mRand.Intn(step) + 1
		}

		if num%2 == 0 {
			bytes = append(bytes, randomLowercaseBytes(num)...)
		} else {
			bytes = append(bytes, randomUppercaseBytes(num)...)
		}
		i += num
	}
	return string(bytes)
}

func RandomName(len int) string {
	bytes := make([]byte, 0, len)
	var num int
	for i := 0; i < len; {
		if len-i < step {
			num = mRand.Intn(len-i) + 1
		} else {
			num = mRand.Intn(step) + 1
		}

		switch num % 3 {
		case 0:
			bytes = append(bytes, randomLowercaseBytes(num)...)
			i += num
		case 1:
			bytes = append(bytes, randomUppercaseBytes(num)...)
			i += num
		case 2:
			bytes = append(bytes, 95)
			i++
		}
	}
	return string(bytes)
}
