package fakecalculator_test

import (
	"context"
	"hasherapi/internal/domain/hash"
	"hasherapi/internal/system/hash/fakecalculator"
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

	calculator := fakecalculator.New()

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			sha3Hashes, _ := calculator.Calculate(context.Background(), tt.args.inputs)
			assert.Equal(t, tt.want, sha3Hashes)
		})
	}
}
