package hash_test

import (
	"hasher/domain/hash"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInputs(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want []hash.Input
	}{
		{
			name: "Inputs ordered correctly.",
			args: args{
				strings: []string{"input-1", "input-2", "input-3"},
			},
			want: []hash.Input{
				hash.NewInput(0, "input-1"),
				hash.NewInput(1, "input-2"),
				hash.NewInput(2, "input-3"),
			},
		},
		{
			name: "Single input.",
			args: args{
				strings: []string{"input-1"},
			},
			want: []hash.Input{
				hash.NewInput(0, "input-1"),
			},
		},
		{
			name: "No inputs.",
			args: args{
				strings: []string{},
			},
			want: []hash.Input{},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, hash.NewInputs(tt.args.strings))
		})
	}
}

func TestSHA3Hashes_ToStrings(t *testing.T) {
	tests := []struct {
		name    string
		outputs hash.SHA3Hashes
		want    []string
	}{
		{
			name: "Ordered items.",
			outputs: []hash.SHA3{
				hash.NewSHA3(0, "hash-1"),
				hash.NewSHA3(1, "hash-2"),
				hash.NewSHA3(2, "hash-3"),
			},
			want: []string{"hash-1", "hash-2", "hash-3"},
		},
		{
			name: "Unordered items.",
			outputs: []hash.SHA3{
				hash.NewSHA3(1, "hash-2"),
				hash.NewSHA3(2, "hash-3"),
				hash.NewSHA3(0, "hash-1"),
			},
			want: []string{"hash-1", "hash-2", "hash-3"},
		},
		{
			name: "One item.",
			outputs: []hash.SHA3{
				hash.NewSHA3(0, "hash-1"),
			},
			want: []string{"hash-1"},
		},
		{
			name:    "No items.",
			outputs: []hash.SHA3{},
			want:    []string{},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.outputs.ToStrings())
		})
	}
}
