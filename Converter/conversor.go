package converter
//func to convert the number to meter. Use os Args to error handling
import (
	"fmt"
	"os"
	"strconv"
)

func Conversor() {
	arg := os.Args[1]
	num, err := strconv.ParseFloat(arg, 64)
	meter := num * 0.3048
	if err != nil {
		fmt.Printf("Error: %q is not a number", arg)
	} else {
		fmt.Printf("%g Feets to  %g is meters\n", num, meter)
	}

}
