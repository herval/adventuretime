package twitter

import (
	"strings"

	"github.com/herval/adventuretime/engine"
)

type TweetParser struct {
	Name string
}

var standardParser = engine.StandardParser{}

func (self *TweetParser) ParseCommand(cmd string) engine.Command {
	sanitized := strings.Trim(
		strings.Replace(cmd, "@"+self.Name, "", -1),
		" ",
	)
	return standardParser.ParseCommand(sanitized)
}
