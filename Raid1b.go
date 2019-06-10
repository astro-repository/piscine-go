package main

func main(){
	Raid1a(1,5)
}

func Raid1a(x,y int){
	for i := 1; i <= y; i++ {
		if x==1 && y!=1 {
			for j := 1; j <= y; j++ {
				if j==1 {
					print("/")
				}else if j==y {
					print("\n\\\n")
				}else if j!=1 && j!=y {
					print("\n*")
				}
			}
			break
		}else if y==1 && x!=1 {
			for j := 1; j <= x; j++ {
				if j==1 {
					print("/")
				}else if j==x {
					print("\\\n")
				}else if j!=1 && j!=y {
					print("*")
				}
			}
			break
		}else{
			for j := 1; j <= x; j++ {
				if (i==1 && j==1) || (i==y && j==x)  {
					print("/")
				}else if (i==1&&j==x) || (i==y && j==1) {
					print("\\")
				}else if (i==1 && (j!=1 || j!=x)) || (i==y && (j!=x || j!=x)) {
					print("*")
				}else if (i!=1 && j==1) {
					print("*")
				}else if (i!=y && j==x){
					print("*")
				}else{
					print(" ")
				}
			}
		}
		print("\n")
	}
}