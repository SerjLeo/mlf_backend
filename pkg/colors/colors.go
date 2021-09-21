package colors

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

var hexRegExp = regexp.MustCompile("^#[a-fA-F0-9]{6}$")

type ColorManager interface {
	IsHEX(color string) bool
	GenerateHex() string
}

type ColorWorker struct{}

func NewColorWorker() *ColorWorker {
	return &ColorWorker{}
}

func (c *ColorWorker) IsHEX(color string) bool {
	if len(color) != 7 {
		return false
	}
	return hexRegExp.MatchString(color)
}

func (c *ColorWorker) GenerateHex() string {
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("%x%x%x", rand.Intn(256), rand.Intn(256), rand.Intn(256))
}
