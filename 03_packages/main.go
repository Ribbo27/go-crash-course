package main

import (
	"fmt"
	"math"
	"github.com/Ribbo27/go-crash-course/03_packages/strUtil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strUtil.Reverse("hello"))
}