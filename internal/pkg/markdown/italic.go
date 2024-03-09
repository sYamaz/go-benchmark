package markdown

type (
	MDItalic struct {
		contents []MDInline
	}
)

func NewMDItalic(contents []MDInline) MDInline {
	c := contents
	if c == nil {
		c = []MDInline{}
	}
	return &MDItalic{
		contents: c,
	}
}

func (md *MDItalic) ToInlineString() string {
	return "*" + joinMDInlines(md.contents) + "*"
}

// Then implements MDInline.
func (md *MDItalic) Then(inline MDInline) MDInlineArray {
	ret := MDInlineArray{}
	ret = append(ret, md, inline)
	return ret
}

// implement MDElement
func (md *MDItalic) ToMDString() string {
	return md.ToInlineString()
}

func (md *MDItalic) ToInlines() MDInlineArray {
	return []MDInline{md}
}
