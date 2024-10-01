package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// 入力
func Input() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("出力するファイル名を入力してください: ")

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
		return ""
	}

	fileName := scanner.Text()
	if fileName == "" {
		fmt.Fprintln(os.Stderr, "ファイル名が空です。")
		return ""
	}

	return fileName
}

// オプション
func Option() bool {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("行番号を表示しますか？(y/n): ")

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
		return false
	}

	option := scanner.Text()
	if option == "" {
		fmt.Fprintln(os.Stderr, "オプションが空です。")
		return false
	}

	if option == "y" {
		return true
	}
	return false
}

// ディレクトリウォーク
func WalkDir(rootDir, fileName string) (string, error) {
	var foundPath string
	found := false

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Base(path) == fileName {
			foundPath = path
			found = true
			// ファイルが見つかったので探索を終了
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if !found {
		return "", fmt.Errorf("ファイルが見つかりませんでした")
	}

	return foundPath, nil
}

// 表示
func PrintFile(filePath string, hasOption bool) error {
	file, err := os.Open(filePath)
	line := 0
	if err != nil {
		return fmt.Errorf("ファイルを開くことができませんでした: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if(hasOption) {
			fmt.Print(line, ": ")
			line++
		}
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ファイルの読み込みエラー: %w", err)
	}

	return nil
}

func main() {

	// ファイル名を入力
	fileName := Input()
	if fileName == "" {
		return
	}

	// 行番号を表示するかどうかのオプション
	option := Option()

	// 現在のディレクトリを取得
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "現在のディレクトリを取得できませんでした:", err)
		return
	}

	// ディレクトリウォーク
	foundPath, err := WalkDir(dir, fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// ファイルの表示
	if err := PrintFile(foundPath, option); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
