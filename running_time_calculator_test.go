package main

import (
	"reflect"
	"testing"
)

func Test_calculateAveragePace(t *testing.T) {
	type args struct {
		kilometers  float64
		timeMinutes float64
		timeSeconds float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				kilometers:  5,
				timeMinutes: 25,
				timeSeconds: 0,
			},
			want: "5:00 /km",
		},
		{
			args: args{
				kilometers:  5,
				timeMinutes: 25,
				timeSeconds: 25,
			},
			want: "5:05 /km",
		},
		{
			args: args{
				kilometers:  10,
				timeMinutes: 50,
				timeSeconds: 0,
			},
			want: "5:00 /km",
		},
		{
			args: args{
				kilometers:  5,
				timeMinutes: 30,
				timeSeconds: 0,
			},
			want: "6:00 /km",
		},
		{
			args: args{
				kilometers:  5,
				timeMinutes: 29,
				timeSeconds: 0,
			},
			want: "5:48 /km",
		},
		{
			args: args{
				kilometers:  5,
				timeMinutes: 24,
				timeSeconds: 0,
			},
			want: "4:48 /km",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateAveragePace(tt.args.kilometers, tt.args.timeMinutes, tt.args.timeSeconds); got != tt.want {
				t.Errorf("calculateAveragePace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateKilometerTimes(t *testing.T) {
	type args struct {
		kilometers int
		pace       string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				kilometers: 5,
				pace:       "5:00 /km",
			},
			want: []string{"5:00", "10:00", "15:00", "20:00", "25:00"},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateKilometerTimes(tt.args.kilometers, tt.args.pace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateKilometerTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
