package main
 

import "fmt"


func main(){
	//var  daromad float64
	//var xarajat float64 
	//var soliqStavka float64
   
    daromad := getUserInput("Daromad: ")
    xarajat := getUserInput("Xarajat: ")
	 soliqStavka := getUserInput("SoliqStavka: ")
     
	 ebt, sofDaromad, ratio:= calculateFinancials(daromad,xarajat,soliqStavka)
	
	fmt.Printf("%f.1\n",ebt)
   fmt.Printf("%f.1\n",sofDaromad)
	fmt.Printf("%f.3\n",ratio)

}
func calculateFinancials(daromad,xarajat,soliqStavka float64) (float64,float64,float64){
	ebt := daromad - xarajat 
	sofDaromad := ebt *(1 - soliqStavka/100)
   ratio := ebt / sofDaromad
	return ebt,sofDaromad,ratio

}
 

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
 }
 