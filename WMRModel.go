package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	var width, length, height float64
	var WMR_CaRoomt0, WMR_CaRoomCaIn, WMR_CaRoomGa, WMR_CaRoomQft3, WMR_proct, WMR_shiftt, WMR_OEL float64
	fmt.Print("Enter width: (ft)")
	fmt.Scan(&width)
	fmt.Print("Enter length: (ft)")
	fmt.Scan(&length)
	fmt.Print("Enter height: (ft)")
	fmt.Scan(&height)
	fmt.Print("Enter the initial room concentration (mg/m3): ")
	fmt.Scan(&WMR_CaRoomt0)
	fmt.Print("Enter the concentration of contaminant in the supply air (mg/m3): ")
	fmt.Scan(&WMR_CaRoomCaIn)
	fmt.Print("Enter the contaminant generation rate: (mg/min)")
	fmt.Scan(&WMR_CaRoomGa)
	fmt.Print("Enter the ventilation rate (ft3/min): ")

	fmt.Scan(&WMR_CaRoomQft3)
	fmt.Print("Enter the process duration (min): ")
	fmt.Scan(&WMR_proct)
	fmt.Print("Enter the shift duration (min): ")
	fmt.Scan(&WMR_shiftt)
	fmt.Print("Enter the occupational exposure limit (mg/m3): ")
	fmt.Scan(&WMR_OEL)
	dimensionsft3 := width * length * height
	dimensionsm3 := dimensionsft3 * 0.0283168
	WMR_CaRoomQm3 := WMR_CaRoomQft3 * 0.0283168
	WMR_SteadyState_ProcT := (WMR_CaRoomt0-WMR_CaRoomCaIn-(WMR_CaRoomGa/WMR_CaRoomQm3))*(math.Exp((-WMR_CaRoomQm3*WMR_proct)/dimensionsm3)) + WMR_CaRoomCaIn + (WMR_CaRoomGa / WMR_CaRoomQm3)
	TWA := (WMR_SteadyState_ProcT * WMR_proct) / WMR_shiftt

	fmt.Printf("The steady-state concentration at the end of the process is: %.2f mg/m3\n", WMR_SteadyState_ProcT)
	if WMR_SteadyState_ProcT > WMR_OEL {
		fmt.Println("The concentration exceeds the occupational exposure limit.")
	} else {
		fmt.Println("The concentration is within the occupational exposure limit.")
	}
	fmt.Printf("The time-weighted average concentration over the shift is: %.2f mg/m3\n", TWA)
	if TWA > WMR_OEL {
		fmt.Println("The time-weighted average concentration exceeds the occupational exposure limit.")
	} else {
		fmt.Println("The time-weighted average concentration is within the occupational exposure limit.")
	}
	fmt.Printf("Execution Time: %s\n", time.Since(start))
}
