package main

import "fmt"

const SDG = "Soli Deo Gloria"

func init() {
	fmt.Println(SDG)
}

// Run - is going to be responsible for
// the instantiation and staretup of our go application
func Run() error {
	fmt.Println("Starting up our application")
	return nil
}

func main() {
	fmt.Println("Hello from Docker within VSCode and Wsl2")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
