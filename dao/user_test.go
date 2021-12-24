package dao

import (
	"fmt"
	"testing"
)

func TestGetUserByPhone(t *testing.T) {
	user, _ := GetUserByPhone("13558619307")
	fmt.Println("user = ", user)
}