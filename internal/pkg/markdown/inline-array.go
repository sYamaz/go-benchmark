package markdown

type (
	MDInlineArray []MDInline
)

func (md *MDInlineArray) Then(inline MDInline) MDInlineArray {
	ret := []MDInline{}

	ret = append(ret, *md...)
	ret = append(ret, inline)

	return ret
}
