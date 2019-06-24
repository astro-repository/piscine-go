package main

import (
	"fmt"
	"strconv"
	"os"
)

func main(){
	param1,err1 := strconv.Atoi(os.Args[1])
	param2,err2 := strconv.Atoi(os.Args[2])
	fmt.Println(os.Args)
	if err1 != nil || err2 != nil{
		fmt.Println("Not a Raid function")
		return
	}else{
		file, err := os.OpenFile("log.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0777)
		if err != nil {
			fmt.Println("Not a Raid function")
		return
		}
		raid1e := "[raid1e] ["+os.Args[1]+"] ["+os.Args[2]+"]\n"
		resultByte := []byte(raid1e)
		file.Write(resultByte)
		file.Close()
	}
	fmt.Print(Raid1e(param1,param2))
}

func Raid1e(x,y int) string{
	varPrint := ""
	if x>0 && y>0 {
		for i := 1; i <= y; i++ {
			if x==1 && y!=1{
				for j := 1; j <= y; j++ {

					if j==1 {
						varPrint+=fmt.Sprintf("A")
					}else if j==y {
						varPrint+=fmt.Sprintf("C")
					}else if j!=1 && j!=y {
						varPrint+=fmt.Sprintf("B")
					}
					varPrint+=fmt.Sprintf("\n")
				}
				break
			}else if y==1 && x!=1 {
				for j := 1; j <= x; j++ {
					if j==1 {
						varPrint+=fmt.Sprintf("A")
					}else if j==y {
						varPrint+=fmt.Sprintf("C")
					}else if j!=1 && j!=x {
						varPrint+=fmt.Sprintf("B")
					}else{
						varPrint+=fmt.Sprintf("C")
					}
				}
				varPrint+=fmt.Sprintf("\n")
				break
			}else{
			//ligne y
			for j := 1; j <= x; j++ {
				//colonne x
				if (i==1 && j==1) || (i==y && j==x)  {
					varPrint+=fmt.Sprintf("A")
				}else if (i==1&&j==x) || (i==y && j==1) {
					varPrint+=fmt.Sprintf("C")
				}else if (i==1 && (j!=1 || j!=x)) || (i==y && (j!=x || j!=x)) {
					varPrint+=fmt.Sprintf("B")
				}else if (i!=1 && j==1) {
					varPrint+=fmt.Sprintf("B")
				}else if (i!=y && j==x){
					varPrint+=fmt.Sprintf("B")
				}else{
					varPrint+=fmt.Sprintf(" ")
				}
			}
			varPrint+=fmt.Sprintf("\n")
		}
	}
}
return varPrint
}