package main

import (
	"NectarPin/internal/Init"
)

func main() {
	Init.Conf()
	Init.Logger()
	Init.Gorm()
}
