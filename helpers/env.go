package helpers

import (
	"fmt"
	"github.com/subosito/gotenv"
)

func LoadEnv() {
	err = gotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
}
