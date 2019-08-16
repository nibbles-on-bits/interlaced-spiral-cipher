package main

import (
	"fmt"
	"math"
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

	//fmt.Printf("GetEncodingPattern(3) => %v\n", GetEncodingPattern(3))
	//fmt.Println()
	/*fmt.Printf("GetEncodingPattern(4) => %v\n", GetEncodingPattern(4))
	fmt.Println()
	fmt.Printf("GetEncodingPattern(5) => %v\n", GetEncodingPattern(5))
	fmt.Println()
	fmt.Printf("GetEncodingPattern(6) => %v\n", GetEncodingPattern(6))
	fmt.Println()
	fmt.Printf("GetEncodingPattern(7) => %v\n", GetEncodingPattern(7))
	fmt.Println()*/

	/*for i := 2; i < 10; i++ {
		PrintEncodingPattern(GetEncodingPattern(i))
	}*/

	fmt.Printf(`Encode("Romani ite domum") = %s`+"\n", Encode("Romani ite domum"))

	// Failing on this
	fmt.Printf(`Encode("Sic transit gloria mundi ") = %s`+"\n", Encode("Sic transit gloria mundi "))

}

func Encode(encodeMe string) string {

	//TODO : test that string length is a square

	dimension := int(math.Sqrt(float64((len(encodeMe)))))

	ret := [][]byte{}

	ret = make([][]byte, dimension)
	for i := 0; i < dimension; i++ {
		ret[i] = make([]byte, dimension)
	}

	ringsNeeded := (dimension + 1) / 2
	cnt := 0

	for ring := 0; ring < ringsNeeded; ring++ { // iterate thru rings 0 is outermost ring
		top := ring
		left := ring
		right := dimension - 1 - ring
		bottom := dimension - 1 - ring

		for x := 0; x < (right - ring); x++ {

			ret[top][left+x] = encodeMe[cnt]

			cnt++
			ret[top+x][right] = encodeMe[cnt]

			cnt++
			ret[bottom][right-x] = encodeMe[cnt]

			cnt++
			ret[bottom-x][left] = encodeMe[cnt]

			if cnt == dimension*dimension-1 { // we are done ex : 4x4 would finish here
				//return ret
			}

			if cnt == dimension*dimension-2 { // we are done ex : 5x5 would finish here
				cnt++

				ret[(dimension+1)/2-1][(dimension+1)/2-1] = encodeMe[cnt]
				// place this in the center!
				//return ret
			}
			cnt++
		}
	}

	encodedString := ""
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			encodedString = encodedString + string(ret[i][j])
		}
	}

	return encodedString
}

func Decode(decodeMe string) string {

	//TODO : test that string length is a square

	dimension := int(math.Sqrt(float64((len(decodeMe)))))

	for x := 0; x < dimension; x++ {
		for y := 0; y < dimension; y++ {

		}
	}

	ret := [][]byte{}

	ret = make([][]byte, dimension)
	for i := 0; i < dimension; i++ {
		ret[i] = make([]byte, dimension)
	}

	ringsNeeded := (dimension + 1) / 2
	cnt := 0

	for ring := 0; ring < ringsNeeded; ring++ { // iterate thru rings 0 is outermost ring
		top := ring
		left := ring
		right := dimension - 1 - ring
		bottom := dimension - 1 - ring

		for x := 0; x < (right - ring); x++ {

			ret[top][left+x] = decodeMe[cnt]

			cnt++
			ret[top+x][right] = decodeMe[cnt]

			cnt++
			ret[bottom][right-x] = decodeMe[cnt]

			cnt++
			ret[bottom-x][left] = decodeMe[cnt]

			if cnt == dimension*dimension-1 { // we are done ex : 4x4 would finish here
				//return ret
			}

			if cnt == dimension*dimension-2 { // we are done ex : 5x5 would finish here
				cnt++

				ret[(dimension+1)/2-1][(dimension+1)/2-1] = decodeMe[cnt]
				// place this in the center!
				//return ret
			}
			cnt++
		}
	}

	encodedString := ""
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			encodedString = encodedString + string(ret[i][j])
		}
	}

	return encodedString
}

func PrintEncodingPattern(pattern [][]int) {
	fmt.Println("Pattern Length : %d", len(pattern))
	for x := 0; x < len(pattern); x++ {
		for y := 0; y < len(pattern); y++ {
			fmt.Printf("%3d ", pattern[x][y])
		}
		fmt.Println()
	}
}

func GetEncodingPattern(dimension int) [][]int {

	fmt.Printf("GetEncodingPattern   dimension=%d\n", dimension)

	ret := make([][]int, dimension)
	for i := 0; i < dimension; i++ {
		ret[i] = make([]int, dimension)
	}

	ringsNeeded := (dimension + 1) / 2
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

				ret[(dimension+1)/2-1][(dimension+1)/2-1] = cnt
				// place this in the center!
				return ret
			}
			cnt++
		}
	}
	return ret
}
