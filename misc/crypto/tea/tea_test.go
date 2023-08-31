package tea

import "testing"

func TestPadKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{""},
			want: "_\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadKey(tt.args.key); got != tt.want {
				t.Errorf("PadKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
