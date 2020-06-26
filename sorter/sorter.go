package sorter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

func Run(config *Config) error {
	lines, err := readLines(config.FileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sort.Strings(lines)
	sortedFileName := addSuffix(config.FileName, config.Suffix)
	err = writeLines(sortedFileName, lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}

func addSuffix(name string, suffix string) string {
	return fmt.Sprintf("%s-%s%s", strings.TrimSuffix(name, path.Ext(name)),
		suffix, filepath.Ext(name))
}

func readLines(file string) (lines []string, err error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		const delim = '\n'
		line, err := r.ReadString(delim)
		if err == nil || len(line) > 0 {
			if err != nil {
				line += string(delim)
			}
			lines = append(lines, line)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return lines, nil
}

func writeLines(file string, lines []string) (err error) {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	for _, line := range lines {
		_, err := w.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}
