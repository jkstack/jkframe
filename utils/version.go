package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Version semantic version
type Version [3]int

// ParseVersion parse semantic version string
func ParseVersion(str string) (Version, error) {
	var ret Version
	tmp := strings.SplitN(str, ".", 3)
	if len(tmp) != 3 {
		return ret, errors.New("invalid version")
	}
	n, err := strconv.ParseInt(tmp[0], 10, 64)
	if err != nil {
		return ret, errors.New("invalid major version")
	}
	ret[0] = int(n)
	n, err = strconv.ParseInt(tmp[1], 10, 64)
	if err != nil {
		return ret, errors.New("invalid minor version")
	}
	ret[1] = int(n)
	n, err = strconv.ParseInt(tmp[2], 10, 64)
	if err != nil {
		return ret, errors.New("invalid patch version")
	}
	ret[2] = int(n)
	return ret, nil
}

// String show semantic version
func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v[0], v[1], v[2])
}

// Greater validate src version is greater than dst version
func (src Version) Greater(dst Version) bool {
	for i := 0; i < len(src); i++ {
		if src[i] > dst[i] {
			return true
		}
	}
	return false
}

// Greater validate src version is greater or equal dst version
func (src Version) GreaterEqual(dst Version) bool {
	for i := 0; i < len(src); i++ {
		if src[i] > dst[i] {
			return true
		} else if src[i] < dst[i] {
			return false
		}
	}
	return true
}

// Greater validate src version is less than dst version
func (src Version) Less(dst Version) bool {
	for i := 0; i < len(src); i++ {
		if src[i] < dst[i] {
			return true
		}
	}
	return false
}

// Greater validate src version is less or equal dst version
func (src Version) LessEqual(dst Version) bool {
	for i := 0; i < len(src); i++ {
		if src[i] < dst[i] {
			return true
		} else if src[i] > dst[i] {
			return false
		}
	}
	return true
}

// Greater validate src version is equal to dst version
func (src Version) Equal(dst Version) bool {
	for i := 0; i < len(src); i++ {
		if src[i] != dst[i] {
			return false
		}
	}
	return true
}
