package utils

import "testing"

func TestIsZeroValue(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-int",
			args: args{
				i: 0,
			},
			want: true,
		},
		{
			name: "test-string",
			args: args{
				i: "",
			},
			want: true,
		},
		{
			name: "test-string-false",
			args: args{
				i: "test",
			},
			want: false,
		},
		{
			name: "test-struct",
			args: args{
				i: struct{}{},
			},
			want: true,
		},
		{
			name: "test-nil",
			args: args{
				i: nil,
			},
			want: true,
		},
		{
			name: "test-bool",
			args: args{
				i: false,
			},
			want: true,
		},
		{
			name: "test-float",
			args: args{
				i: 0.0,
			},
			want: true,
		},
		{
			name: "test-custom-true",
			args: args{
				i: args{},
			},
			want: true,
		},

		{
			name: "test-custom-false",
			args: args{
				i: args{i: 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZeroValue(tt.args.i); got != tt.want {
				t.Errorf("IsZeroValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
