# benchmarkを実行してlog.txtに結果を出力する
go test -bench . ./benchs/... > ./tmp/log.txt

# markdownに変換してresultsに保存する
go run ./cmd/reporter/main.go -format md -out ./results/log.md ./tmp/log.txt