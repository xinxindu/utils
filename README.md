# utils
## 工具类
#### 随机库(支持并发)
```go
// 随机len个小写字符
RandomLowercaseStr(len int) string 
RandomLowercaseStrC(len int) string

// 随机len个大写字符
RandomUppercaseStr(len int) string
RandomUppercaseStrC(len int) string

// 随机len个大小写字符
RandomLowerUppercaseStr(len int) string
RandomLowerUppercaseStrC(len int) string

// 随机len个大小写字符和下划线
RandomName(len int) string
RandomNameC(len int) string

// 指定范围内随机num个数字
RandomSlice(start, end, num int) []int
RandomSliceC(start, end, num int) []int

```
