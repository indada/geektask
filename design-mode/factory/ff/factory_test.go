package ff

import (
	"reflect"
	"testing"
)

func TestNewIRuleConfigParser(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParser
	}{
		{
			name: "jsons",
			args: args{t: "jsons"},
			want: jsonRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIRuleConfigParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}