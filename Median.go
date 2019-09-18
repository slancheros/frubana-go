package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var listToCalculate []int
	reader :=bufio.NewReader(os.Stdin)
	print("Enter number of operations to perform:")
	strNumberOperations,_:= reader.ReadString('\n')
	numberOperations, err := strconv.Atoi(strNumberOperations)
	if err!=nil{
		print("Invalid number of operations" )
	}
	for i:=1; i<=numberOperations; i++{
		command, value := handle_operation()
		if strings.Compare(command,"a") == 0{
			addToMedian(listToCalculate,value)
		}
		if strings.Compare(command,"r") == 0{
			addToMedian(listToCalculate,value)
		}
	}
}

func calculateMedian(numberList []int) float64{

	sort.Ints(numberList)
	length := len(numberList)
	medianCalculated := -1.0

	if length != 0 {
		if length == 1 {
			medianCalculated = float64(numberList[0])
		} else if  length % 2 == 0{
			secondIndex := length / 2
			firstIndex  := secondIndex - 1
			medianCalculated = float64((numberList[firstIndex] + numberList[secondIndex]) / 2)
		}else if length % 2 != 0{
            medianIndex := math.Floor( float64(length) / 2.0 )
            medianCalculated = float64(numberList[int(medianIndex)])
		}
	}
   return medianCalculated
}

func isIntegral( val float64) bool{
	return val == float64(int(val))
}

func addToMedian(numberList []int, addValue int){
	numberList = append(numberList, addValue)
	medianValue := calculateMedian(numberList)
	printMedian(medianValue)
}

func removeToMedian(numberList[]int, oldValue int) {
	if len(numberList) != 0 {

		if contains(numberList, oldValue) {
			position := getPosition(numberList, oldValue)
			numberList = removeElement(numberList, position)
			if len(numberList) == 0 {
				printError()
			}else {
				median := calculateMedian(numberList)
				printMedian(median)
			}
		}else {
			printError()
			}
	}else {
		printError()
	}
}


func getPosition ( numberList[]int, valueToFind int) int{
	for position, v := range numberList {
		if v == valueToFind {
			return position
		}
	}
	return -1
}

func printMedian( medianValue float64 ){
	strMedian :=""
	if isIntegral(medianValue){
		medianInt := int(medianValue)
		strMedian =fmt.Sprintf("%d", medianInt)

	}else {
		strMedian = fmt.Sprintf("%f", medianValue)
	}
	print("This is the median of the list of numbers: " + strMedian)
}

func printError(){
	print("Wrong!")
}

func printInvalidOperation(){
	print("Invalid operation")
}

func contains(list []int, number int) bool {
	for _, item := range list {
		if item == number {
			return true
		}
	}
	return false
}

func removeElement(numberList []int, index int) []int {
	return append(numberList[:index], numberList[index+1:]...)
}

func handle_operation()(string, int){
	reader :=bufio.NewReader(os.Stdin)
	print("Enter operation:")
	operation,_:= reader.ReadString('\n')
	var command string
	var value_var int
	var err error

	if len(operation) ==0{
		printInvalidOperation()
	}else{
		tokensOperation := strings.Fields(operation)
		if len(tokensOperation) == 2 {
			command = tokensOperation[0]
			value_var, err = strconv.Atoi(tokensOperation[1])
			if err !=nil{
				printInvalidOperation()
			}
		}else{
			printInvalidOperation()
		}
	}
   return command, value_var
}