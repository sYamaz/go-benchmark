# BenchmarkJsonEncode

## encoder

|key|N|ns/op|
|---|---|---|
|byte配列を取得する-4|919708|1213|
|既存のwriterに書き込む-4|1023951|1168|

## marshall

|key|N|ns/op|
|---|---|---|
|byte配列を取得する-4|907201|1137|
|writerに書き込む-4|910286|1266|

# BenchmarkJsonDecode

## decoder

|key|N|ns/op|
|---|---|---|
|byte配列からデコード-4|199902|5684|
|readerからデコード-4|110485|11206|

## unmarshall

|key|N|ns/op|
|---|---|---|
|byte配列からデコード-4|221516|5462|
|readerからデコード-4|115298|11714|
