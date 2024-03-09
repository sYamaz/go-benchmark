package markdown

import (
	"errors"
	"strings"
)

type (
	MDHeadline struct {
		// 1 ~ 6
		level    int
		contents []MDInline
	}
)

func NewMDHeadline(level int, content []MDInline) *MDHeadline {
	c := content
	if c == nil {
		c = []MDInline{}
	}

	if level < 1 || 6 < level {
		panic(errors.New("level must be 1 ~ 6"))
	}

	return &MDHeadline{
		level:    level,
		contents: c,
	}
}

func (md *MDHeadline) ToBlockString() string {
	return strings.Repeat("#", md.level) + " " + joinMDInlines(md.contents)
}

// implement MDElement
func (md *MDHeadline) ToMDString() string {
	return md.ToBlockString() + "\n"
}
