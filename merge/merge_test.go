package merge

import (
	"reflect"
	"testing"
)

func TestMerge_Sort_Int(t *testing.T) {
	unsortedArr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	expectedArr := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}

	result := Merge[int]{}.Sort(unsortedArr)

	if !reflect.DeepEqual(expectedArr, result) {
		t.Errorf("expected %v, got %v", expectedArr, result)
	}
}

func TestMerge_Sort_Int8(t *testing.T) {
	unsortedArr := []int8{12, 34, -56, 78, 90, -12, 34, -56, 78, 90}
	expectedArr := []int8{-56, -56, -12, 12, 34, 34, 78, 78, 90, 90}

	result := Merge[int8]{}.Sort(unsortedArr)

	if !reflect.DeepEqual(expectedArr, result) {
		t.Errorf("expected %v, got %v", expectedArr, result)
	}
}

func TestMerge_Sort_Uint32(t *testing.T) {
	unsortedArr := []uint32{4294967295, 0, 123456789, 47, 643, 987654321, 456789123}
	expectedArr := []uint32{0, 47, 643, 123456789, 456789123, 987654321, 4294967295}

	result := Merge[uint32]{}.Sort(unsortedArr)

	if !reflect.DeepEqual(expectedArr, result) {
		t.Errorf("expected %v, got %v", expectedArr, result)
	}
}

func TestMerge_Sort_Float64(t *testing.T) {
	unsortedArr := []float64{3.14, 2.71, 1.61, 0.0, 1.41, 2.0, 3.14, 2.71, 1.61}
	expectedArr := []float64{0.0, 1.41, 1.61, 1.61, 2.0, 2.71, 2.71, 3.14, 3.14}

	result := Merge[float64]{}.Sort(unsortedArr)

	if !reflect.DeepEqual(expectedArr, result) {
		t.Errorf("expected %v, got %v", expectedArr, result)
	}
}

func Test_merge(t *testing.T) {
	leftArr := []int{1, 2, 3}
	rightArr := []int{4, 5, 6}
	expectedArr := []int{1, 2, 3, 4, 5, 6}

	result := merge(leftArr, rightArr)

	if !reflect.DeepEqual(expectedArr, result) {
		t.Errorf("expected %v, got %v", expectedArr, result)
	}
}
