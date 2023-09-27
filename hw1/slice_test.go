package hw1

import (
	"errors"
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {
	// 测试正常情况下的删除
	src1 := []int{1, 2, 3, 4, 5}
	expected1 := []int{1, 2, 4, 5}
	result1, err := Remove(src1, 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Expected %v, got %v", expected1, result1)
	}

	// 测试删除首元素
	src2 := []int{1, 2, 3, 4, 5}
	expected2 := []int{2, 3, 4, 5}
	result2, err := Remove(src2, 0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Expected %v, got %v", expected2, result2)
	}

	// 测试删除尾元素
	src3 := []int{1, 2, 3, 4, 5}
	expected3 := []int{1, 2, 3, 4}
	result3, err := Remove(src3, 4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Expected %v, got %v", expected3, result3)
	}

	// 测试超出索引范围的情况
	src4 := []int{1, 2, 3, 4, 5}
	_, err4 := Remove(src4, 5)
	expectedError4 := errors.New("index out of range")
	if !reflect.DeepEqual(err4, expectedError4) {
		t.Errorf("Expected error %v, got %v", expectedError4, err4)
	}
}

func TestDecreaseSlice(t *testing.T) {
	// 测试缩容到合适的大小（cap=4）
	src1 := []int{1, 2, 3, 4, 5}
	result1, _ := Remove(src1, 4)
	expected1Cap := 5
	if cap(result1) != expected1Cap {
		t.Errorf("Expected capacity %d, got %d", expected1Cap, cap(result1))
	}

	// 测试小内存的情况下（cap=1）
	src2 := []int{1}
	expected2Cap := 1
	result2, _ := Remove(src2, 0)
	if cap(result2) != expected2Cap {
		t.Errorf("Expected capacity %d, got %d", expected2Cap, cap(result2))
	}

	// 测试大内存的情况下（cap=500）
	src3 := make([]int, 1000)
	expected3Cap := 500

	for i := 0; i < 500; i++ {
		src3, _ = Remove(src3, 400)
	}
	if cap(src3) != expected3Cap {
		t.Errorf("Expected capacity %d, got %d", expected3Cap, cap(src3))
	}
}
