package trace

import (
	"bytes"
	"testing"
)

// Goにおいてテストは大切
// テスト対象と同じフォルダに置かれ、名前はTestで始まり_test.goで終わる。
// *testing.T型を受け取る。
func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Newからの戻り値がnilです。")
	} else {
		tracer.Trace("こんにちは、traceパッケージ")
		if buf.String() != "こんにちは、traceパッケージ\n" {
			t.Errorf("'%s'という誤った文字列が出力されました", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("データ")
}
