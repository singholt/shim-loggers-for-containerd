package fluent_logger_golang

//go:generate msgp
type TestMessage struct {
	Foo  string `msg:"foo" json:"foo,omitempty"`
	Hoge string `msg:"hoge" json:"hoge,omitempty"`
}
