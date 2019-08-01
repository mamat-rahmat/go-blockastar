package blockastar

import (
	"fmt"
	"os"
)

// TestData is pair of two coordinate start and end cell
type TestData struct {
	x1, y1, x2, y2 int
	len            float64
}

// BuildTestDataFromScen build Testdata from .scen file
func BuildTestDataFromScen(path string) []TestData {
	var (
		err                            error
		ver, num, R, C, x1, y1, x2, y2 int
		name                           string
		len                            float64
	)

	// redirect input from file`
	os.Stdin, err = os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Scanf("version %d\n", &ver)
	result := make([]TestData, 0)
	for i := 0; i < 5; i++ {
		n, _ := fmt.Scanf("%d\t%s\t%d\t%d\t%d\t%d\t%d\t%d\t%f\n", &num, &name, &R, &C, &y1, &x1, &y2, &x2, &len)
		if n != 9 {
			break
		}
		result = append(result, TestData{
			x1,
			y1,
			x2,
			y2,
			len,
		})
	}
	return result
}
