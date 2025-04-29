package main

import (
	"errors"
	"fmt"
)

//type myError struct {
//	message string
//}
//
//func (m *myError) Error() string {
//	return fmt.Sprintf("Error: %s", m.message)
//}
//
//func eprocess() error {
//	return &myError{"Custom error message"}
//}
//
//func readConfig() error {
//	return errors.New("Config error")
//}
//
//func readData() error {
//	err := readConfig()
//	if err != nil {
//		return fmt.Errorf("Error reading config: %w", err)
//	}
//	return nil
//}

type customError struct {
	code    int
	message string
	er      error
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error: %d: %s, %v\n", e.code, e.message, e.er)
}

func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong!",
			er:      err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("Internal error")
}

func main() {
	//err1 := eprocess()
	//if err1 != nil {
	//	fmt.Println(err1)
	//	return
	//}
	//fmt.Println("Data processed successfully")

	//if err := readData(); err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}

	err := doSomething()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Operation successful")
}
