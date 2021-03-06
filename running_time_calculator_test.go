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
			name: "5KM in 25mins",
			args: args{
				kilometers:  5,
				timeMinutes: 25,
				timeSeconds: 0,
			},
			want: "5:00 /km",
		},
		{
			name: "5KM in 25m 25s",
			args: args{
				kilometers:  5,
				timeMinutes: 25,
				timeSeconds: 25,
			},
			want: "5:05 /km",
		},
		{
			name: "10KM in 50m",
			args: args{
				kilometers:  10,
				timeMinutes: 50,
				timeSeconds: 0,
			},
			want: "5:00 /km",
		},
		{
			name: "5KM in 30m",
			args: args{
				kilometers:  5,
				timeMinutes: 30,
				timeSeconds: 0,
			},
			want: "6:00 /km",
		},
		{
			name: "5KM in 29m",
			args: args{
				kilometers:  5,
				timeMinutes: 29,
				timeSeconds: 0,
			},
			want: "5:48 /km",
		},
		{
			name: "5KM in 24m",
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
			name: "5KM at 5:00 /km",
			args: args{
				kilometers: 5,
				pace:       "5:00 /km",
			},
			want: []string{"5:00", "10:00", "15:00", "20:00", "25:00"},
		},
		{
			name: "5KM at 4:30 /km",
			args: args{
				kilometers: 5,
				pace:       "4:30 /km",
			},
			want: []string{"4:30", "9:00", "13:30", "18:00", "22:30"},
		},
		{
			name: "10KM at 5:00 /km",
			args: args{
				kilometers: 10,
				pace:       "5:00 /km",
			},
			want: []string{"5:00", "10:00", "15:00", "20:00", "25:00", "30:00", "35:00", "40:00", "45:00", "50:00"},
		},
		{
			name: "10KM at 5:59 /km",
			args: args{
				kilometers: 10,
				pace:       "5:59 /km",
			},
			want: []string{"5:59", "11:58", "17:57", "23:56", "29:55", "35:54", "41:53", "47:52", "53:51", "59:50"},
		},
		{
			name: "1KM at 5:00 /km",
			args: args{
				kilometers: 1,
				pace:       "5:00 /km",
			},
			want: []string{"5:00"},
		},
		{
			name: "0KM at 5:00 /km",
			args: args{
				kilometers: 0,
				pace:       "5:00 /km",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateKilometerTimes(tt.args.kilometers, tt.args.pace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateKilometerTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertStringToPaceStruct(t *testing.T) {
	type args struct {
		pace string
	}
	tests := []struct {
		name string
		args args
		want Pace
	}{
		{
			args: args{
				pace: "5:00",
			},
			want: Pace{5, 0},
		},
		{
			args: args{
				pace: "5:24",
			},
			want: Pace{5, 24},
		},
		{
			args: args{
				pace: "4:00",
			},
			want: Pace{4, 0},
		},
		{
			args: args{
				pace: "10:10",
			},
			want: Pace{10, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertStringToPaceStruct(tt.args.pace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertStringToPaceStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
