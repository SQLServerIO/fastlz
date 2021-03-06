package fastlz

// #include <stdlib.h>
// #include "fastlz.h"
import "C"

// Other imports
import (
	"errors"
	"fmt"
	"unsafe"
)

// Compress applies fastlz_compress function to input
func Compress(input []byte) ([]byte, error) {
	length := len(input)
	if length == 0 {
		return nil, errors.New("fastlz: empty input")
	}

	// Output buffer
	output := make([]byte, int(float64(length)*1.4))

	// Run
	num := C.fastlz_compress(unsafe.Pointer(&input[0]), C.int(length), unsafe.Pointer(&output[0]))

	// Empty compression result
	if num == 0 {
		return nil, errors.New("fastlz: compression error, empty result")
	}

	return output[:num], nil
}

// Decompress applies fastlz_decompress function to input
func Decompress(input []byte, maxOut uint) ([]byte, error) {
	length := len(input)
	if length == 0 {
		return nil, errors.New("fastlz: empty input")
	}
	if maxOut == 0 {
		return nil, errors.New("fastlz: invalid max out value")
	}

	// Output buffer
	output := make([]byte, int(maxOut))

	// Run
	num := C.fastlz_decompress(unsafe.Pointer(&input[0]), C.int(length), unsafe.Pointer(&output[0]), C.int(maxOut))
	goNum := int(num)

	// Empty decompression result
	if goNum == 0 {
		return nil, errors.New("fastlz: decompression error, empty result")
	}

	if goNum != int(maxOut) {
		return nil, fmt.Errorf("fastlz: decompression error, shit happens! Max out: %d, num: %d, input data:%#v", maxOut, goNum, input)
	}

	return output[:goNum], nil
}
