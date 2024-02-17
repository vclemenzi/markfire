package tokenizer

type Token struct {
	Kind    int
	SubKind int
	Content string
	Line    int
}

type OpenableTokens struct {
	List       *List
	Blockquote *Blockquote
}

type List struct {
	IsOpen  bool
	Index   int
	Subkind int
	Closure int
}

type Blockquote struct {
	IsOpen  bool
	Index   int
	Closure int
}
