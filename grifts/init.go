package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/kkarimi/buffalo-react/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
