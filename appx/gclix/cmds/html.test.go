package cmds

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/urfave/cli/v2"
	"github.com/ynwcel/gox/xos"
)

var (
	html_test_flag_port = 8080
	htmlTestCmd         = &cli.Command{
		Name:      "html-test",
		Usage:     "html dev test run",
		UsageText: fmt.Sprintf("%s html-test [options...]", appName),
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				DefaultText: "8080",
				Value:       html_test_flag_port,
				Destination: &html_test_flag_port,
			},
		},
		Action: htmlTestAction,
	}
)

func htmlTestAction(ctx *cli.Context) error {

	svr := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", html_test_flag_port),
		Handler: new_htmlTesthandler(),
	}
	log.Printf("svr.addr=%s\n", svr.Addr)
	if err := svr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

type htmlTesthandler struct {
	viewpool *sync.Pool
}

func new_htmlTesthandler() *htmlTesthandler {
	return &htmlTesthandler{
		viewpool: &sync.Pool{
			New: func() any {
				view := gview.New()
				view.SetPath("./")
				return view
			},
		},
	}
}

func (self *htmlTesthandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		request_uri      = r.RequestURI
		request_filename = fmt.Sprintf("./%s", strings.TrimLeft(request_uri, "/"))
		view             = self.viewpool.Get().(*gview.View)
	)
	log.Printf("uri=%s,file=%s\n", request_uri, request_filename)
	if request_uri == "/" {
		if xos.PathIsFile("./index.html") {
			w.Header().Set("Location", "/index.html")
			w.WriteHeader(http.StatusTemporaryRedirect)
		} else if xos.PathIsFile("./index.htm") {
			w.Header().Set("Location", "/index.htm")
			w.WriteHeader(http.StatusTemporaryRedirect)
		} else if xos.PathIsFile("./default.html") {
			w.Header().Set("Location", "/default.html")
			w.WriteHeader(http.StatusTemporaryRedirect)
		} else if xos.PathIsFile("./default.htm") {
			w.Header().Set("Location", "/default.htm")
			w.WriteHeader(http.StatusTemporaryRedirect)
		}
		return
	} else if !xos.PathIsFile(request_filename) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - NotFound"))
	} else {
		file_mimetype, err := mimetype.DetectFile(request_filename)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "get file mimetype error:%s\n", err.Error())
			return
		}
		if strings.ToLower(file_mimetype.Extension()) == ".html" || strings.ToLower(file_mimetype.Extension()) == ".htm" {
			if content, err := view.Parse(r.Context(), request_filename); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, err.Error())
			} else {
				w.Header().Set("Content-Type", file_mimetype.String())
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, content)
			}
		} else {
			http.ServeFile(w, r, request_filename)
		}
	}
}
