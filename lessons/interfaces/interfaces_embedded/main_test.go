package main

import (
	"testing"
)

func Test_cube_area(t *testing.T) {
	type fields struct {
		edge  float64
		color string
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cube{
				edge:  tt.fields.edge,
				color: tt.fields.color,
			}
			if got := c.area(); got != tt.want {
				t.Errorf("cube.area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cube_volume(t *testing.T) {
	type fields struct {
		edge  float64
		color string
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cube{
				edge:  tt.fields.edge,
				color: tt.fields.color,
			}
			if got := c.volume(); got != tt.want {
				t.Errorf("cube.volume() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cube_getColor(t *testing.T) {
	type fields struct {
		edge  float64
		color string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cube{
				edge:  tt.fields.edge,
				color: tt.fields.color,
			}
			if got := c.getColor(); got != tt.want {
				t.Errorf("cube.getColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_measure(t *testing.T) {
	type args struct {
		g geometry
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := measure(tt.args.g)
			if got != tt.want {
				t.Errorf("measure() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("measure() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
