package function

import (
	"crypto/md5"
	"fmt"
)

func CheckHash(check string) bool {
	hashStandard := "ac85e83127e49ec42487f272d9b9db8b"
	hashShadow := "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	hashThinkertoy := "5d0f26221815d2cb76574c5615eaab54"
	if check == hashStandard {
		return true
	}
	if check == hashShadow {
		return true
	}
	if check == hashThinkertoy {
		return true
	}
	return false
}

func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}
