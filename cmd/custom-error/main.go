/*
Go programs express error state with error values.
The error type is a built-in interface similar to fmt.Stringer:

type error interface {
	Error() string
}

自分なりのエラーを出力したい場合はError()メソッドを作ってやれば良い。
*/

package main

import "fmt"

// UserNotFound is a struct that has "Username".
type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	err := &UserNotFound{Username: "mike"}
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
