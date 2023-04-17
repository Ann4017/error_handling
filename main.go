package main

import (
	"errors"
	"fmt"
	"os"
)

type custom_err struct {
	len int
	std int
}

func main() {
	err := scan_error("test")
	if err != nil {
		fmt.Println(err)
	}
}

var new_err = errors.New("New error")

func wrap_error() error {
	w_err := fmt.Errorf("error: %w\n", new_err)

	return w_err
}

func unwrap_error() error {
	w_err := fmt.Errorf("error: %w\n", new_err)

	unw_err := errors.Unwrap(w_err)
	return unw_err
}

func scan_error(file_name string) error {

	file, err := os.Open(file_name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("same error: %w, %w", err, os.ErrNotExist)
		}
		return err
	}
	defer file.Close()

	return nil
}

func (c custom_err) Error() string {
	return "암호 길이가 짧습니다."
}

func register_account(id, pwd string) error {
	if len(pwd) < 10 {
		return custom_err{len(pwd), 10}
	}
	return nil
}

func conversion_err_type() {
	err := register_account("abh4017", "12345")

	if err != nil {
		if err_info, ok := err.(custom_err); ok {
			fmt.Println(err_info.len)
		}
		fmt.Println(err)
	}
}
