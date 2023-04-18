package main

import (
	"errors"
	"fmt"
	"os"
)

type password_err struct {
	len int
	std int
}

var (
	open_err     = errors.New("file open error")
	open_err_get = fmt.Errorf("발생한 오류: %w\n", open_err)
)

func open_panic_test(file_name string) error {
	file, err := os.Open(file_name)
	if err != nil {
		panic(open_err_get)
	}
	defer file.Close()

	return nil
}

func open_test(file_name string) error {
	file, err := os.Open(file_name)
	if err != nil {
		if err == open_err {
			//
		}
		return open_err_get
	}
	defer file.Close()

	return nil
}

func (c password_err) Error() string {
	return "암호 길이가 짧습니다."
}

func register_account(id, pwd string) error {
	if len(pwd) < 10 {
		return password_err{len(pwd), 10}
	}
	return nil
}

func main() {
	//custom err
	err := register_account("id", "pwd")
	if err != nil {
		if err_info, ok := err.(password_err); ok {
			fmt.Printf("%s len: %d, std: %d\n", err_info, err_info.len, err_info.std)
		}
	}

	//panic recover
	func() {
		err = open_panic_test("test")
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recover from: %v", r)
			}
		}()
	}()
}

// var new_err = errors.New("New error")

// func wrap_error() error {
// 	w_err := fmt.Errorf("error: %w\n", new_err)

// 	return w_err
// }

// func unwrap_error() error {
// 	w_err := fmt.Errorf("error: %w\n", new_err)

// 	unw_err := errors.Unwrap(w_err)
// 	return unw_err
// }

// func scan_error(file_name string) error {

// 	file, err := os.Open(file_name)
// 	if err != nil {
// 		if errors.Is(err, os.ErrNotExist) {
// 			return fmt.Errorf("same error: %w, %w", err, os.ErrNotExist)
// 		}
// 		return err
// 	}
// 	defer file.Close()

// 	return nil
// }

// func conversion_err_type() {
// 	err := register_account("abh4017", "12345")

// 	if err != nil {
// 		if err_info, ok := err.(password_err); ok {
// 			fmt.Println(err_info.len)
// 		}
// 		fmt.Println(err)
// 	}
// }
