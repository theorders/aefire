package aefire

import (
	"fmt"
	"time"
)

func TimeNowRef() *time.Time {
	now := time.Now()
	return &now
}

func Timezone(hours, minutes int) *time.Location {
	return time.FixedZone(fmt.Sprintf("UTC%d:%02d", hours, minutes), hours*60*60+minutes*60)
}

func KRTimezone() *time.Location {
	return Timezone(9, 0)
}
