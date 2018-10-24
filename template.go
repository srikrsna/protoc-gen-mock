package main

import (
	"io"
	"strings"
	"text/template"

	"github.com/gobuffalo/packr"
)

var T *template.Template

func init() {
	T = template.New("defaults")
	if err := packr.NewBox("./templates").Walk(func(s string, file packr.File) error {
		var sb strings.Builder
		if _, err := io.Copy(&sb, file); err != nil {
			return err
		}
		var err error
		T, err = T.Parse(sb.String())
		return err
	}); err != nil {
		panic(err)
	}
}
