# BenchmarkJsonDecode

## unmarshall

|key|N|ns/op|
|---|---|---|
|byte配列からデコード-4|220459|5469|
|readerからデコード-4|109617|11007|


## decoder

|key|N|ns/op|
|---|---|---|
|byte配列からデコード-4|205557|5752|
|readerからデコード-4|113503|11154|


# BenchmarkJsonEncode

## marshall

|key|N|ns/op|
|---|---|---|
|byte配列を取得する-4|915340|1149|
|writerに書き込む-4|772330|1390|


## encoder

|key|N|ns/op|
|---|---|---|
|byte配列を取得する-4|882656|1225|
|既存のwriterに書き込む-4|995218|1183|
