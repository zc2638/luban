/**
 * Created by zc on 2020/11/26.
 */
package printer

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

const (
	tabwriterMinWidth = 6
	tabwriterWidth    = 4
	tabwriterPadding  = 3
	tabwriterPadChar  = ' '
)

func New() *tabwriter.Writer {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, 0)
	return w
}

func NewTab(header ...string) *Tab {
	return &Tab{Header: header}
}

type Tab struct {
	Header []string
	Data   [][]interface{}
}

func (t *Tab) Add(vs ...string) {
	var data []interface{}
	for _, v := range vs {
		data = append(data, v)
	}
	t.Data = append(t.Data, data)
}

func (t *Tab) Print() {
	var num int
	if t.Header != nil {
		num = len(t.Header)
	} else if len(t.Data) > 0 {
		num = len(t.Data[0])
	} else {
		return
	}
	formats := make([]string, 0, num)
	for i := 0; i < num; i++ {
		formats = append(formats, "%s")
	}
	format := strings.Join(formats, "\t") + "\n"
	w := New()

	if t.Header != nil {
		var header []interface{}
		for _, h := range t.Header {
			header = append(header, h)
		}
		fmt.Fprintf(w, format, header...)
	}
	for _, v := range t.Data {
		fmt.Fprintf(w, format, v...)
	}
	w.Flush()
	t.Header = nil
	t.Data = nil
}
