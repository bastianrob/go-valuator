package valuator_test

import (
	"testing"

	valuator "github.com/bastianrob/go-valuator"
	"github.com/stretchr/testify/assert"
)

func Test_Invalid_Evaluator(t *testing.T) {
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
		name:    "Evaluator does not exists",
		args:    args{"eval", "something-something", "No", ""},
		obj:     map[string]interface{}{},
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
