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
	Resp      Message
	BizErr    BizError
	CommonErr CommonError
	Err       error
}

// ResponseHandler is the function to handle http response.
// It may bind result or return an error
type ResponseHandler func(res []byte, bizErr []byte, commonErr []byte) *Response
