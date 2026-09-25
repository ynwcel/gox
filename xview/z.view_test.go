package xview

import (
	"os"
	"testing"
)

func TestViewHtml(t *testing.T) {
	datas := map[string]any{
		"content": "<h3>TestViewHtml</h3>",
	}
	if r, err := NewHtmlView(os.DirFS("./z.test.htmls")).Render("index.html", datas); err != nil {
		t.Fatal(err)
	} else {
		t.Log("output:", string(r))
	}
}

func TestViewText(t *testing.T) {
	datas := map[string]any{
		"content": "<h3>TestViewText</h3>",
	}
	if r, err := NewTextView(os.DirFS("./z.test.htmls")).Render("index.html", datas); err != nil {
		t.Fatal(err)
	} else {
		t.Log("output:", string(r))
	}
}
