package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/revel/revel"
)

type Color int
type App struct {
	*revel.Controller
}

const (
	Red  Color = 1
	Blue       = 2
)

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Show() revel.Result {
	return c.Render()
}

func (c App) Form() revel.Result {
	return c.Render()
}

func f(ctx context.Context, str []byte) error {
	var m map[string]string
	m["A"] = "default"                             // main.go:13:2: assignment to nil map (SA5000)
	if err := json.Unmarshal(str, m); err != nil { // main.go:14:32: json.Unmarshal expects to unmarshal into a pointer, but the provided value is not a pointer (SA1014)
		return err
	}
	l := *&m // main.go:17:7: *&x will be simplified to x. It will not copy x. (SA4001)
	l["A"] = ""
	if strings.ToLower(m["A"]) == strings.ToLower(m["X"]) { // main.go:19:5: should use strings.EqualFold instead (SA6005)
		fmt.Println(fmt.Sprintf("%v -> %s", m, l)) // main.go:27:3: should use fmt.Printf instead of fmt.Println(fmt.Sprintf(...)) (but don't forget the newline) (S1038)
	}
	return nil
}
