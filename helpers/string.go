package helpers

import "strconv"

func StrToUint(str string) (n uint, err error) {

	unsigned_int, err := strconv.ParseUint(str, 10, 32)

	return uint(unsigned_int), err
}
