package cube_test

import (
	"reflect"
	"testing"

	"github.com/murat/testing-practices/geometry/cube"
)

func TestNew(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name    string
		args    args
		want    *cube.Cube
		wantErr bool
	}{
		{
			name:    "happy path",
			args:    args{10},
			want:    &cube.Cube{10},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cube.New(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCube_Name(t *testing.T) {
	type fields struct {
		Lenght float64
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
			c := &cube.Cube{
				Lenght: tt.fields.Lenght,
			}
			if got := c.Name(); got != tt.want {
				t.Errorf("Cube.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCube_Volume(t *testing.T) {
	type fields struct {
		Lenght float64
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
			c := &cube.Cube{
				Lenght: tt.fields.Lenght,
			}
			if got := c.Volume(); got != tt.want {
				t.Errorf("Cube.Volume() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCube_Surface(t *testing.T) {
	type fields struct {
		Lenght float64
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
			c := &cube.Cube{
				Lenght: tt.fields.Lenght,
			}
			if got := c.Surface(); got != tt.want {
				t.Errorf("Cube.Surface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCube_Perim(t *testing.T) {
	type fields struct {
		Lenght float64
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
			c := &cube.Cube{
				Lenght: tt.fields.Lenght,
			}
			if got := c.Perim(); got != tt.want {
				t.Errorf("Cube.Perim() = %v, want %v", got, tt.want)
			}
		})
	}
}
