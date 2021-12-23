package hash

import "context"

type Calculator interface {
	Calculate(context.Context, []Input) (SHA3Hashes, error)
}
