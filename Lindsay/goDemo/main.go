//based on this demo https://www.youtube.com/watch?v=C8LgvuEBraI&t=652s

//The first line needs to be the name of the package most of the time main will be the package name
package main
//the next step we need to import fmt which will give us in and out functions like c stdio
import (
    "fmt"
    "math/rand"
)

func main() {
  //use fmt to print and P in print is uppercase
  fmt.Println("Hello Big Data Class")

  //there are two ways to decare varibles
  //way one the Object oriented way
  var x int =5
  //way to the "old school way" go will figure out the type
  y:=7
  fmt.Println("x",x,"y",y)
  //lets use a for loop to build a slice or in other languages an array list
  //there is no while loop in go only for.
  sli:=[]int{}
  //for is similar to java
  for i:=0;i<5;i++{
    //random is similar to java rand.Intn and range
    sli=append(sli,rand.Intn(100))
    //every append creates a new array
  }
  //now we have a random list. I want to know if the sum of that list is even or not lets make a function for that
   fmt.Println(randomSum(sli))

 
}
//functions
//func name(parameter)return can return more than one thing
func randomSum(arr []int)(int,string){
// if the sum is even lets return the sum and a string saying its even otherwise lets return the sum and a string saying its odd
sum:=0
//enhanced for loop
 for index,value :=range arr{
   fmt.Println("index:",index,"value",value)
    sum =sum+value
  }
  //condition statment is very similar to other language conditional statments
  if sum %2==0{
    return sum, "Its even";
  }
  return sum, "Its odd"
  //even though multiple return is fun the main use is to return errors since go does not have exceptions
}
//few things to note there are no classes it uses structs like C. It also has pointers like C
