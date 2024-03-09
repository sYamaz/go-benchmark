package markdown

type (
	MDText struct {
		Text string
	}
)

func NewMDText(text string) MDInline {
	return &MDText{
		Text: text,
	}
}

func (md *MDText) ToInlineString() string {
	return md.Text
}

func (md *MDText) ToInlines() MDInlineArray {
	return []MDInline{md}
}

// Then implements MDInline.
func (md *MDText) Then(inline MDInline) MDInlineArray {
	ret := MDInlineArray{}
	ret = append(ret, md, inline)
	return ret
}

// implement MDElement
func (md *MDText) ToMDString() string {
	return md.ToInlineString()
}
