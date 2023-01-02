package storage

import (
	"reflect"
	"testing"
)

func TestNewInMemoryDB(t *testing.T) {
	tests := map[string]struct {
		want InMemoryDB
	}{
		"Empty-DB": {
			want: nil,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := NewInMemoryDB(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: NewInMemoryDB() = %v, want %v", got, tc.want)
			}
		})
	}
}
