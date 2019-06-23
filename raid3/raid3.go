package main

import (
	"os"
	"fmt"
	"strconv"
)

func main(){
	if len(os.Args) > 3{
		os.Args = os.Args[1:]
		param1, err1 := strconv.Atoi(os.Args[1])
		param2, err2 := strconv.Atoi(os.Args[2])

		for i,v := range os.Args{
			if i == 0 {
				os.Args[i] = v[2:]
				os.Args[i] += ".go"
			}
		}

		if len(os.Args) == 3{
			// for i,v := range os.Args{
			// 	if i == 0 {
			// 		os.Args[i] = v[2:]
			// 		os.Args[i] += ".go"
			// 	}
			// }
			if (err1 != nil) || (err2 != nil) {
				return 
			}else{
				fmt.Print(isExist(os.Args[0], param1, param2))
			}
		}else if len(os.Args) == 5 && os.Args[3] == "|" && os.Args[4] == "./raid3" {
			
			tabAssociatif := map[int]string{
				0:raid1a(param1, param2),
				1:raid1b(param1, param2),
				2:raid1c(param1, param2),
				3:raid1d(param1, param2),
				4:raid1e(param1, param2),
			}
			tabAssociatif2 := map[int]string{
				0:"[raid1a]",
				1:"[raid1b]",
				2:"[raid1c]",
				3:"[raid1d]",
				4:"[raid1e]",
			}
			// fmt.Print("TabAssociatif :\n",tabAssociatif)
			// fmt.Println()
			// fmt.Println("isExist -->",isExist(os.Args[0], param1, param2))
			listFonction := []string{}

			for i, v := range tabAssociatif{
				if isExist(os.Args[0], param1, param2) == v {
					listFonction = append(listFonction,tabAssociatif2[i])
				}
			}

			for i,_ := range listFonction{
				if i < len(listFonction) -1  {
					fmt.Print(listFonction[i], " [",param1,"]"," [",param2,"] || ")
				}else{
					fmt.Println(listFonction[i], " [",param1,"]"," [",param2,"]")
				}
			}
		}
	}
	
}

func isExist (str string, param1, param2 int)string{
	_, err := os.Open(str)
	if err != nil {
		return "Not a Raid function"
	}else{
		str = str[:len(str)-3]
		// return str
		if str == "raid1a" {
			return (raid1a(param1, param2))
		}else if str == "raid1b" {
			return (raid1b(param1, param2))
		}else if str == "raid1c" {
			return (raid1c(param1, param2))
		}else if str == "raid1d" {
			return (raid1d(param1, param2))
		}else if str == "raid1e" {
			return (raid1e(param1, param2))
		}
		return ""
	}
}

func raid1a(x,y int) string{
	varPrint :=""
	if x>0 && y>0 {

		for i := 1; i <= y; i++ {
			if x==1 && y!=1 {
			for j := 1; j <= y; j++ {

				if j==1 || j==y {
					varPrint+= fmt.Sprintf("o\n")
				}else if j!=1 && j!=y {
					varPrint+= fmt.Sprintf("|\n")
				}
			}
			break
			}else{
				for j := 1; j <= x; j++ {
					if (i==1 && j==1) || (i==1&&j==x) || (i==y && j==1) || (i==y && j==x)  {
						varPrint+= fmt.Sprintf("o")
					}else if (i==1 && (j!=1 || j!=x)) || (i==y && (j!=x || j!=x)) {
						varPrint+=fmt.Sprintf("-")
					}else if (i!=1 && j==1) || (i!=y && j==x){
						varPrint+= fmt.Sprintf("|")
					}else{
						varPrint+= fmt.Sprintf(" ")
					}
				}
			}
			varPrint+=fmt.Sprintf("\n")
		}
	}
	return varPrint
}

func raid1b(x,y int) string{
	varPrint := ""
	if x>0 && y>0 {
		for i := 1; i <= y; i++ {
			if x==1 && y!=1 {
				for j := 1; j <= y; j++ {
					if j==1 {
						varPrint+=fmt.Sprintf("/")
					}else if j==y {
						varPrint+=fmt.Sprintf("\n\\\n")
					}else if j!=1 && j!=y {
						varPrint+=fmt.Sprintf("\n*")
					}
				}
				break
			}else if y==1 && x!=1 {
				for j := 1; j <= x; j++ {
					if j==1 {
						varPrint+=fmt.Sprintf("/")
					}else if j==x {
						varPrint+=fmt.Sprintf("\\\n")
					}else if j!=1 && j!=y {
						varPrint+=fmt.Sprintf("*")
					}
				}
				break
			}else{
				for j := 1; j <= x; j++ {
					if (i==1 && j==1) || (i==y && j==x)  {
						varPrint+=fmt.Sprintf("/")
					}else if (i==1&&j==x) || (i==y && j==1) {
						varPrint+=fmt.Sprintf("\\")
					}else if (i==1 && (j!=1 || j!=x)) || (i==y && (j!=x || j!=x)) {
						varPrint+=fmt.Sprintf("*")
					}else if (i!=1 && j==1) {
						varPrint+=fmt.Sprintf("*")
					}else if (i!=y && j==x){
						varPrint+=fmt.Sprintf("*")
					}else{
						varPrint+=fmt.Sprintf(" ")
					}
				}
			}
			varPrint+=fmt.Sprintf("\n")
		}
	}
	return varPrint
}

func raid1c(x,y int) string{
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
			}else{
			//ligne y
				for j := 1; j <= x; j++ {
					//colonne x
					
					if (i==1 && j==1) || (i==1 && x==j) {
						varPrint+=fmt.Sprintf("A")
					}else if (i==y && j==1) || (i==y && j==x) {
						varPrint+=fmt.Sprintf("C")
					}else if (i==1 && (j!=1 || j!=x)) || (i==y && (j!=1 || j!=x)) || (j==1 && (i!=1 || i==y)) || (j==x && (i!=1 || i!=y)){
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

func raid1d(x,y int) string{
	varPrint := ""
	if x>0 && y>0 {
		for i := 1; i <= y; i++ {
			//ligne y
			for j := 1; j <= x; j++ {
				//colonne x
				if (i==1 && j==1) || (i==y && j==1)  {
					varPrint+=fmt.Sprintf("A")
				}else if (i==1&&j==x) ||  (i==y && j==x){
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
	return varPrint
}

func raid1e(x,y int) string{
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
