package main

import "fmt"

func Attendance(employeeTime string) string {

	hour := int(employeeTime[0]-'0')*10 + int(employeeTime[1]-'0')
	minute := int(employeeTime[3]-'0')*10 + int(employeeTime[4]-'0')

	if hour < 8 || (hour == 8 && minute == 0) {
		return "On Time in " + employeeTime
	}

	return "Late in " + employeeTime

}

func main() {
	attendanceTime := "10:59"
	attendanceTime2 := "06:01"

	fmt.Println(Attendance(attendanceTime))
	fmt.Println(Attendance(attendanceTime2))
}
