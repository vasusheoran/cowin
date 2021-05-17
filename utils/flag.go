package utils

import (
	"fmt"
	"strings"
)

type DuplicateStringFlag []string

func (ff *DuplicateStringFlag) String() string {
	return fmt.Sprintf("%s", *ff)
}

func (ff *DuplicateStringFlag) Set(value string) error {
	fmt.Sprintf("%s\n", value)
	*ff = append(*ff, strings.TrimSpace(value))
	return nil
}

//
//type DuplicateIntFlag []int
//
//func (ff *DuplicateIntFlag) String() string {
//	return fmt.Sprintf("%d", *ff)
//}
//
//func (ff *DuplicateIntFlag) Set(value string) error {
//	tmp, err := strconv.Atoi(value)
//	if err != nil {
//		*ff = append(*ff, -1)
//	} else {
//		*ff = append(*ff, tmp)
//	}
//	return nil
//}
