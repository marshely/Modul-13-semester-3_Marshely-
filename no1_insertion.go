package main
import (
	"fmt"
	"sort"
)

func checkJarak(arr []int) string {
	if len(arr) < 2 {
		return "Data berjarak tidak tetap"
	}

	gap := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != gap {
			return "Data berjarak tidak tetap"
		}
	}
	return fmt.Sprintf("Data berjarak %d", gap)
}
func main() {
	var input []int
	var num int

	for {
	fmt.Scan(&num)
	if num < 0 {
		break
	}
	input = append(input, num)
}

	sort.Ints(input)

	status := checkJarak(input)

	fmt.Println("Keluaran:")
	for i, val := range input {
		if i == len(input)-1 {
			fmt.Printf("%d\n", val)
		} else {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println(status)
}