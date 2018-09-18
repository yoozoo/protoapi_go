package protoapigo

const (
	HAPPYPATH   = 200
	BIZERROR    = 400
	COMMONERROR = 420
	ERROR       = 500
)

// Message protoconf configuration object interface
type Message interface {
}

type Response struct {
	Resp      *Message
	BizErr    BizError
	CommonErr CommonError
	Err       error
}
