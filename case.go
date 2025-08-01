package got

type Namer interface {
	Name() string
}

// Case is a interface that defines the structure of a test case.
type Case interface {
	Namer

	Input() any    // the input of the test case
	Want() any     // the expected result of the test case
	WantErr() bool // whether the test case should return an error
	Err() error    // the error of the test case
}

type caseImpl struct {
	name    string
	input   any
	want    any
	wantErr bool
	err     error
}

func (c *caseImpl) Name() string {
	return c.name
}

func (c *caseImpl) Input() any {
	return c.input
}

func (c *caseImpl) Want() any {
	return c.want
}

func (c *caseImpl) WantErr() bool {
	return c.wantErr
}

func (c *caseImpl) Err() error {
	return c.err
}

func NewCase(name string, input any, want any, wantErr bool, err error) Case {
	return &caseImpl{name: name, input: input, want: want, wantErr: wantErr, err: err}
}

func CaseBuilder(name string) *caseBuilder {
	return &caseBuilder{caseImpl: caseImpl{name: name}}
}

type caseBuilder struct {
	caseImpl
}

func (b *caseBuilder) Name(name string) *caseBuilder {
	b.name = name
	return b
}

func (b *caseBuilder) Input(input any) *caseBuilder {
	b.input = input
	return b
}

func (b *caseBuilder) Want(want any) *caseBuilder {
	b.want = want
	return b
}

func (b *caseBuilder) WantErr(wantErr bool) *caseBuilder {
	b.wantErr = wantErr
	return b
}

func (b *caseBuilder) Err(err error) *caseBuilder {
	b.err = err
	return b
}

func (b *caseBuilder) Build() Case {
	return &b.caseImpl
}
