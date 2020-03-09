package main

const (
	SecondsPerMinute = 60
	SecondPerHour    = SecondsPerMinute * 60
	SecondPerDay     = SecondPerHour * 24
)

func resolveTime(seconds int) (day, hour, minute int) {
	day = seconds / SecondPerDay
	hour = seconds / SecondPerHour
	minute = seconds / SecondsPerMinute
	return
}

func main() {

}
