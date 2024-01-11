package arrays

import "errors"

type Array struct {
	length int
	data   []interface{}
}

func (ar *Array) Push(value interface{}) {
	ar.data = append(ar.data, value)
	ar.length++
}

func (ar *Array) Pop() (interface{}, error) {
	if ar.length <= 0 {
		return nil, errors.New("Array is empty")
	}
	value := ar.data[ar.length-1]
	ar.data = ar.data[:ar.length-1]
	ar.length--

	return value, nil
}

func (ar *Array) Delete(index int) error {

	for i := index; i < ar.length-1; i++ {
		ar.data[i] = ar.data[i+1]
	}

	ar.data = ar.data[:ar.length-1]
	ar.length--

	return nil
}

func (ar *Array) Size() int {
	return ar.length
}
