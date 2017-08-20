package trace

// Tracer is an interface コード内でのできごとを記録できるオブジェクトを表す
type Tracer interface {
	Trace(...interface{})
	// ...interface{} は任意の型を何個でも受け取れることを意味する
	// fmt.Sprintやlog.Fatalなどもこのような書き方をしていてGoコミュニティーにとっては馴染みの深いパターン
}
