package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string
	Data InitialData
}

type InitialData struct {
	InitialAmount float64
	SplitAmount   float64
	// Balance float
}

type Settlement struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

// type Queue struct {
// 	Position int
// 	Amount   float64
// }

var StackPos []int
var QueuePos []int
var QueueAmount []float64
var StackAmount []float64

func calculator(c *gin.Context) {

	data := []Person{
		{
			Name: "deb",
			Data: InitialData{
				InitialAmount: 3500,
				SplitAmount:   500,
			},
		},
		{
			Name: "saheb",
			Data: InitialData{
				InitialAmount: 1500,
				SplitAmount:   500,
			},
		},
		{
			Name: "raju",
			Data: InitialData{
				InitialAmount: 0,
				SplitAmount:   2000,
			},
		},
		{
			Name: "minu",
			Data: InitialData{
				InitialAmount: 0,
				SplitAmount:   2000,
			},
		},
	}
	var settled []Settlement
	// var stackAmount, queueAmount []float64
	// var stackPos, queuePos []int

	QueueAmount, QueuePos = createQueue(data)
	StackAmount, StackPos = createStack(data)
	var length = len(QueueAmount)
	fmt.Println(length)
	for i := 0; i <= len(QueueAmount)-1; i++ {
		var amt float64
		balance := QueueAmount[i]
		for balance != 0 {
			var stackBalance float64
			var pos int
			StackAmount, StackPos, stackBalance, pos = popStack(StackAmount, StackPos)
			fmt.Printf("pos, balance ==> %v, %v\n", StackPos, stackBalance)

			if stackBalance >= balance {
				amt = balance
				stackBalance = stackBalance - balance
				balance = 0
				StackAmount, StackPos = pushStack(StackAmount, StackPos, stackBalance, pos)
			} else {
				amt = stackBalance
				balance = balance - stackBalance
				stackBalance = 0
			}
			// if stackBalance != 0 {
			// 	StackAmount, StackPos = pushStack(StackAmount, StackPos, stackBalance, pos)
			// }
			settled = append(settled, Settlement{From: data[pos].Name, To: data[QueuePos[i]].Name, Amount: amt})
			fmt.Printf("settled ==> %v\n\n", settled)

		}
	}

	// for _, per := range data {
	// 	if per.Data.InitialAmount > per.Data.SplitAmount {
	// 		balance := per.Data.InitialAmount - per.Data.SplitAmount
	// 		fmt.Printf("balance ==> %v\n", balance)
	// 		f := 2
	// 		for f != 0 {
	// 			negPos, negBalance := findNegative(data)
	// 			fmt.Printf("pos, balance ==> %v, %v\n", negPos, negBalance)

	// 			if negPos == 100 {
	// 				fmt.Printf("error ==> ")
	// 				return
	// 			}
	// 			var amt float64
	// 			if negBalance >= balance {
	// 				amt = negBalance - balance
	// 				balance = negBalance - balance
	// 				data[negPos].Data.InitialAmount = per.Data.SplitAmount
	// 			} else {
	// 				amt = negBalance
	// 				fmt.Printf("amount %v\n", amt)
	// 				balance = balance - negBalance
	// 				data[negPos].Data.InitialAmount = negBalance
	// 				// amt
	// 			}
	// 			s := Settlement{
	// 				From:   data[negPos].Name,
	// 				To:     per.Name,
	// 				Amount: amt,
	// 			}
	// 			fmt.Printf("settled ==> %v\n\n", s)
	// 			settled = append(settled, s)

	// 			// data[negPos].Data.InitialAmount += balance
	// 			if balance == 0 {
	// 				break
	// 			}
	// 			f = f - 1
	// 			// balance = 0
	// 		}

	// 	}
	// }
	c.JSON(http.StatusOK, settled)
}

func createQueue(data []Person) ([]float64, []int) {
	queueAmount := []float64{}
	queuePos := []int{}
	for i, per := range data {
		if per.Data.InitialAmount > per.Data.SplitAmount {
			queuePos = append(queuePos, i)
			queueAmount = append(queueAmount, per.Data.InitialAmount-per.Data.SplitAmount)
		}
	}
	return queueAmount, queuePos
}

func createStack(data []Person) ([]float64, []int) {

	stackAmount := []float64{}
	stackPos := []int{}
	for i, per := range data {
		if per.Data.InitialAmount < per.Data.SplitAmount {
			stackPos = append(stackPos, i)
			stackAmount = append(stackAmount, per.Data.SplitAmount-per.Data.InitialAmount)
		}
	}
	return stackAmount, stackPos
}

func pushStack(stackAmount []float64, stackPos []int, amount float64, pos int) ([]float64, []int) {

	stackAmount = append(stackAmount, amount)
	stackPos = append(stackPos, pos)

	return stackAmount, stackPos
}

func popStack(stackAmount []float64, stackPos []int) ([]float64, []int, float64, int) {

	lastAmount := stackAmount[len(stackAmount)-1]
	stackAmount = stackAmount[:len(stackAmount)-1]

	lastPos := stackPos[len(stackPos)-1]
	stackPos = stackPos[:len(stackPos)-1]

	return stackAmount, stackPos, lastAmount, lastPos
}

// func findNegative(data []Person) (int, float64) {
// 	var flag int
// 	for j := 0; j <= len(data)-1; j++ {
// 		if data[j].Data.InitialAmount < data[j].Data.SplitAmount {
// 			flag = j
// 			fmt.Printf("name ==> %v %v %v\n", data[j].Name, data[j].Data.InitialAmount, data[j].Data.SplitAmount)
// 			break
// 		}
// 	}
// 	return flag, data[flag].Data.SplitAmount - data[flag].Data.InitialAmount
// }

func main() {

	r := gin.Default()

	r.GET("/data", calculator)
	r.Run()
}
