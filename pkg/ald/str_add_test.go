package ald

import "testing"

func TestStrAdd(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args:    args{"1", "1"},
			want:    "2",
			wantErr: false,
		},
		{
			args:    args{"", ""},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StrAdd(tt.args.str1, tt.args.str2)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrAdd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StrAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
