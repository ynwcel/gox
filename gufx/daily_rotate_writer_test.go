package gufx

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestDailyRotateWriter_WriteAndRotate(t *testing.T) {
	log_file := fmt.Sprintf("%s/app.log", t.TempDir())

	writer := NewDailyRotateWriter(log_file)

	// 写入一条日志
	logLine1 := "hello world 1\n"
	if _, err := writer.Write([]byte(logLine1)); err != nil {
		t.Fatalf("write failed: %v", err)
	}

	// 检查当天日志文件是否存在
	today := time.Now().Format("20060102")
	log_file_ext := filepath.Ext(log_file)
	log_file_base := strings.TrimSuffix(log_file, log_file_ext)
	logFile := log_file_base + "." + today + log_file_ext

	data, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}
	if !strings.Contains(string(data), "hello world 1") {
		t.Errorf("log content mismatch, got: %s", string(data))
	}

	// 模拟跨天：手动改 currDate，再写一条日志
	drw := writer.(*DailyRotateWriter)
	yesterday := time.Now().AddDate(0, 0, -1).Format("20060102")
	drw.currDate = yesterday // 强制让 writer 以为是昨天

	logLine2 := "hello world 2\n"
	if _, err := writer.Write([]byte(logLine2)); err != nil {
		t.Fatalf("write failed on rotation: %v", err)
	}

	// 确认今天文件包含第二条日志
	data2, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("failed to read log file after rotation: %v", err)
	}
	if !strings.Contains(string(data2), "hello world 2") {
		t.Errorf("log content after rotation mismatch, got: %s", string(data2))
	}
	drw.file.Close()
}

func BenchmarkDailyRotateWriter_Write(b *testing.B) {
	log_file := fmt.Sprintf("%s/app.log", b.TempDir())

	writer := NewDailyRotateWriter(log_file)

	data := []byte("benchmark log line\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := writer.Write(data); err != nil {
			b.Fatalf("write failed: %v", err)
		}
	}
	drw := writer.(*DailyRotateWriter)
	drw.file.Close()
}
