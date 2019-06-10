package main

func main(){
	Raid1a(1,5)
}

func Raid1a(x,y int){
	for i := 1; i <= y; i++ {
		if x==1 && y!=1{
			for j := 1; j <= y; j++ {

				if j==1 {
					print("A")
				}else if j==y {
					print("\nC")
				}else if j!=1 && j!=y {
					print("\nB")
				}
			}
			break
		}else if y==1 && x!=1 {
			for j := 1; j <= x; j++ {
				if j==1 {
					print("A")
				}else if j==y {
					print("C")
				}else if j!=1 && j!=x {
					print("B")
				}else{
					print("C")
				}
			}
			break
		}else{
		//ligne y
		for j := 1; j <= x; j++ {
			//colonne x
			
			if (i==1 && j==1) || (i==y && j==x)  {
				print("A")
			}else if (i==1&&j==x) || (i==y && j==1) {
				print("C")
			}else if (i==1 && (j!=1 || j!=x)) || (i==y && (j!=x || j!=x)) {
				print("B")
			}else if (i!=1 && j==1) {
				print("\nB")
			}else if (i!=y && j==x){
				print("B\n")
			}else{
				print(" ")
			}
		}
	}
}
	print("\n")
}