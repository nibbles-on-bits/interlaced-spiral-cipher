package main

import (
	"fmt"
)

func main() {
	/*fmt.Printf("LoopsNeeded(16) => %d\n", LoopsNeeded(16))
	fmt.Printf("LoopsNeeded(25) => %d\n", LoopsNeeded(25))
	fmt.Printf("LoopsNeeded(36) => %d\n", LoopsNeeded(36))
	fmt.Printf("LoopsNeeded(49) => %d\n", LoopsNeeded(49))
	fmt.Printf("LoopsNeeded(64) => %d\n", LoopsNeeded(64))
	fmt.Printf("LoopsNeeded(81) => %d\n", LoopsNeeded(81))
	fmt.Printf("LoopsNeeded(100) => %d\n", LoopsNeeded(100))
	fmt.Println()
	fmt.Printf("DimensionNeeded(16) => %d\n", DimensionNeeded(16))
	fmt.Printf("DimensionNeeded(25) => %d\n", DimensionNeeded(25))
	fmt.Printf("DimensionNeeded(36) => %d\n", DimensionNeeded(36))
	fmt.Printf("DimensionNeeded(49) => %d\n", DimensionNeeded(49))
	fmt.Printf("DimensionNeeded(64) => %d\n", DimensionNeeded(64))
	fmt.Printf("DimensionNeeded(81) => %d\n", DimensionNeeded(81))
	fmt.Printf("DimensionNeeded(100) => %d\n", DimensionNeeded(100))
	fmt.Println()
	fmt.Println('Encode("ABCDEFGHIJKLMNOP") => %v\n', Encode("ABCDEFGHIJKLMNOP"))*/

	fmt.Printf("GetEncodingPattern(3) => %v\n", GetEncodingPattern(3))
	fmt.Println()
	fmt.Printf("GetEncodingPattern(4) => %v\n", GetEncodingPattern(4))
	fmt.Println()
	fmt.Printf("GetEncodingPattern(5) => %v\n", GetEncodingPattern(5))
	fmt.Println()
	fmt.Printf("GetEncodingPattern(6) => %v\n", GetEncodingPattern(6))
	fmt.Println()
	fmt.Printf("GetEncodingPattern(7) => %v\n", GetEncodingPattern(7))
	fmt.Println()

}

// LoopsNeeded will determine how many loops are necessary to encode a nxn size string
// using Interlaced Spiral Cipher
//func RingsNeeded(n int) int {
//	return int((n + 1) / 2)
//}

/*func DimensionNeeded(n int) int {
	return RingsNeeded(n) + 2

}*/

func Encode(encodeMe string) [][]int {
	ret := [][]int{}

	return ret
}

// GetEncodingPattern will generate an encoding sequence 2D array of
// dimension x dimension size
// GetEncodingPattern will generate an encoding sequence 2D array of
// dimension x dimension size
func GetEncodingPattern(dimension int) [][]int {

	fmt.Printf("GetEncodingPattern   dimension=%d\n", dimension)

	ret := make([][]int, dimension)
	for i := 0; i < dimension; i++ {
		ret[i] = make([]int, dimension)
	}

	//ringsNeeded := RingsNeeded(dimension)
	ringsNeeded := (dimension + 1) / 2
	lastCnt := dimension*dimension - 1
	fmt.Printf("GetEncodingPattern   ringsNeeded=%d\n", ringsNeeded)
	fmt.Printf("GetEncodingPattern   lastCnt=%d\n", lastCnt)

	cnt := 0

	for ring := 0; ring < ringsNeeded; ring++ { // iterate thru rings 0 is outermost ring

		top := ring
		left := ring
		right := dimension - 1 - ring
		bottom := dimension - 1 - ring
		fmt.Printf("ring=%d top=%d left=%d bottom=%d right=%d\n", ring, top, left, bottom, right)

		for x := 0; x < (right - ring); x++ {

			ret[top][left+x] = cnt

			cnt++
			ret[top+x][right] = cnt

			cnt++
			ret[bottom][right-x] = cnt

			cnt++
			ret[bottom-x][left] = cnt

			if cnt == dimension*dimension-1 { // we are done ex : 4x4 would finish here
				return ret
			}

			if cnt == dimension*dimension-2 { // we are done ex : 5x5 would finish here
				cnt++
				ret[bottom-top][right-left] = cnt
				// place this in the center!
				return ret
			}
			cnt++
		}
	}
	return ret
}