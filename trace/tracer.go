package trace

import (
	"fmt"
	"io"
)

// Tracerはコード内での出来事を記録できるオブジェクトを表すインターフェースです。
// Traceという1つのメソッドを持ちます。
type Tracer interface {
	// Trace()は任意の型の引数をいくらでも受け取ります。GoのIFあるある。
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// OffはTraceメソッドの呼び出しを無視するTracerを返します。
func Off() Tracer {
	return &nilTracer{}
}
