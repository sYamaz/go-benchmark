// table-row is not markdown element
// it is only table content

package markdown

import "strings"

type (
	MDTableRow struct {
		cols []MDTableCell
	}
)

func NewMDTableRow(cols []MDTableCell) *MDTableRow {
	c := cols
	if c == nil {
		c = []MDTableCell{}
	}

	return &MDTableRow{
		cols: c,
	}
}

func (md *MDTableRow) toInlineString() string {
	strs := []string{}
	for _, c := range md.cols {
		strs = append(strs, c.toString())
	}

	return "|" + strings.Join(strs, "|") + "|"
}

func (md *MDTableRow) toString() string {
	return md.toInlineString()
}
