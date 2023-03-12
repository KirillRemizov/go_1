package main

import "fmt"

func main() {
	fmt.Println("Hello GO!")

	a, err := firstFunc("a", 2)
	if err != nil {
		// TODO: handle err
	}

	a = 5

	fmt.Println(a)

}

func firstFunc(in string, in2 int) (a int, err error) {
	return 1, nil
}
