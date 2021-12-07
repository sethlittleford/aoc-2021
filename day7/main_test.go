package main

import "testing"

func Test_median(t *testing.T) {
	tests := []struct {
		name    string
		dataSet []float64
		want    float64
	}{
		{
			name: "Success on odd set",
			dataSet: []float64{1, 2, 2, 3, 4, 7, 9},
			want: 3,
		},
		{
			name: "Success on even set",
			dataSet: []float64{16,1,2,0,4,2,7,1,2,14},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := median(tt.dataSet); got != tt.want {
				t.Errorf("median() = %v, want %v", got, tt.want)
			}
		})
	}
}
