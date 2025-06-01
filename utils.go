package main

import (
	"bufio"
	"context"
	"io"
	"os"
	"strings"
)

func IsFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func ReadLinesWithContext(ctx context.Context, reader io.Reader) ([]string, error) {
	std := bufio.NewScanner(reader)

	var lines []string
	for {
		select {
		case <-ctx.Done():
			return lines, ctx.Err()
		default:
			if !std.Scan() {
				if err := std.Err(); err != nil {
					return lines, err
				}
				return lines, nil
			}

			line := strings.TrimSpace(std.Text())
			if line != "" {
				lines = append(lines, line)
			}
		}
	}
}
