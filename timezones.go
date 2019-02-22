package go_timezones

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

type Timezone string

type Timezones []Timezone

func (t Timezones) Len() int {
	return len(t)
}

func (t Timezones) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Timezones) Less(i, j int) bool {
	return t[i] < t[j]
}

var timezonePatterns = []string{
	"/usr/share/zoneinfo/*",
	"/usr/share/zoneinfo/*/*",
	"/usr/share/zoneinfo/*/*/*",
}

func GetTimezones() Timezones {
	timezones := make(Timezones, 0, 658)

	for _, pattern := range timezonePatterns {
		paths, err := filepath.Glob(pattern)
		if err != nil {
			panic(fmt.Sprintf("glob failed, pattern is %s", pattern))
		}
		for _, path := range paths {
			fileInfo, err := os.Stat(path)
			if err != nil {
				panic("os.Stat")
			}
			if !fileInfo.IsDir() {
				reg := regexp.MustCompile(`\/([A-Z][^\/]{1,}(\/[A-Z][^\/]{1,})?(\/[A-Z][^\/]{1,})?)$`)
				if match := reg.FindStringSubmatch(path); match != nil {
					timezones = append(timezones, Timezone(match[1]))
				}
			}

		}
	}

	sort.Sort(timezones)

	//if len(timezones) > 0 {
	//	prev := timezones[0]
	//	for i := 1; i < len(timezones); i++ {
	//		if prev == timezones[i] {
	//			timezones = append(timezones[:i], timezones[i+1:]...)
	//		}
	//		prev = timezones[i]
	//	}
	//}

	return timezones
}
