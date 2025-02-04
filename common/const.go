package common

import "fmt"

func Recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from ", r)
	}
}

type TokenPayload struct {
	UId   int    `json:"uid"`
	URole string `json:"role"`
}

func (p TokenPayload) GetUId() int {
	return p.UId
}

func (p TokenPayload) GetURole() string {
	return p.URole
}
