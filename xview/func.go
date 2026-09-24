package xview

import (
	"bytes"
	phtmltpl "html/template"
	"maps"
)

func (lv *localView) build_func_map() map[string]any {
	return map[string]any{
		"include": lv.build_include_func(),
	}
}

func (lv *localView) build_include_func() any {
	var render = func(file string, datas ...map[string]any) (string, error) {
		var (
			output   = bytes.NewBuffer(nil)
			lv_datas = lv.datas
		)
		for _, d := range datas {
			if len(d) > 0 {
				maps.Insert(lv_datas, maps.All(d))
			}
		}
		if lv.isHtml {
			if _, err := lv.renderHtml(output, file, lv_datas); err != nil {
				return "", err
			}
		} else {
			if _, err := lv.renderText(output, file, lv_datas); err != nil {
				return "", err
			}
		}
		return output.String(), nil
	}
	if lv.isHtml {
		return func(file string, datas ...map[string]any) (phtmltpl.HTML, error) {
			if content, err := render(file, datas...); err != nil {
				return "", err
			} else {
				return phtmltpl.HTML(content), nil
			}
		}
	} else {
		return render
	}
}
