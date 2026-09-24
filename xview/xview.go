package xview

import (
	"bytes"
	phtmltpl "html/template"
	"io"
	"io/fs"
	"maps"
	"os"
	ptexttpl "text/template"
)

type Viewer interface {
	Assign(datas map[string]any) Viewer
	Render(tpl string, datas ...map[string]any) ([]byte, error)
	RenderWriter(w io.Writer, tpl string, datas ...map[string]any) (int, error)
}

type localView struct {
	isHtml bool
	viewFS fs.FS
	datas  map[string]any
}

func NewHtmlView(tplpath string) Viewer {
	return &localView{
		isHtml: true,
		viewFS: os.DirFS(tplpath),
		datas:  make(map[string]any),
	}
}

func NewTextView(tplpath string) Viewer {
	return &localView{
		isHtml: false,
		viewFS: os.DirFS(tplpath),
		datas:  make(map[string]any),
	}
}

func (lv *localView) Assign(datas map[string]any) Viewer {
	maps.Insert(lv.datas, maps.All(datas))
	return lv
}

func (lv *localView) Render(tpl string, datas ...map[string]any) ([]byte, error) {
	var output = bytes.NewBuffer(nil)
	if _, err := lv.RenderWriter(output, tpl, datas...); err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}
func (lv *localView) RenderWriter(w io.Writer, tpl string, datas ...map[string]any) (int, error) {
	for _, d := range datas {
		if len(d) > 0 {
			lv.Assign(d)
		}
	}

	if lv.isHtml {
		return lv.renderHtml(w, tpl, lv.datas)
	} else {
		return lv.renderText(w, tpl, lv.datas)
	}
}

func (lv *localView) renderHtml(w io.Writer, tpl string, data map[string]any) (int, error) {
	var (
		output   = bytes.NewBuffer(nil)
		html_tpl = phtmltpl.New("html").Funcs(lv.build_func_map())
		err      error
	)
	if html_tpl, err = html_tpl.ParseFS(lv.viewFS, "*"); err != nil {
		return 0, err
	}
	if err = html_tpl.ExecuteTemplate(output, tpl, data); err != nil {
		return 0, err
	}
	return w.Write(output.Bytes())
}
func (lv *localView) renderText(w io.Writer, tpl string, data map[string]any) (int, error) {
	var (
		output   = bytes.NewBuffer(nil)
		html_tpl = ptexttpl.New("text").Funcs(lv.build_func_map())
		err      error
	)
	if html_tpl, err = html_tpl.ParseFS(lv.viewFS, "*"); err != nil {
		return 0, err
	}
	if err = html_tpl.ExecuteTemplate(output, tpl, data); err != nil {
		return 0, err
	}
	return w.Write(output.Bytes())
}
