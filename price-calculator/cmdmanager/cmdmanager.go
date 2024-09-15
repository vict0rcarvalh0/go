package cmdmanager

import "fmt"

type CmdManager struct {

}

func (cmd CmdManager) ReadLinesFromFile() ([]string, error) {
	fmt.Println("Prease enter yout prices. Confirm every price with ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}


func (cmd CmdManager) WriteJson(data interface{}) error {
	fmt.Println(data)
	return nil
}

func NewCmdManager() CmdManager{
	return CmdManager{}
} 