package collatzconjecture

import "errors"

//CollatzConjecture returns an int and error
func CollatzConjecture(input int) (step int, err error) {
	if input < 1 {
		return 0, errors.New("True")
	} else if input == 1 {
		return 0, nil
	}

	for {
		if input == 1 {
			break
		} else if input%2 == 0 {
			input = input / 2
			step++
		} else if input%2 != 0 {
			input = input*3 + 1
			step++
		}
	}

	return
}
