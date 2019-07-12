package valuator_test

import (
	"testing"

	valuator "github.com/bastianrob/go-valuator"
	"github.com/stretchr/testify/assert"
)

func Test_Intersect_Evaluator(t *testing.T) {
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
		args:    args{"status", "intersect", `["Sold", "Returned", "Void"]`, ""},
		obj:     map[string]interface{}{},
		want:    false,
		wantErr: false,
	}, {
		name:    "Property intersects of",
		args:    args{"status", "intersect", `["Sold", "Returned", "Void"]`, "Status intersects"},
		obj:     map[string]interface{}{"status": []string{"Hold", "Sold"}},
		want:    true,
		wantErr: false,
	}, {
		name:    "Property does not instersects",
		args:    args{"status", "intersect", `["Sold", "Returned", "Void"]`, "Status intersects"},
		obj:     map[string]interface{}{"status": []string{"Not A Status"}},
		want:    false,
		wantErr: false,
	}, {
		name:    "Expression is not valid",
		args:    args{"status", "intersect", `I am not an array`, "Status intersects"},
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
