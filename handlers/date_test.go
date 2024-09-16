package api

import (
	"reflect"
	"testing"
)

func TestReadDate(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    DateEntry
		wantErr bool
	}{
		{
			name:    "Successful case",
			args:    args{id: "1"},
			want:    DateEntry{ID: 1, Dates: []string{"*23-08-2019", "*22-08-2019", "*20-08-2019", "*26-01-2020", "*28-01-2020", "*30-01-2019", "*07-02-2020", "*10-02-2020"}},
			wantErr: false,
		},
		{
			name:    "API returns error status",
			args:    args{id: "404"},
			want:    DateEntry{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDate(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
