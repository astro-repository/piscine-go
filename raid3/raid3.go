package main

import (
	"os"
	"fmt"
	// "io"
	"log"
	"strings"
	// "strconv"
	// "os/exec"
	// "io/ioutil"
	"bufio"
	// "bytes"
)
const FICHIER = "log.txt"

func main(){

	if len(os.Args)==1{
		file, err := os.Open(FICHIER)
		file2, _ := os.Open(FICHIER)
		file3, _ := os.Open(FICHIER)

	    if err != nil {
	        log.Fatal(err)
	    }
	    defer file.Close()

	    scanner := bufio.NewScanner(file)
	    scanner2 := bufio.NewScanner(file2)
	    scanner3 := bufio.NewScanner(file3)
	    compteur:=0

	    for scanner.Scan() {
	    	compteur++
	    }

	    i:=0
	    lastElmt := ""

	    for scanner2.Scan() {
	    	i++
	    	if compteur == i{
	    		lastElmt = fmt.Sprintf(scanner.Text())
	    	}
	    }

	    splitLastElmt := strings.Index(lastElmt," ")
	    lastElmt2 := []rune(lastElmt)
	    lastElmt = string(lastElmt2[splitLastElmt+1:])

	    resultat := []string{}

	    for scanner3.Scan() {
	    	if strings.Index(scanner3.Text(),lastElmt) != -1{
	    		resultat = append(resultat,scanner3.Text())
	    	}
	    }

	    fmt.Println(len(resultat))

	    for i,v := range resultat{
	    	if i<len(resultat)-1{
	    		fmt.Print(v," || ")
	    	}else{
	    		fmt.Println(v)
	    	}
	    }
	    if err := scanner.Err(); err != nil {
	        log.Fatal(err)
	    }
		if err != nil{
			return
		}

	}else{
		fmt.Println("Not a Raid function")
	}
}

// --- FAIL: TestRaid3 (2.40s)
//     raid3_test.go:82: ./student/raid3 raid1a/raid1a  prints "" instead of "[raid1a] [37] [19] || [raid1b] [37] [19] || [raid1c] [37] [19] || [raid1d] [37] [19] || [raid1e] [37] [19]\n"
//     raid3_test.go:82: ./student/raid3 raid1b/raid1b  prints "" instead of "[raid1a] [5] [26] || [raid1b] [5] [26] || [raid1c] [5] [26] || [raid1d] [5] [26] || [raid1e] [5] [26]\n"
//     raid3_test.go:82: ./student/raid3 raid1c/raid1c  prints "" instead of "[raid1a] [4] [8] || [raid1b] [4] [8] || [raid1c] [4] [8] || [raid1d] [4] [8] || [raid1e] [4] [8]\n"
//     raid3_test.go:82: ./student/raid3 raid1d/raid1d  prints "" instead of "[raid1a] [49] [45] || [raid1b] [49] [45] || [raid1c] [49] [45] || [raid1d] [49] [45] || [raid1e] [49] [45]\n"
//     raid3_test.go:82: ./student/raid3 raid1e/raid1e  prints "" instead of "[raid1a] [10] [33] || [raid1b] [10] [33] || [raid1c] [10] [33] || [raid1d] [10] [33] || [raid1e] [10] [33]\n"
// FAIL
// exit status 1
// FAIL	_/jail	2.403s

// --- FAIL: TestRaid3 (2.29s)
//     raid3_test.go:82: ./student/raid3 raid1a/raid1a  prints "[/jail/student/raid3/raid3]\n" 
//     instead of "[raid1a] [39] [33] || [raid1b] [39] [33] || [raid1c] [39] [33] || [raid1d] [39] [33] || [raid1e] [39] [33]\n"
//     raid3_test.go:82: ./student/raid3 raid1b/raid1b  prints "[/jail/student/raid3/raid3]\n" 
//     instead of "[raid1a] [35] [20] || [raid1b] [35] [20] || [raid1c] [35] [20] || [raid1d] [35] [20] || [raid1e] [35] [20]\n"
//     raid3_test.go:82: ./student/raid3 raid1c/raid1c  prints "[/jail/student/raid3/raid3]\n" 
//     instead of "[raid1a] [27] [4] || [raid1b] [27] [4] || [raid1c] [27] [4] || [raid1d] [27] [4] || [raid1e] [27] [4]\n"
//     raid3_test.go:82: ./student/raid3 raid1d/raid1d  prints "[/jail/student/raid3/raid3]\n" 
//     instead of "[raid1a] [13] [27] || [raid1b] [13] [27] || [raid1c] [13] [27] || [raid1d] [13] [27] || [raid1e] [13] [27]\n"
//     raid3_test.go:82: ./student/raid3 raid1e/raid1e  prints "[/jail/student/raid3/raid3]\n" 
//     instead of "[raid1a] [45] [16] || [raid1b] [45] [16] || [raid1c] [45] [16] || [raid1d] [45] [16] || [raid1e] [45] [16]\n"
// FAIL
// exit status 1
// FAIL	_/jail	2.290s