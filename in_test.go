package valuator_test

import (
	"testing"

	valuator "github.com/bastianrob/go-valuator"
	"github.com/stretchr/testify/assert"
)

func Test_IN_Evaluator(t *testing.T) {
	type args struct {
		prop string
		op   string
		val  string
		desc string
	}
	tests := []struct {
		name    string
		args    args
		obj     map[string]interface{}
		want    bool
		wantErr bool
	}{{
		name:    "Property does not exists",
		args:    args{"status", "in", `["Sold", "Returned", "Void"]`, "Status is one of"},
		obj:     map[string]interface{}{},
		want:    false,
		wantErr: false,
	}, {
		name:    "Property is one of",
		args:    args{"status", "in", `["Sold", "Returned", "Void"]`, "Status is one of"},
		obj:     map[string]interface{}{"status": "Sold"},
		want:    true,
		wantErr: false,
	}, {
		name:    "Property is not one of",
		args:    args{"status", "in", `["Sold", "Returned", "Void"]`, "Status is one of"},
		obj:     map[string]interface{}{"status": "Not A Status"},
		want:    false,
		wantErr: false,
	}, {
		name:    "Expression is not valid",
		args:    args{"status", "in", `I am not an array`, "Status is one of"},
		obj:     map[string]interface{}{"status": "Not"},
		want:    false,
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluator, err := valuator.NewValuator(tt.args.prop, tt.args.op, tt.args.val, tt.args.desc)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, evaluator.Evaluate(tt.obj))
			}
		})
	}
}
