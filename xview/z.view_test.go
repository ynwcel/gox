package xview

import "testing"

func TestViewHtml(t *testing.T) {
	datas := map[string]any{
		"content": "<h3>TestViewHtml</h3>",
	}
	if r, err := NewHtmlView("./z.test.htmls").Render("index.html", datas); err != nil {
		t.Fatal(err)
	} else {
		t.Log("output:", string(r))
	}
}

func TestViewText(t *testing.T) {
	datas := map[string]any{
		"content": "<h3>TestViewText</h3>",
	}
	if r, err := NewTextView("./z.test.htmls").Render("index.html", datas); err != nil {
		t.Fatal(err)
	} else {
		t.Log("output:", string(r))
	}
}
