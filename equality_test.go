package valuator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	valuator "github.com/bastianrob/go-valuator"
)

func Test_Equality_NoProperty(t *testing.T) {
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
		args:    args{"name", "=", "Name of a property", "Must be equal"},
		obj:     map[string]interface{}{},
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

func Test_Equality_String(t *testing.T) {
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
		name:    "String Equality",
		args:    args{"name", "=", "Name of a property", "Must be equal"},
		obj:     map[string]interface{}{"name": "Name of a property"},
		want:    true,
		wantErr: false,
	}, {
		name:    "String Equality - negative case",
		args:    args{"name", "=", "Name of a property", "Must be equal"},
		obj:     map[string]interface{}{"name": "Name of a property is different from expected"},
		want:    false,
		wantErr: false,
	}, {
		name:    "String Inequality",
		args:    args{"name", "!=", "Name of a property", "Must be equal"},
		obj:     map[string]interface{}{"name": "Name of a property is different from expected"},
		want:    true,
		wantErr: false,
	}, {
		name:    "String Inequality - negative case",
		args:    args{"name", "!=", "Name of a property", "Must be equal"},
		obj:     map[string]interface{}{"name": "Name of a property"},
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

func Test_Equality_Numerical(t *testing.T) {
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
		name:    "Integer Equality",
		args:    args{"number", "=", "50", "Must be equal"},
		obj:     map[string]interface{}{"number": "50"},
		want:    true,
		wantErr: false,
	}, {
		name:    "Integer Equality - negative case",
		args:    args{"number", "=", "50", "Must be equal"},
		obj:     map[string]interface{}{"number": "200"},
		want:    false,
		wantErr: false,
	}, {
		name:    "Float Inequality",
		args:    args{"number", "!=", "50.5", "Must be equal"},
		obj:     map[string]interface{}{"number": "50.55"},
		want:    true,
		wantErr: false,
	}, {
		name:    "Float Inequality - negative case",
		args:    args{"number", "!=", "50.5", "Must be equal"},
		obj:     map[string]interface{}{"number": "50.5"},
		want:    false,
		wantErr: false,
	}, {
		name:    "Object value is not a numeric but must be equal",
		args:    args{"number", "=", "50.5", "Must be equal"},
		obj:     map[string]interface{}{"number": "I am not a number"},
		want:    false,
		wantErr: false,
	}, {
		name:    "Object value is not a numeric but must not be equal",
		args:    args{"number", "!=", "50.5", "Must be equal"},
		obj:     map[string]interface{}{"number": "I am not a number"},
		want:    true,
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
