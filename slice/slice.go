package main

import "fmt"

func main() {
	sli := []int{1, 2, 3}
	fmt.Printf("init\n slice %v len:%d,cap:%d addr:%p \n", sli, len(sli), cap(sli), sli)
	appendSli(sli)
	fmt.Printf("after appendSli\n  slice %v len:%d,cap:%d addr:%p\n", sli, len(sli), cap(sli), sli)
	appendSliP(&sli)
	fmt.Printf("after appendSliP\n slice %v len:%d,cap:%d addr:%p\n", sli, len(sli), cap(sli), sli)
	for _, v := range sli {
		fmt.Println(v)
	}
}

func appendSli(sli []int) {
	sli = append(sli, 4)
	fmt.Printf("appendSli\n slice %v len:%d,cap:%d addr:%p\n", sli, len(sli), cap(sli), sli)
}

func appendSliP(sli *[]int) {
	*sli = append(*sli, 4, 5)
	fmt.Printf("appendSliP\n slice %v len:%d,cap:%d addr:%p\n", *sli, len(*sli), cap(*sli), *sli)
}
