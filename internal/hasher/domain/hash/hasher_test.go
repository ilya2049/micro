package hash_test

import (
	"hasher/domain/hash"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSHA3HashSum(t *testing.T) {
	type args struct {
		inputs []hash.Input
	}
	tests := []struct {
		name string
		args args
		want hash.SHA3Hashes
	}{
		{
			name: "No inputs. Expected an empty output collection.",
			args: args{
				inputs: []hash.Input{},
			},
			want: hash.SHA3Hashes{},
		},
		{
			name: "One input. Expected one sha3 hash in the output collection.",
			args: args{
				inputs: []hash.Input{
					hash.NewInput(0, "1"),
				},
			},
			want: hash.SHA3Hashes{
				hash.NewSHA3(0, "67b176705b46206614219f47a05aee7ae6a3edbe850bbbe214c536b989aea4d2"),
			},
		},
		{
			name: "Two inputs. Expected two sha3 hashes in the output collection in arbitrary order.",
			args: args{
				inputs: []hash.Input{
					hash.NewInput(0, "1"),
					hash.NewInput(1, "2"),
				},
			},
			want: hash.SHA3Hashes{
				hash.NewSHA3(0, "67b176705b46206614219f47a05aee7ae6a3edbe850bbbe214c536b989aea4d2"),
				hash.NewSHA3(1, "b1b1bd1ed240b1496c81ccf19ceccf2af6fd24fac10ae42023628abbe2687310"),
			},
		},
		{
			name: "Three inputs. Expected three sha3 hashes in the output collection in arbitrary order.",
			args: args{
				inputs: []hash.Input{
					hash.NewInput(0, "1"),
					hash.NewInput(1, "2"),
					hash.NewInput(2, "3"),
				},
			},
			want: hash.SHA3Hashes{
				hash.NewSHA3(0, "67b176705b46206614219f47a05aee7ae6a3edbe850bbbe214c536b989aea4d2"),
				hash.NewSHA3(1, "b1b1bd1ed240b1496c81ccf19ceccf2af6fd24fac10ae42023628abbe2687310"),
				hash.NewSHA3(2, "1bf0b26eb2090599dd68cbb42c86a674cb07ab7adc103ad3ccdf521bb79056b9"),
			},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.ElementsMatch(t, tt.want, hash.CalculateSHA3HashSum(tt.args.inputs))
		})
	}
}
