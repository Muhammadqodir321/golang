package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)
const accountBalanceFile = "balance.txt"
 
func getBalanceFromFile() (float64,error) {
	 data, err:= os.ReadFile( accountBalanceFile)

	 if err != nil{
      return 1000, errors.New("failed to find balance file")
	 }


	 balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText,64)

	if err != nil{
	  return 1000, errors.New("failed to parse stored balance value")
	}
	return balance, nil
 }


func writeBalanceToFile(balance float64){
balanceText :=fmt.Sprint(balance)
	os.WriteFile( accountBalanceFile,[]byte(balanceText), 0644)
}

func main(){
	var accountBalance,err = getBalanceFromFile()
	if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
	}
	fmt.Println("Welcome to Go bank!")
	for{

fmt.Println("What do you want to do?")
fmt.Println("1.Check balance")
fmt.Println("2.Deposit money")
fmt.Println("3.Withdraw money")
fmt.Println("4.Exit")


var choice int
fmt.Scan(&choice)
fmt.Println("Your choice:",choice)

switch choice{
case 1:
	fmt.Println("Your balance:",accountBalance) 
case 2:
	fmt.Print("Your deposit:")
	var depositAmount float64 
	fmt.Scan(&depositAmount)
	if depositAmount <= 0{
		fmt.Println("Invalid amount.Must be greater than 0.")
		//return
	continue
	}
	accountBalance += depositAmount
	fmt.Println("Your balance updated! New amount:",accountBalance)
   writeBalanceToFile(accountBalance)
case 3:
fmt.Print("Withdraw money:")
var withdrawAmount float64 
fmt.Scan(&withdrawAmount)
if withdrawAmount <= 0{
	fmt.Println("Invalid amount.Must be greater than 0.")
	continue
}
if withdrawAmount > accountBalance{
	fmt.Println("Sorry,you can't withdraw than you have!")
	continue
}

accountBalance -= withdrawAmount
fmt.Println("Your balance updated! New amount:",accountBalance)
writeBalanceToFile(accountBalance)
default:
fmt.Println("Goodbye!")
fmt.Println("Thanks for choosing our bank!")
return
//break

}

}
}