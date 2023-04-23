package collections_test

import (
	"errors"
	"linq/collections"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmptyList(t *testing.T) {
	intListPointer := collections.NewEmptyList[int](1)
	stringListPointer := collections.NewEmptyList[string](1)

	assert.NotNil(t, intListPointer)
	assert.NotNil(t, stringListPointer)
}

func TestNewListFromArray(t *testing.T) {
	inputs := []struct {
		array []int
	}{
		{
			array: []int{1, 2, 3},
		},
		{
			array: []int{1, 2, 3, 4, 5},
		},
		{
			array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, input := range inputs {
		l := collections.NewListFromArray(input.array)
		assert.Equal(t, input.array, l.ToArray())
	}
}

func TestListAdd(t *testing.T) {
	l := collections.NewEmptyList[int](10)
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, input := range inputs {
		l.Add(input)
		assert.Equal(t, inputs[0:index+1], l.ToArray())
	}
}

func TestListContains(t *testing.T) {
	l := collections.NewListFromArray([]int{1, 2, 2, 4, 6, 7, 9, 10})
	testCases := []struct {
		input  int
		output bool
	}{
		{input: 1, output: true},
		{input: 2, output: true},
		{input: 3, output: false},
		{input: 4, output: true},
		{input: 5, output: false},
		{input: 6, output: true},
		{input: 7, output: true},
		{input: 8, output: false},
		{input: 9, output: true},
		{input: 10, output: true},
	}

	for _, tc := range testCases {
		actualOutput := l.Contains(tc.input)
		assert.Equal(t, tc.output, actualOutput)
	}
}

func TestListCountOf(t *testing.T) {
	l := collections.NewListFromArray([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3})
	testCases := []struct {
		input  int
		output int
	}{
		{input: 1, output: 4},
		{input: 2, output: 3},
		{input: 3, output: 5},
		{input: 4, output: 1},
		{input: 5, output: 2},
	}

	for _, tc := range testCases {
		actualOutput := l.CountOf(tc.input)
		assert.Equal(t, tc.output, actualOutput)
	}
}

func TestListDistinct(t *testing.T) {
	testCases := []struct {
		input  *collections.List[int]
		output *collections.List[int]
	}{
		{
			input:  collections.NewListFromArray([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			output: collections.NewListFromArray([]int{1, 2, 3, 5, 4}),
		},
		{
			input:  collections.NewListFromArray([]int{1, 1, 2, 2, 3, 3, 4, 5, 4}),
			output: collections.NewListFromArray([]int{1, 2, 3, 4, 5}),
		},
		{
			input:  collections.NewListFromArray([]int{1, 2, 3, 4, 5}),
			output: collections.NewListFromArray([]int{1, 2, 3, 4, 5}),
		},
		{
			input:  collections.NewListFromArray([]int{5, 4, 3, 2, 1}),
			output: collections.NewListFromArray([]int{5, 4, 3, 2, 1}),
		},
	}

	for _, tc := range testCases {
		actualOutput := tc.input.Distinct()
		assert.Equal(t, tc.output, actualOutput)
	}
}

func TestListExtend(t *testing.T) {
	testCases := []struct {
		list1  *collections.List[int]
		list2  *collections.List[int]
		output *collections.List[int]
	}{
		{
			list1:  collections.NewListFromArray([]int{1, 2, 2, 1, 3, 5, 5}),
			list2:  collections.NewListFromArray([]int{3, 3, 1, 4, 2, 1, 3, 3}),
			output: collections.NewListFromArray([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
		},
		{
			list1:  collections.NewListFromArray([]int{1, 2, 3, 4, 5}),
			list2:  collections.NewListFromArray([]int{6, 7, 8, 9, 10}),
			output: collections.NewListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		},
		{
			list1:  collections.NewListFromArray([]int{1, 1, 1, 1, 1}),
			list2:  collections.NewListFromArray([]int{0, 0}),
			output: collections.NewListFromArray([]int{1, 1, 1, 1, 1, 0, 0}),
		},
	}

	for _, tc := range testCases {
		tc.list1.Extend(tc.list2)
		assert.Equal(t, tc.output, tc.list1)
	}
}

func TestListGet(t *testing.T) {
	l := collections.NewListFromArray([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	testCases := []struct {
		input  int
		output int
		err    error
	}{
		{input: -1, output: 0, err: errors.New("INDEX_OUT_OF_RANGE")},
		{input: 0, output: 0, err: nil},
		{input: 1, output: 1, err: nil},
		{input: 2, output: 2, err: nil},
		{input: 3, output: 3, err: nil},
		{input: 4, output: 4, err: nil},
		{input: 5, output: 5, err: nil},
		{input: 6, output: 6, err: nil},
		{input: 7, output: 7, err: nil},
		{input: 8, output: 8, err: nil},
		{input: 9, output: 9, err: nil},
		{input: 10, output: 10, err: nil},
		{input: 11, output: 0, err: errors.New("INDEX_OUT_OF_RANGE")},
	}

	for _, tc := range testCases {
		output, err := l.Get(tc.input)
		assert.Equal(t, tc.output, output)
		assert.Equal(t, tc.err, err)
	}
}

func TestIndexOf(t *testing.T) {
	l := collections.NewListFromArray([]int{1, 2, 3, 4, 5})
	testCases := []struct {
		input  int
		output int
	}{
		{input: 0, output: -1},
		{input: 1, output: 0},
		{input: 2, output: 1},
		{input: 3, output: 2},
		{input: 4, output: 3},
		{input: 5, output: 4},
		{input: 6, output: -1},
	}

	for _, tc := range testCases {
		output := l.IndexOf(tc.input)
		assert.Equal(t, tc.output, output)
	}
}

func TestListRemoveAll(t *testing.T) {
	l := collections.NewListFromArray([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3})
	testCases := []struct {
		input int
		err   error
	}{
		{input: 1, err: nil},
		{input: 2, err: nil},
		{input: 3, err: nil},
		{input: 4, err: nil},
		{input: 5, err: nil},
		{input: 1, err: errors.New("ITEM_NOT_FOUND")},
	}

	for _, tc := range testCases {
		err := l.RemoveAll(tc.input)
		removed := !l.Contains(tc.input)
		assert.Equal(t, true, removed)
		assert.Equal(t, tc.err, err)
	}
}

func TestListRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		input  *collections.List[int]
		output *collections.List[int]
	}{
		{
			input:  collections.NewListFromArray([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			output: collections.NewListFromArray([]int{1, 2, 3, 5, 4}),
		},
		{
			input:  collections.NewListFromArray([]int{1, 1, 2, 2, 3, 3, 4, 5, 4}),
			output: collections.NewListFromArray([]int{1, 2, 3, 4, 5}),
		},
		{
			input:  collections.NewListFromArray([]int{1, 2, 3, 4, 5}),
			output: collections.NewListFromArray([]int{1, 2, 3, 4, 5}),
		},
		{
			input:  collections.NewListFromArray([]int{5, 4, 3, 2, 1}),
			output: collections.NewListFromArray([]int{5, 4, 3, 2, 1}),
		},
	}

	for _, tc := range testCases {
		tc.input.RemoveDuplicates()
		assert.Equal(t, tc.input.ToArray(), tc.output.ToArray())
	}
}

func TestListRemoveFirst(t *testing.T) {
	testCases := []struct {
		inputList    *collections.List[int]
		inputElement int
		output       *collections.List[int]
		err          error
	}{
		{
			inputList:    collections.NewListFromArray([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			inputElement: 1,
			output:       collections.NewListFromArray([]int{2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			err:          nil,
		},
		{
			inputList:    collections.NewListFromArray([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			inputElement: 4,
			output:       collections.NewListFromArray([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 2, 1, 3, 3}),
			err:          nil,
		},
		{
			inputList:    collections.NewListFromArray([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			inputElement: 10,
			output:       collections.NewListFromArray([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			err:          errors.New("ITEM_NOT_FOUND"),
		},
	}

	for _, tc := range testCases {
		err := tc.inputList.RemoveFirst(tc.inputElement)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.inputList.ToArray(), tc.output.ToArray())
	}
}

func TestListWhere(t *testing.T) {
	testCases := []struct {
		inputList  *collections.List[int]
		outputList *collections.List[int]
		filterFunc func(int) bool
	}{
		{
			inputList:  collections.NewListFromArray([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			outputList: collections.NewListFromArray([]int{2, 2, 4, 2}),
			filterFunc: func(i int) bool { return i%2 == 0 },
		},
		{
			inputList:  collections.NewListFromArray([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			outputList: collections.NewListFromArray([]int{1, 2, 2, 1, 1, 2, 1}),
			filterFunc: func(i int) bool { return i < 3 },
		},
	}

	for _, tc := range testCases {
		outputList := tc.inputList.Where(tc.filterFunc)
		assert.Equal(t, tc.outputList.ToArray(), outputList.ToArray())
	}
}

func TestListSize(t *testing.T) {
	testCases := []struct {
		inputList *collections.List[int]
		output    int
	}{
		{
			inputList: collections.NewListFromArray([]int{}),
			output:    0,
		},
		{
			inputList: collections.NewListFromArray([]int{1}),
			output:    1,
		},
		{
			inputList: collections.NewListFromArray([]int{1, 2, 3, 4, 5}),
			output:    5,
		},
	}

	for _, tc := range testCases {
		output := tc.inputList.Size()
		assert.Equal(t, tc.output, output)
	}
}

func TestListString(t *testing.T) {
	testCases := []struct {
		inputList *collections.List[int]
		output    string
	}{
		{
			inputList: collections.NewListFromArray([]int{}),
			output:    "[]",
		},
		{
			inputList: collections.NewListFromArray([]int{1}),
			output:    "[1]",
		},
		{
			inputList: collections.NewListFromArray([]int{1, 2, 3, 4, 5}),
			output:    "[1,2,3,4,5]",
		},
	}

	for _, tc := range testCases {
		output := tc.inputList.String()
		assert.Equal(t, tc.output, output)
	}
}

func TestListType(t *testing.T) {
	l := collections.NewEmptyList[int](0)

	assert.Equal(t, collections.TypeList, l.Type())
}
