package grpcutil

import (
	"google.golang.org/grpc/metadata"
)

func MetadataGetFirst(key string, m metadata.MD) string {
	metadataValues := m.Get(key)
	if len(metadataValues) > 0 {
		return metadataValues[0]
	}

	return ""
}
