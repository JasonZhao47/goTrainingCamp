package hw1

import (
	"errors"
)

func Remove[T comparable](src []T, idx int) ([]T, error) {
	// 为了防止大数据量时占用过多内存，采用原地搬移数组方式，而不是进行复制。
	srcLen := len(src)
	if idx < 0 || idx > srcLen-1 {
		return nil, errors.New("index out of range")
	}
	for i := idx; i < srcLen-1; i++ {
		src[i] = src[i+1]
	}

	src = src[:srcLen-1]

	return decreaseSlice(src), nil
}

func decreaseSlice[T comparable](src []T) []T {
	newLen := uint(len(src))
	oldCap := uint(cap(src))

	doubleLen := newLen + newLen
	newCap := oldCap
	// 缩容机制设置
	// 如果有大量内存空缺 > 2倍
	// 则直接/2
	if doubleLen <= oldCap && doubleLen > 0 {
		newCap = oldCap / 2
	}

	newArr := make([]T, newLen, newCap)

	copy(newArr, src)
	return newArr
}
