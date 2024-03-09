package markdown

type (
	MDEmphasis struct {
		contents []MDInline
	}
)

func NewMDEmphasis(contents []MDInline) MDInline {
	c := contents
	if c == nil {
		c = []MDInline{}
	}
	return &MDEmphasis{
		contents: c,
	}
}

func (md *MDEmphasis) ToInlineString() string {
	return "**" + joinMDInlines(md.contents) + "**"
}

// implement MDElement
func (md *MDEmphasis) ToMDString() string {
	return md.ToInlineString()
}

// Then implements MDInline.
func (md *MDEmphasis) Then(inline MDInline) MDInlineArray {
	ret := MDInlineArray{}
	ret = append(ret, md, inline)
	return ret
}

func (md *MDEmphasis) ToInlines() MDInlineArray {
	return []MDInline{md}
}
