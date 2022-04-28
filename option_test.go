package redis

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_getOption(t *testing.T) {
	d := 100 * time.Second
	s := EncoderJSON
	s2 := EncoderGOB

	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		{
			name: "",
			args: args{
				opts: []Option{WithExpireTime(100 * time.Second), withEncoder(EncoderJSON)},
			},
			want: Option{
				ExpireDuration: &d,
				Encoder:        &s,
			},
		},
		{
			name: "",
			args: args{
				opts: []Option{WithJSONEncoder()},
			},
			want: Option{
				ExpireDuration: nil,
				Encoder:        &s,
			},
		},
		{
			name: "",
			args: args{
				opts: []Option{WithGOBEncoder()},
			},
			want: Option{
				ExpireDuration: nil,
				Encoder:        &s2,
			},
		},
		{
			name: "",
			args: args{
				opts: []Option{WithDefaultEncoder()},
			},
			want: Option{
				ExpireDuration: nil,
				Encoder:        &s2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getOption(tt.args.opts...)
			assert.Equal(t, tt.want, got)
		})
	}
}
