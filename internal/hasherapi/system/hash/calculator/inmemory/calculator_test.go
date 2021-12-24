package inmemory_test

import (
	"context"
	"hasherapi/domain/hash"
	"hasherapi/system/hash/calculator/inmemory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashCalculator_Calculate(t *testing.T) {
	type args struct {
		inputs []hash.Input
	}

	tests := []struct {
		name string
		args args
		want hash.SHA3Hashes
	}{
		{
			name: "No hashes.",
			args: args{
				inputs: []hash.Input{},
			},
			want: []hash.SHA3Hash{},
		},
		{
			name: "One hash.",
			args: args{
				inputs: []hash.Input{
					hash.Input("1"),
				},
			},
			want: []hash.SHA3Hash{
				hash.SHA3Hash("hash-of-1"),
			},
		},
		{
			name: "Two hashes.",
			args: args{
				inputs: []hash.Input{
					hash.Input("1"),
					hash.Input("2"),
				},
			},
			want: []hash.SHA3Hash{
				hash.SHA3Hash("hash-of-1"),
				hash.SHA3Hash("hash-of-2"),
			},
		},
	}

	calculator := inmemory.NewHashCalculator()

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sha3Hashes, _ := calculator.Calculate(context.Background(), tt.args.inputs)
			assert.Equal(t, tt.want, sha3Hashes)
		})
	}
}
