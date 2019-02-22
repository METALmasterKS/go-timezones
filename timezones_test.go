package go_timezones

import (
	"testing"
)

func TestTimezones(t *testing.T) {
	tests := []struct {
		name string
		want Timezone
	}{
		{"UTC", Timezone("UTC")},
		{"Europe/Moscow", Timezone("Europe/Moscow")},
		{"America/North_Dakota/Beulah", Timezone("America/North_Dakota/Beulah")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTimezones()
			if !foundedTimezone(tt.want, got) {
				t.Errorf("Timezones() = %v, want %v", got, tt.want)
			}
		})
	}
}

func foundedTimezone(found Timezone, sl Timezones) bool {
	for _, v := range sl {
		if v == found {
			return true
		}
	}
	return false
}
