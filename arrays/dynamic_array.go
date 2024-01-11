package arrays

type DynamicArray struct {
	Capacity int
	Length   int
	Data     []int
}

func NewDynamicArray(capacity int) *DynamicArray {
	if capacity > 0 {
		return &DynamicArray{
			Capacity: capacity,
			Length:   0,
			Data:     make([]int, capacity),
		}
	}

	return &DynamicArray{
		Capacity: 0,
		Length:   0,
		Data:     make([]int, 0),
	}
}

func (d *DynamicArray) Get(i int) int {
	if i > d.Length {
		return -1
	}
	return d.Data[i]
}

func (d *DynamicArray) Set(i, n int) {
	if i <= d.Length {
		d.Data[i] = n
		d.Length++
	}
}

func (d *DynamicArray) GetSize() int {
	return d.Length
}

func (d *DynamicArray) GetCapacity() int {
	return d.Capacity
}

func (d *DynamicArray) Resize() {
	d.Capacity = d.Capacity * 2
	newArray := make([]int, d.Capacity)
	copy(newArray, d.Data)
	d.Data = newArray
}

func (d *DynamicArray) Pushback(n int) {
	if d.Length == d.Capacity {
		d.Resize()
	}
	d.Data[d.Length] = n
	d.Length++
}

func (d *DynamicArray) Popback() int {
	d.Data = d.Data[:d.Length-1]
	d.Length--
	return d.Data[d.Length-1]
}
