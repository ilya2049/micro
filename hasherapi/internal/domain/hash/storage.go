package hash

import "context"

type Storage interface {
	Save(context.Context, SHA3Hashes) ([]IdentifiedSHA3Hash, error)
	Get(context.Context, []ID) ([]IdentifiedSHA3Hash, error)
}
