package objects

import (
	"testing"
)

var cr = map[string]map[string]float64{
	"TWD": {"TWD": 1, "JPT": 3.669, "USD": 0.03281},
	"JPY": {"TWD": 0.26956, "JPY": 1, "USD": 0.00885},
	"USD": {"TWD": 30.444, "JPY": 111.801, "USD": 1},
}

func TestCurrencies_Convert(t *testing.T) {
	type fields struct {
		Currencies map[string]map[string]float64
	}
	fs := fields{cr}
	type args struct {
		amount string
		from   string
		to     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "testcase1",
			fields: fs,
			args: args{
				amount: "12345",
				from:   "TWD",
				to:     "USD",
			},
			want:    "405.04",
			wantErr: false,
		},
		{
			name:   "testcase2",
			fields: fs,
			args: args{
				amount: "12345.6789",
				from:   "TWD",
				to:     "USD",
			},
			want:    "405.06",
			wantErr: false,
		},
		{
			name:   "testcase3",
			fields: fs,
			args: args{
				amount: "0",
				from:   "TWD",
				to:     "USD",
			},
			want:    "0.00",
			wantErr: false,
		},
		{
			name:   "testcase4",
			fields: fs,
			args: args{
				amount: "123456.789",
				from:   "JPY",
				to:     "TWD",
			},
			want:    "33,279.01",
			wantErr: false,
		},
		{
			name:   "testcase5",
			fields: fs,
			args: args{
				amount: "-9992414.24",
				from:   "USD",
				to:     "TWD",
			},
			want:    "-304,209,059.12",
			wantErr: false,
		},
		{
			name:   "testcase6",
			fields: fs,
			args: args{
				amount: "9999999.99999",
				from:   "USD",
				to:     "JPY",
			},
			want:    "1,118,010,000.00",
			wantErr: false,
		},
		{
			name:   "testcase7",
			fields: fs,
			args: args{
				amount: "-1",
				from:   "JPY",
				to:     "USD",
			},
			want:    "-0.01",
			wantErr: false,
		},
		{
			name:   "testcase8",
			fields: fs,
			args: args{
				amount: "-.01",
				from:   "JPY",
				to:     "USD",
			},
			want:    "-0.00",
			wantErr: false,
		},
		{
			name:   "testcase_no_amount_args",
			fields: fs,
			args: args{
				from: "TWD",
				to:   "USD",
			},
			want:    "",
			wantErr: true,
		},
		{
			name:   "testcase_invalid_amount_1",
			fields: fs,
			args: args{
				amount: "123.456.3",
				from:   "TWD",
				to:     "USD",
			},
			want:    "",
			wantErr: true,
		},
		{
			name:   "testcase_invalid_amount_2",
			fields: fs,
			args: args{
				amount: "00e123456.3",
				from:   "TWD",
				to:     "USD",
			},
			want:    "",
			wantErr: true,
		},
		{
			name:   "testcase_no_from_args",
			fields: fs,
			args: args{
				amount: "123456.789",
				to:     "TWD",
			},
			want:    "",
			wantErr: true,
		},
		{
			name:   "testcase_invalid_from",
			fields: fs,
			args: args{
				amount: "123456.789",
				from:   "EUR",
				to:     "TWD",
			},
			want:    "",
			wantErr: true,
		},
		{
			name:   "testcase_no_to_args",
			fields: fs,
			args: args{
				amount: "123456.789",
				from:   "JPY`",
			},
			want:    "",
			wantErr: true,
		},
		{
			name:   "testcase_invalid_to",
			fields: fs,
			args: args{
				amount: "123456.789",
				from:   "JPY",
				to:     "CAD",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Currencies{
				Currencies: tt.fields.Currencies,
			}
			got, err := c.Convert(tt.args.amount, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Convert() got = %v, want %v", got, tt.want)
			}
		})
	}
}
