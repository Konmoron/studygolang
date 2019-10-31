package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	uid := uuid.New()
	fmt.Printf("uid, type: %T, value: %s", uid, uid.String())
}
