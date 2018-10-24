// +build mage

package main

import "github.com/magefile/mage/sh"

func BuildExample() error {
	return sh.Run(
		"protoc",
		"-I", "/usr/local/include",
		"-I", "./example",
		"--go_out=:./example",
		"./example/example.proto",
	)
}

func BuildDefaults() error {
	return sh.Run(
		"protoc",
		"-I", "/usr/local/include",
		"-I", "./example",
		"--defaults_out=:./example",
		"./example/example.proto",
	)
}
