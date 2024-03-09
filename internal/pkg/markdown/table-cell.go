// table-cell is not markdown element
// it is only table content
package markdown

type (
	MDTableCell struct {
		contents []MDInline
	}
)

func NewMDTableCell(contents []MDInline) *MDTableCell {
	c := contents
	if c == nil {
		c = []MDInline{}
	}

	return &MDTableCell{
		contents: c,
	}
}

func (md *MDTableCell) toInlineString() string {
	return joinMDInlines(md.contents)
}

func (md *MDTableCell) toString() string {
	return md.toInlineString()
}
