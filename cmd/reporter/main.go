package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sYamaz/mdstruct"
	"golang.org/x/tools/benchmark/parse"
)

type (
	OSExitError struct {
		Internal error
		ExitCode int
	}

	BenchmarkGroup struct {
		Depth       int               `json:"depth"`
		Key         string            `json:"key"`
		ChildGroups []BenchmarkGroup  `json:"childlen,omitempty"`
		Value       []parse.Benchmark `json:"value,omitempty"` // b.Runで同じ名前を指定する場合があるかもしれない
	}
)

// frags
var (
	format  = flag.String("format", "json", "出力形式を指定します json / md。default=json")
	outFile = flag.String("out", "", "出力ファイルを指定します")

	formatMap = map[string]func(groups []BenchmarkGroup) ([]byte, error){
		"json": convertToPrettyJson,
		"md":   convertToMarkdown,
	}
)

// Error implements error.
func (o *OSExitError) Error() string {
	return fmt.Sprintf("exitCode=%v, internal=%v", o.ExitCode, o.Internal)
}

func usage() {
	usageText := "Usage: \n"
	usageText = usageText + fmt.Sprintln("golangのベンチマーク出力 (e.g. `go test -bench . > log.txt`)したテキストファイルを読み込んで")
	usageText = usageText + fmt.Sprintln("行ごとに値を読み取ります。")
	usageText = usageText + fmt.Sprintln("その後、Nameをdelimiterで分割し木構造に変換し、出力します。")
	usageText = usageText + fmt.Sprintln("")
	fmt.Fprintln(flag.CommandLine.Output(), usageText)
	flag.PrintDefaults()
}

func getWriter(outFile *string) (io.Writer, error) {
	if *outFile == "" {
		return log.Writer(), nil
	}

	f, err := os.Create(*outFile)
	if err != nil {
		return nil, &OSExitError{
			ExitCode: 1,
			Internal: err,
		}
	}

	return f, nil
}

func init() {
	log.SetPrefix("")
	log.SetFlags(0)

	flag.Usage = usage
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		usage()
		os.Exit(2)
	}

	writer, err := getWriter(outFile)
	if err != nil {
		if exitErr, ok := err.(*OSExitError); ok {
			log.Println(exitErr.Internal)
			os.Exit(exitErr.ExitCode)
		}
		log.Fatal(err)
	}

	file, err := os.Open(flag.Arg(0))
	defer func() {
		file.Close()
	}()
	if err != nil {
		log.Fatal(err)
	}

	// readfile and convert to []Benchmark
	arr, err := parseToBenchResultArray(file)
	if err != nil {
		if exitErr, ok := err.(*OSExitError); ok {
			log.Println(exitErr.Internal)
			os.Exit(exitErr.ExitCode)
		}
		log.Fatal(err)
	}

	// group Benchmarks by name
	group := groupByName(arr, "/", 0)

	// output
	f := formatMap[*format]
	b, err := f(group)
	if err != nil {
		if exitErr, ok := err.(*OSExitError); ok {
			log.Println(exitErr.Internal)
			os.Exit(exitErr.ExitCode)
		}
		log.Fatal(err)
	}

	writer.Write(b)

	os.Exit(0)
}

func parseToBenchResultArray(reader io.Reader) ([]parse.Benchmark, error) {
	set, err := parse.ParseSet(reader)
	if err != nil {
		return nil, &OSExitError{
			Internal: err,
			ExitCode: 1,
		}
	}

	retArray := []parse.Benchmark{}

	for _, v := range set {
		for _, vEle := range v {
			retArray = append(retArray, *vEle)
		}
	}

	return retArray, nil
}

func groupByName(benchmarks []parse.Benchmark, nameSplitter string, depth int) []BenchmarkGroup {
	ret := []BenchmarkGroup{}

	// 枝
	branches := map[string][]parse.Benchmark{}
	// 葉
	leafs := map[string]*BenchmarkGroup{}

	for _, b := range benchmarks {
		keys := strings.Split(b.Name, nameSplitter)
		// Nameが分割できない場合benchmark末端（葉）と判断する
		if len(keys) == 1 {
			if _, ok := leafs[keys[0]]; !ok {
				leafs[keys[0]] = &BenchmarkGroup{
					Depth:       depth,
					Key:         keys[0],
					Value:       []parse.Benchmark{},
					ChildGroups: []BenchmarkGroup{},
				}
			}
			leafs[keys[0]].Value = append(leafs[keys[0]].Value, b)
			continue
		}

		// それ以外の場合、枝と判断する
		b.Name = strings.Join(keys[1:], nameSplitter)
		branches[keys[0]] = append(branches[keys[0]], b)
	}

	// 葉を確定する
	for _, v := range leafs {
		ret = append(ret, *v)
	}

	// 枝を再起的にグルーピングする
	for k, v := range branches {
		group := groupByName(v, nameSplitter, depth+1)
		ret = append(ret, BenchmarkGroup{
			Depth:       depth,
			Key:         k,
			ChildGroups: group,
			Value:       nil,
		})
	}

	return ret
}

// func convertToJson(groups []BenchmarkGroup) ([]byte, error) {
// 	return json.Marshal(groups)
// }

func convertToPrettyJson(groups []BenchmarkGroup) ([]byte, error) {
	return json.MarshalIndent(groups, "", "  ")
}

// func convertToXml(groups []BenchmarkGroup) ([]byte, error) {
// 	return xml.Marshal(groups)
// }

// func convertToPrettyXml(groups []BenchmarkGroup) ([]byte, error) {
// 	return xml.MarshalIndent(groups, "", "  ")
// }

func convertToMarkdown(groups []BenchmarkGroup) ([]byte, error) {
	doc := mdstruct.NewMDDocument(nil)

	for _, group := range groups {
		convertToMarkdownCore(doc, group)
	}

	str := doc.ToMDString()

	return []byte(str), nil
}

func convertToMarkdownCore(doc *mdstruct.MDDocument, group BenchmarkGroup) {
	// 見出し
	headline := mdstruct.NewMDHeadline(group.Depth+1, mdstruct.NewMDText(group.Key))
	doc.Add(headline)

	// テーブル
	rows := []mdstruct.MDTableRow{}

	for _, ch := range group.ChildGroups {
		// childrenが葉の場合、テーブルに突っ込む
		if len(ch.ChildGroups) == 0 { // 葉の場合
			for _, v := range ch.Value {
				rows = append(rows, *mdstruct.NewMDTableRow(
					[]mdstruct.MDTableCell{
						*mdstruct.NewMDTableCell(mdstruct.NewMDText(v.Name)),
						*mdstruct.NewMDTableCell(mdstruct.NewMDText(strconv.Itoa(v.N))),
						*mdstruct.NewMDTableCell(mdstruct.NewMDText(strconv.FormatFloat(v.NsPerOp, 'f', -1, 64))),
					},
				))
			}
			continue
		}

		// 葉じゃない場合
		convertToMarkdownCore(doc, ch)

	}

	if len(rows) != 0 {
		table := mdstruct.NewMDTable(
			// header
			*mdstruct.NewMDTableRow(
				[]mdstruct.MDTableCell{
					*mdstruct.NewMDTableCell(mdstruct.NewMDText("key")),
					*mdstruct.NewMDTableCell(mdstruct.NewMDText("N")),
					*mdstruct.NewMDTableCell(mdstruct.NewMDText("ns/op")),
				},
			),
			rows,
		)

		doc.Add(table)
	}
}
