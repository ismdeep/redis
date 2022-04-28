package redis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_client_SaveObject(t *testing.T) {
	type args struct {
		key   string
		value interface{}
		opts  []Option
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				key:   "",
				value: nil,
				opts:  nil,
			},
			wantErr: true,
		},
		{
			name: "",
			args: args{
				key:   "test-name",
				value: nil,
				opts:  nil,
			},
			wantErr: true,
		},
		{
			name: "",
			args: args{
				key:   "",
				value: "test-content",
				opts:  nil,
			},
			wantErr: true,
		},
		{
			name: "",
			args: args{
				key:   "test-key",
				value: "test-value",
				opts:  nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantErr, testClient.SaveObject(tt.args.key, tt.args.value, tt.args.opts...) != nil)
		})
	}
}

func TestClient_GetObject(t *testing.T) {
	assert.NoError(t, testClient.SaveObject("test-key-001", "value"))

	var s string
	assert.NoError(t, testClient.GetObject("test-key-001", &s))
	assert.Equal(t, "value", s)
}
