package reader

type Reader interface {
	next() *string
	peek() *string
}

type TokenReader struct {
	Tokens   []string
	Position int
}

func (tr *TokenReader) Peek() *string {
	if tr.Position > len(tr.Tokens) {
		return nil
	}
	return &tr.Tokens[tr.Position]
}

func (tr *TokenReader) Pop() *string {
	if tr.Position > len(tr.Tokens) {
		return nil
	}
	token := &tr.Tokens[tr.Position]
	tr.Position = tr.Position + 1
	return token
}
func (tr *TokenReader) Next() *string { return tr.next() }
func (tr *TokenReader) next() *string {
	if tr.Position > len(tr.Tokens) {
		return nil
	}
	token := tr.Tokens[tr.Position]
	tr.Position = tr.Position + 1
	return &token
}

func (tr *TokenReader) peek() *string {
	if tr.Position > len(tr.Tokens) {
		return nil
	}
	return &tr.Tokens[tr.Position]
}
