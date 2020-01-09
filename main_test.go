package microcli

import (
	"reflect"
	"testing"
)

func TestCall(t *testing.T) {
	type args struct {
		ns       string
		svcName  string
		endPoint string
		body     string
		result   interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ns:       "micro",
				svcName:  "authen",
				endPoint: "login",
				body:     "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Call(tt.args.ns, tt.args.svcName, tt.args.endPoint, tt.args.body, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("Call() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStrToMap(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToMap(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
