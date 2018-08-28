package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/mile_high/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
