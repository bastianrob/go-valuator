package valuator_test

import (
	"testing"

	valuator "github.com/bastianrob/go-valuator"
	"github.com/stretchr/testify/assert"
)

func Test_GTE_Evaluator(t *testing.T) {
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
		args:    args{"number", ">=", "100", "Must be more than 100"},
		obj:     map[string]interface{}{},
		want:    false,
		wantErr: false,
	}, {
		name:    "Integer evaluator",
		args:    args{"number", ">=", "100", "Must be more than 100"},
		obj:     map[string]interface{}{"number": 100},
		want:    true,
		wantErr: false,
	}, {
		name:    "Integer evaluator - negative case",
		args:    args{"number", ">=", "100", "Must be more than 100"},
		obj:     map[string]interface{}{"number": 99},
		want:    false,
		wantErr: false,
	}, {
		name:    "Float evaluator",
		args:    args{"number", ">=", "99.9", "Must be more than 99.9"},
		obj:     map[string]interface{}{"number": 99.9},
		want:    true,
		wantErr: false,
	}, {
		name:    "Float evaluator",
		args:    args{"number", ">=", "99.9", "Must be more than 99.9"},
		obj:     map[string]interface{}{"number": 99.8},
		want:    false,
		wantErr: false,
	}, {
		name:    "String evaluator",
		args:    args{"number", ">=", "99.9", "Must be more than 99.9"},
		obj:     map[string]interface{}{"number": "99.9"},
		want:    true,
		wantErr: false,
	}, {
		name:    "String evaluator",
		args:    args{"number", ">=", "99.9", "Must be more than 99.9"},
		obj:     map[string]interface{}{"number": "99.8"},
		want:    false,
		wantErr: false,
	}, {
		name:    "Not a Number",
		args:    args{"number", ">=", "100", "Must be more than 100"},
		obj:     map[string]interface{}{"number": "I am not a number"},
		want:    false,
		wantErr: false,
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
