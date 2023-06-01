package main


// This solves the problem mentioned in the   
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	data := `Dubai-1, 11014, SP-1 , 17803, SX, 4923, HV, 2069, P 
Mumbai-3, 22547, SP-1, 9389, SP-2, 4829, SX, 3375, HV, 3371, P, 309, Mu`

	z := strings.Split(data, "\n")

	panelNames := map[string]string{
		"SP-1": "Standard Panel 1",
		"SP-2": "Standard Panel 2",
		"SX":   "Solar Extension",
		"HV":   "High Voltage",
		"P":    "Premium",
		"Mu":   "Multi-Panel",
	}

	for _, line := range z {
		words := strings.Split(line, ",")
		dealershipName := words[0]
		fmt.Println("Dealership Name:", dealershipName)

		panelPairs := make([]PanelPair, 0)
		totalSales := 0

		for i := 1; i < len(words); i += 2 {
			num, err := strconv.Atoi(strings.TrimSpace(words[i]))
			if err != nil {
				// Handle the error
				continue
			}
			panelCode := strings.TrimSpace(words[i+1])
			panelName := panelNames[panelCode]

			panelPairs = append(panelPairs, PanelPair{PanelsSold: num, PanelCode: panelCode, PanelName: panelName})
			totalSales += num
		}

		fmt.Println("Panel Pairs:", panelPairs)
		fmt.Println("Total Sales:", totalSales)

		// Calculate and display the percentages
		for _, pair := range panelPairs {
			percentage := float64(pair.PanelsSold) / float64(totalSales) * 100
			fmt.Printf("%s: %.2f%%\n", pair.PanelName, percentage)
		}

		fmt.Println()
	}
}

type PanelPair struct {
	PanelsSold int
	PanelCode  string
	PanelName  string
}
