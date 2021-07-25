package main

import (
	"database/sql"
	"errors"
	"fmt"

	xerrors "github.com/pkg/errors"
)

type User struct {
	Name string
	Age  int
}

func getUser() (User, error) {
	err := sql.ErrNoRows
	return User{}, xerrors.Wrapf(err, "getUser error")
}

func main() {
	user, err := getUser()

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("Original error:\n%T %v\n\n", xerrors.Cause(err), err)
			fmt.Printf("stack trace:\n%+v\n", err)
			return
		}
	}

	fmt.Println(user)
}
