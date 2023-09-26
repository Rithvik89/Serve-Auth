package utils

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	fmt.Println("DB name pls ---> ", GetDB())
	if GetDB() != "auth_db" {
		t.Errorf("Got %s, Correct %s", GetDB(), "auth_db")
	}
}
