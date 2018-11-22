package commonslice

// RemoveDuplicateByte receives a slice of bytes and returns without the duplicate ones.
func RemoveDuplicateByte(elements []byte) []byte {
	var encountered = map[byte]bool{}
	var index int
	var key byte
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []byte{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateComplex128 receives a slice of complex128s and returns without the duplicate ones.
func RemoveDuplicateComplex128(elements []complex128) []complex128 {
	var encountered = map[complex128]bool{}
	var index int
	var key complex128
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []complex128{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateComplex64 receives a slice of complex64s and returns without the duplicate ones.
func RemoveDuplicateComplex64(elements []complex64) []complex64 {
	var encountered = map[complex64]bool{}
	var index int
	var key complex64
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []complex64{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateFloat32 receives a slice of float32s and returns without the duplicate ones.
func RemoveDuplicateFloat32(elements []float32) []float32 {
	var encountered = map[float32]bool{}
	var index int
	var key float32
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []float32{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateFloat64 receives a slice of float64s and returns without the duplicate ones.
func RemoveDuplicateFloat64(elements []float64) []float64 {
	var encountered = map[float64]bool{}
	var index int
	var key float64
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []float64{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateInt receives a slice of ints and returns without the duplicate ones.
func RemoveDuplicateInt(elements []int) []int {
	var encountered = map[int]bool{}
	var index int
	var key int
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []int{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateInt16 receives a slice of int16s and returns without the duplicate ones.
func RemoveDuplicateInt16(elements []int16) []int16 {
	var encountered = map[int16]bool{}
	var index int
	var key int16
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []int16{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateInt32 receives a slice of int32s and returns without the duplicate ones.
func RemoveDuplicateInt32(elements []int32) []int32 {
	var encountered = map[int32]bool{}
	var index int
	var key int32
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []int32{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateInt64 receives a slice of int64s and returns without the duplicate ones.
func RemoveDuplicateInt64(elements []int64) []int64 {
	var encountered = map[int64]bool{}
	var index int
	var key int64
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []int64{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateInt8 receives a slice of int8s and returns without the duplicate ones.
func RemoveDuplicateInt8(elements []int8) []int8 {
	var encountered = map[int8]bool{}
	var index int
	var key int8
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []int8{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateRune receives a slice of runes and returns without the duplicate ones.
func RemoveDuplicateRune(elements []rune) []rune {
	var encountered = map[rune]bool{}
	var index int
	var key rune
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []rune{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateUint receives a slice of uints and returns without the duplicate ones.
func RemoveDuplicateUint(elements []uint) []uint {
	var encountered = map[uint]bool{}
	var index int
	var key uint
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []uint{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateUint16 receives a slice of uint16s and returns without the duplicate ones.
func RemoveDuplicateUint16(elements []uint16) []uint16 {
	var encountered = map[uint16]bool{}
	var index int
	var key uint16
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []uint16{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateUint32 receives a slice of uint32s and returns without the duplicate ones.
func RemoveDuplicateUint32(elements []uint32) []uint32 {
	var encountered = map[uint32]bool{}
	var index int
	var key uint32
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []uint32{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateUint64 receives a slice of uint64s and returns without the duplicate ones.
func RemoveDuplicateUint64(elements []uint64) []uint64 {
	var encountered = map[uint64]bool{}
	var index int
	var key uint64
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []uint64{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateUint8 receives a slice of uint8s and returns without the duplicate ones.
func RemoveDuplicateUint8(elements []uint8) []uint8 {
	var encountered = map[uint8]bool{}
	var index int
	var key uint8
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []uint8{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}

// RemoveDuplicateUintptr receives a slice of uintptrs and returns without the duplicate ones.
func RemoveDuplicateUintptr(elements []uintptr) []uintptr {
	var encountered = map[uintptr]bool{}
	var index int
	var key uintptr
	for index = range elements {
		encountered[elements[index]] = true
	}
	var result = []uintptr{}
	for key = range encountered {
		result = append(result, key)
	}
	return result
}
