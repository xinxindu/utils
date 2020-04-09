package rand

import (
	"math/rand"
)

/*
	支持并发的随机
*/

// RandomSlice: chose num items from [start, end)
// start: at least is 1 .
func RandomSliceC(r *rand.Rand, start, end, num int) []int {
	if end-start <= num {
		return nil
	}
	ret := make([]int, 0, num)
	cnt := end - start
	for i := start; i < end; i++ {
		r := r.Int() % cnt
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

func RandomLowercaseStrC(r *rand.Rand, len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 97
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func RandomUppercaseStrC(r *rand.Rand, len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func randomLowercaseBytesC(r *rand.Rand, len int) []byte {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 97
		bytes[i] = byte(b)
	}
	return bytes
}

func randomUppercaseBytesC(r *rand.Rand, len int) []byte {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return bytes
}

func RandomLowerUppercaseStrC(r *rand.Rand, len int) string {
	bytes := make([]byte, 0, len)
	var num int
	for i := 0; i < len; {
		if len-i < step {
			num = r.Intn(len-i) + 1
		} else {
			num = r.Intn(step) + 1
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

func RandomNameC(r *rand.Rand, len int) string {
	bytes := make([]byte, 0, len)
	var num int
	for i := 0; i < len; {
		if len-i < step {
			num = r.Intn(len-i) + 1
		} else {
			num = r.Intn(step) + 1
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
