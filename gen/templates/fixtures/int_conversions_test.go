package fixtures

import (
	"math/rand"
	"testing"
)

func TestIsNumber(t *testing.T) {
	type args struct {
		in int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1 is NumberOne",
			args{
				1,
			},
			true,
		},
		{
			"2 is NumberTwo",
			args{
				2,
			},
			true,
		},
		{
			"3 is NumberThree",
			args{
				3,
			},
			true,
		},
		{
			"errors for a random number is not a Number",
			args{
				rand.Int(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumber(tt.args.in); got != tt.want {
				t.Errorf("IsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToNumber(t *testing.T) {
	type args struct {
		in int
	}
	tests := []struct {
		name  string
		args  args
		want  Number
		want1 bool
	}{
		{
			"1 is NumberOne",
			args{
				1,
			},
			NumberOne,
			true,
		},
		{
			"2 is NumberTwo",
			args{
				2,
			},
			NumberTwo,
			true,
		},
		{
			"3 is NumberThree",
			args{
				3,
			},
			NumberThree,
			true,
		},
		{
			"errors for a random number is not a Number",
			args{
				rand.Int(),
			},
			Number(0),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToNumber(tt.args.in)
			if got != tt.want {
				t.Errorf("ToNumber() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToNumber() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToNumberErr(t *testing.T) {
	type args struct {
		in int
	}
	tests := []struct {
		name    string
		args    args
		want    Number
		wantErr bool
	}{
		{
			"1 is NumberOne",
			args{
				1,
			},
			NumberOne,
			false,
		},
		{
			"2 is NumberTwo",
			args{
				2,
			},
			NumberTwo,
			false,
		},
		{
			"3 is NumberThree",
			args{
				3,
			},
			NumberThree,
			false,
		},
		{
			"errors for a random number is not a Number",
			args{
				rand.Int(),
			},
			Number(0),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToNumberErr(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToNumberErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToNumberErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustToNumber(t *testing.T) {
	type args struct {
		in int
	}
	tests := []struct {
		name      string
		args      args
		want      Number
		wantPanic bool
	}{
		{
			"1 is NumberOne",
			args{
				1,
			},
			NumberOne,
			false,
		},
		{
			"2 is NumberTwo",
			args{
				2,
			},
			NumberTwo,
			false,
		},
		{
			"3 is NumberThree",
			args{
				3,
			},
			NumberThree,
			false,
		},
		{
			"panics for a random number is not a Number",
			args{
				rand.Int(),
			},
			Number(0),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				defer func() { recover() }()
			}

			if got := MustToNumber(tt.args.in); got != tt.want {
				t.Errorf("MustToNumber() = %v, want %v", got, tt.want)
			}

			if tt.wantPanic {
				t.Errorf("did not panic")
			}
		})
	}
}
