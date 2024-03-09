package markdown

import "strings"

type (
	MDTable struct {
		header MDTableRow
		body   []MDTableRow
	}
)

func NewMDTable(header MDTableRow, rows []MDTableRow) *MDTable {
	rs := rows
	if rs == nil {
		rs = []MDTableRow{}
	}

	return &MDTable{
		header: header,
		body:   rs,
	}
}

func (md *MDTable) ToMDString() string {
	return md.ToMDBlockString() + "\n"
}

func (md *MDTable) ToMDBlockString() string {

	colsCount := len(md.header.cols)
	divider := "|" + strings.Repeat("---|", colsCount)

	strs := []string{md.header.toString(), divider}

	for _, r := range md.body {
		strs = append(strs, r.toString())
	}

	return strings.Join(strs, "\n")
}
