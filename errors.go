package protoapigo

// BizError is for business error
type BizError interface {
	error
}

// CommonError is for common error
type CommonError interface {
	error
}
