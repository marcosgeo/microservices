package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort_WorseCase(t *testing.T) {
	// Initialization
	elements := []int{10, 8, 6, 4, 2}

	// Execution
	BubbleSort(elements)

	// Validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, 2, elements[0])
	assert.EqualValues(t, 4, elements[1])
	assert.EqualValues(t, 6, elements[2])
	assert.EqualValues(t, 8, elements[3])
	assert.EqualValues(t, 10, elements[4])
}

func TestBubbleSort_BestCase(t *testing.T) {
	elem := []int{1, 2, 3, 4, 5}

	BubbleSort(elem)

	assert.NotNil(t, elem)
	assert.EqualValues(t, 5, len(elem))
	assert.EqualValues(t, 1, elem[0])
	assert.EqualValues(t, 2, elem[1])
	assert.EqualValues(t, 3, elem[2])
	assert.EqualValues(t, 4, elem[3])
	assert.EqualValues(t, 5, elem[4])
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	els := getElements(5)
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 4, els[0])
	assert.EqualValues(t, 3, els[1])
	assert.EqualValues(t, 2, els[2])
	assert.EqualValues(t, 1, els[3])
	assert.EqualValues(t, 0, els[4])
}

func BenchmarkBubbleSort_10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort_1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort_100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}
