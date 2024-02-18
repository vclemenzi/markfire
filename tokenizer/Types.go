package tokenizer

type Token struct {
	Kind    int
	SubKind int
	Content string
	Line    int
}

type OpenableTokens struct {
	Configuration *Configuration
	List          *List
	Blockquote    *Blockquote
}

type List struct {
	IsOpen  bool
	Index   int
	Subkind int
}

type Blockquote struct {
	IsOpen bool
	Index  int
}

type Configuration struct {
	IsOpen bool
	Index  int
}
