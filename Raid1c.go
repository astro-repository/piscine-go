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
					print("C")
				}else if j!=1 && j!=y {
					print("B")
				}
				print("\n")
			}
			break
		}else{
		//ligne y
		for j := 1; j <= x; j++ {
			//colonne x
			
			if (i==1 && j==1) || (i==1 && x==j) {
				print("A")
			}else if (i==y && j==1) || (i==y && j==x) {
				print("C")
			}else if (i==1 && (j!=1 || j!=x)) || (i==y && (j!=1 || j!=x)) || (j==1 && (i!=1 || i==y)) || (j==x && (i!=1 || i!=y)){
				print("B")
			}else{
				print(" ")
			}
		}
		print("\n")
	}
}
	
}