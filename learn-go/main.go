package main 

import (
	"fmt"
	"errors"
)

// func main(){
//     ans,err := devide(1,020)

// 	if err!=nil{
// 		fmt.Print("Error",err)
// 		return
// 	}

// 	fmt.Print(ans)
// }

// func main(){
// 	defer fmt.Println("first")
// 	 fmt.Println("2nd")

// 	 prices := []int {1,2,3,4,5}
// 	 prices = append(prices,6)
// 	 fmt.Println(prices)
// 	 fmt.Println(len(prices))
// 	 fmt.Println(prices[3])
// 	fmt.Println("loop ")
// 	 for i ,v := range prices {
// 		fmt.Println(i,v)
// 	 }

// 	 stocks := map[string]float64{
// 		"apple":100,
// 		"banana":200,
// 		"orange":300,
// 	 }

// 	 stocks["owner"]=5
// 	 price,ok := stocks["apple"]

// 	 if !ok{
// 		fmt.Println("apple not found")
// 	 }

// 	 fmt.Println("apple found the stocks we have ",price)

// 	 for k,v := range stocks{
// 		fmt.Println(k,v)
// 	 }
// 	fmt.Println(stocks)

// 	fmt.Println("apple stcks we have this ",stocks["apple"])

// 	hellow()
// 	min,max,avg,err := stats([]float64{1,2,3,4,5})
// 	fmt.Println("min",min)
// 	fmt.Println("max",max)
// 	fmt.Println("avg",avg)
// 	fmt.Println("err",err)

// }
func stats(prices []float64) (min, max, avg float64, err error){
	if len(prices)==0{
		return 0,0,0,errors.New("empty slice")
	}

	min = prices[0]
	max = prices[0]
	avg = 0
	for _,v := range prices{
		if v>max{
			max=v
		}
		if v<min{
			min=v
		}
		avg+=v
	}
	avg = avg/float64(len(prices))
	return 
}

func hellow(){
	fmt.Println("hello")
	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	fmt.Println(add(1,2))
	fmt.Print(sub(1,2))
}

func add(a int,b int) int {
	return a+b
}

func sub(a,b int) int {
	return a-b
}


func devide(a,b float64) (float64,error){
	if b==0{
		return 0,errors.New("cannot devide by zero")
	}

	return a/b,nil
}


type Stock struct {
    Symbol string
    Price  float64
    Qty    int
}


func main (){
	s1:= Stock {
		Symbol:"apple",
		Price:100,
		Qty:10,
	}

	var s2 Stock

	s2.Symbol="apple"
	s2.Price=100
	s2.Qty=10

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s2.totalValue())
	s1.addQuntity(10)
	fmt.Println(s1)


	s4:= newStock("INFY", 100)
	s4.addQuntity(20)

	fmt.Println(s4)
}


func (s Stock ) totalValue() float64{
	return s.Price * float64(s.Qty)
}

func (s *Stock) addQuntity(n int){
	s.Qty += n 
}

func newStock(symbol string, price float64 ) *Stock{
	return &Stock{Symbol:symbol,Price:price}
}