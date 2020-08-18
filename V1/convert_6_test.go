package V1

import (
	"bytes"
	"fmt"
	"time"
)

func main6() {

	closed := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("3333")
		close(closed)
	}()

	select {
	case e, ok := <-closed:
		fmt.Println("====> ", e, ok)
	}

	fmt.Println("22222")

	// fmt.Println(convert("LEETCODEISHIRING", 3))

	// fmt.Println(convert("LEETCODEISHIRING", 4))

}

func convert(s string, numRows int) string {

	if numRows == 1 || numRows >= len(s) {
		return s
	}
	resSlice := make([]string, numRows)

	loc := 0
	down := false

	for i := 0; i < len(s); i++ {
		resSlice[loc] += string(s[i : i+1])

		if loc == 0 || loc == numRows-1 {
			down = !down
		}

		if down {
			loc++
		} else {
			loc += -1
		}

	}

	buf := bytes.NewBuffer([]byte{})

	for _, v := range resSlice {
		buf.WriteString(v)
	}

	return buf.String()

}
