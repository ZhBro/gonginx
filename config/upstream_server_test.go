package config

import (
	"reflect"
	"testing"
)

func TestNewUpstreamServer(t *testing.T) {
	type args struct {
		directive *Directive
	}
	tests := []struct {
		name       string
		args       args
		want       *UpstreamServer
		wantString string
	}{
		{
			name: "new upstream server",
			args: args{
				directive: &Directive{
					Name:       "server",
					Parameters: []string{"127.0.0.1:8080"},
				},
			},
			want: &UpstreamServer{
				Directive: &Directive{
					Name:       "server",
					Parameters: []string{"127.0.0.1:8080"},
				},
				Address:    "127.0.0.1:8080",
				Flags:      make([]string, 0),
				Parameters: make(map[string]string, 0),
			},
			wantString: "server 127.0.0.1:8080;",
		},
		{
			name: "new upstream server with weight",
			args: args{
				directive: &Directive{
					Name:       "server",
					Parameters: []string{"127.0.0.1:8080", "weight=5"},
				},
			},
			want: &UpstreamServer{
				Directive: &Directive{
					Name:       "server",
					Parameters: []string{"127.0.0.1:8080", "weight=5"},
				},
				Address: "127.0.0.1:8080",
				Flags:   make([]string, 0),
				Parameters: map[string]string{
					"weight": "5",
				},
			},
			wantString: "server 127.0.0.1:8080 weight=5;",
		},
		{
			name: "new upstream server with weight and a flag",
			args: args{
				directive: &Directive{
					Name:       "server",
					Parameters: []string{"127.0.0.1:8080", "weight=5", "down"},
				},
			},
			want: &UpstreamServer{
				Directive: &Directive{
					Name:       "server",
					Parameters: []string{"127.0.0.1:8080", "weight=5", "down"},
				},
				Address: "127.0.0.1:8080",
				Flags:   []string{"down"},
				Parameters: map[string]string{
					"weight": "5",
				},
			},
			wantString: "server 127.0.0.1:8080 weight=5 down;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpstreamServer(tt.args.directive)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUpstreamServer() = %v, want %v", got, tt.want)
			}
			if got.ToString() != tt.wantString {
				t.Errorf("NewUpstreamServer().ToString = %v, want %v", got.ToString(), tt.wantString)
			}
		})
	}
}
