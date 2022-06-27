package validator

import (
	"errors"
	"strconv"
)

func CustomConvertNumber(inputString string) (int, error) {
	convertToNumber, err := strconv.Atoi(inputString)
	if err != nil {
		if ok := errors.Is(err.(*strconv.NumError).Err, strconv.ErrSyntax); ok {
			return convertToNumber, errors.New("The value is invalid format, It should be number.")
		} else if ok := errors.Is(err.(*strconv.NumError).Err, strconv.ErrRange); ok {
			return convertToNumber, errors.New("The value is invalid value, It is out of range.")
		} else {
			return convertToNumber, err
		}
	}
	return convertToNumber, nil
}
