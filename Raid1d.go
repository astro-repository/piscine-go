package main

func main(){
	Raid1a(1,5)
}

func Raid1a(x,y int){
	for i := 1; i <= y; i++ {
		//ligne y
		for j := 1; j <= x; j++ {
			//colonne x
			if (i==1 && j==1) || (i==y && j==1)  {
				print("A")
			}else if (i==1&&j==x) ||  (i==y && j==x){
				print("C")
			}else if (i==1 && (j!=1 || j!=x)) || (i==y && (j!=x || j!=x)) {
				print("B")
			}else if (i!=1 && j==1) {
				print("B")
			}else if (i!=y && j==x){
				print("B")
			}else{
				print(" ")
			}
		}
		print("\n")
	}
}