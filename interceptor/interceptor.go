package interceptor

type contextKey int

const (
	//lint:ignore U1000 because it's actually used
	contextKeySegment contextKey = iota
)
