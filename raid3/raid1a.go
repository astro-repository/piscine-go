package main

import "fmt"
import "os"
import "strconv"

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
		raid1a := "[raid1a] ["+os.Args[1]+"] ["+os.Args[2]+"]\n"
		resultByte := []byte(raid1a)
		file.Write(resultByte)
		file.Close()
		fmt.Print(Raid1a(param1,param2))
	}
	
}

func Raid1a(x,y int)string{
	varPrint:=""
	if x>0 && y>0 {

		for i := 1; i <= y; i++ {
			if x==1 && y!=1 {
			for j := 1; j <= y; j++ {

				if j==1 || j==y {
					varPrint+=fmt.Sprintf("o\n")
				}else if j!=1 && j!=y {
					varPrint+=fmt.Sprintf("|\n")
				}
			}
			break
			}else{
				for j := 1; j <= x; j++ {
					if (i==1 && j==1) || (i==1&&j==x) || (i==y && j==1) || (i==y && j==x)  {
						varPrint+=fmt.Sprintf("o")
					}else if (i==1 && (j!=1 || j!=x)) || (i==y && (j!=x || j!=x)) {
						varPrint+=fmt.Sprintf("-")
					}else if (i!=1 && j==1) || (i!=y && j==x){
						varPrint+=fmt.Sprintf("|")
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
// documentation_lire_fichier_externe.go
// raid1a
// raid1a.go
// raid1b.go
// raid1c.go
// raid1d.go
// raid1e.go
// raid3
// raid3.go
// documentation_lire_fichier_externe.go
// raid1a
// raid1a.go
// raid1b.go
// raid1c.go
// raid1d.go
// raid1e.go
// raid3
// raid3.go
// documentation_lire_fichier_externe.go
// raid1a
// raid1a.go
// raid1b.go
// raid1c.go
// raid1d.go
// raid1e.go
// raid3
// raid3.go
// documentation_lire_fichier_externe.go
// raid1a
// raid1a.go
// raid1b.go
// raid1c.go
// raid1d.go
// raid1e.go
// raid3
// raid3.go
// documentation_lire_fichier_externe.go
// raid1a
// raid1a.go
// raid1b.go
// raid1c.go
// raid1d.go
// raid1e.go
// raid3
// raid3.go
// documentation_lire_fichier_externe.go
// raid1a
// raid1a.go
// raid1b.go
// raid1c.go
// raid1d.go
// raid1e.go
// raid3
// raid3.go
// documentation_lire_fichier_externe.go
// raid1a
// raid1a.go
// raid1b.go
// raid1c.go
// raid1d.go
// raid1e.go
// raid3
// raid3.go
// documentation_lire_fichier_externe.go
// raid1a
// raid1a.go
// raid1b.go
// raid1c.go
// raid1d.go
// raid1e.go
// raid3
// raid3.go
// documentation_lire_fichier_externe.go
// raid1a
// raid1a.go
// raid1b.go
// raid1c.go
// raid1d.go
// raid1e.go
// raid3
// raid3.go
// documentation_lire_fichier_externe.go
// raid1a
// raid1a.go
// raid1b.go
// raid1c.go
// raid1d.go
// raid1e.go
// raid3
// raid3.go
// documentation_lire_fichier_externe.go
// raid1a
// raid1a.go
// raid1b.go
// raid1c.go
// raid1d.go
// raid1e.go
// raid3
// raid3.go
// documentation_lire_fichier_externe.go
// raid1a
// raid1a.go
// raid1b.go
// raid1c.go
// raid1d.go
// raid1e.go
// raid3
// raid3.go
