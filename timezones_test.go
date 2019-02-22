package go_timezones

import (
	"reflect"
	"testing"
)

func TestTimezones(t *testing.T) {
	tests := []struct {
		name string
		want Timezone
	}{
		{"Europe/Moscow", Timezone("Europe/Moscow")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTimezones(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Timezones() = %v, want %v", got, tt.want)
			}
		})
	}
}
