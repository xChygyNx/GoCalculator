package change

import (
	"errors"
	"fmt"
	"github.com/xChygyNx/change/scan_input"
	"os"
	"sort"
)

func Run() {
	if len(os.Args) != 3 {
		Usage()
		fmt.Println(len(os.Args))
		return
	}
	coins, errC := scan_input.ScanCoins(os.Args[1])
	target, errT := scan_input.ScanTarget(os.Args[2])
	if errC != nil || errT != nil {
		Usage()
		fmt.Println()
		if errC != nil {
			fmt.Println(errC)
		}
		if errT != nil {
			fmt.Println(errT)
		}
		return
	}
	change, err := Change(coins, target)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Change: %v\n", change)
}

func Change(coins []int, target int) ([]int, error) {
	mapCoins := createMap(coins, target)
	if len(mapCoins[target]) == 0 && target != 0 {
		return []int{}, errors.New("Solution not exist")
	}
	return mapCoins[target], nil
}

func createMap(coins []int, target int) map[int][]int {
	tabCount := make([]int, target+1)
	mapNominal := make(map[int][]int, target+1)
	countCoins := len(coins)
	tabCount[0] = 0
	mapNominal[0] = []int{}
	for i := 1; i <= target; i++ {
		tabCount[i] = target + 1
		for j := 0; j < countCoins && coins[j] <= i; j++ {
			k := i - coins[j]
			if k == 0 || (k >= 0 && tabCount[i] > tabCount[k]+1) {
				tabCount[i] = tabCount[k] + 1
				elem := make([]int, len(mapNominal[k]))
				copy(elem, mapNominal[k])
				mapNominal[i] = append(elem, coins[j])
				sort.Ints(mapNominal[i])
			}
		}
	}
	return mapNominal
}

func Usage() {
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("\tgo run main.go \"coins_nominals\" target")
	fmt.Println()
	fmt.Println("coins_nominals\t- list of coins nominals (integers separated by spaces)")
	fmt.Println("target\t\t- sum which need to collect by coins of this nominals")
}
