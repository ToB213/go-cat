「cat」コマンドのようなもの．ただそれだけ．


## 使用方法

### 1. ソースコードを実行

以下のコマンドでプログラムを直接実行できます．

```
go run .
```
### 2. バイナリをビルドして実行

プログラムをビルドして実行ファイルを生成することもできます．

```
go build
./cat
```
## 実行例

以下は、main.go ファイルを出力する際の例です．ファイル名を入力し，行番号を表示するかどうかを選択します．

```
$ go run .
出力するファイル名を入力してください: main.go
行番号を表示しますか？(y/n): y
0: package main
1:
2: import (
3: 	"bufio"
4: 	"fmt"
5: 	"os"
6: 	"path/filepath"
7: )
8:
...
```
