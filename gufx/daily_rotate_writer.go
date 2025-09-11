package gufx

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type DailyRotateWriter struct {
	mu       sync.Mutex
	wfile    string
	wext     string
	file     *os.File
	currDate string
}

// NewDailyRotateWriter 构造函数
func NewDailyRotateWriter(file string) io.Writer {
	var (
		fileExt  = filepath.Ext(file)
		fileBase = strings.TrimSuffix(file, fileExt)
	)

	return &DailyRotateWriter{
		wfile: fileBase,
		wext:  fileExt,
	}
}

func (w *DailyRotateWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	today := time.Now().Format("20060102")
	if w.file == nil || today != w.currDate {
		if err := w.rotate(today); err != nil {
			return 0, err
		}
	}
	return w.file.Write(p)
}

func (w *DailyRotateWriter) rotate(date string) error {
	if w.file != nil {
		_ = w.file.Close()
	}
	if err := os.MkdirAll(filepath.Dir(w.wfile), 0755); err != nil {
		return err
	}
	filename := w.wfile + "." + date + w.wext
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	w.file = f
	w.currDate = date
	return nil
}
