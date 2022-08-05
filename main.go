package main

import (
	"L2WB/pattern"
	"fmt"
)

func main() {
	w1 := pattern.NewStockWorker("IGOOOOOOR", 100)
	w2 := pattern.NewStockWorker("ILLYAA", 150)
	w3 := pattern.NewStockWorker("Misha", 155) // просто миша
	w1.SetNext(w2)
	w2.SetNext(w3)
	work1 := pattern.NewWorkinStock(80, "1 kg ch")
	work2 := pattern.NewWorkinStock(120, "2 kg ch")
	work3 := pattern.NewWorkinStock(152, "3 kg ch")
	work4 := pattern.NewWorkinStock(9999, "spawn lightning")
	w1.Work(work1)
	w1.Work(work2)
	w1.Work(work3)
	w1.Work(work4)
	fmt.Println("\n")
	work1.Work(w1)
	work2.Work(w1)
	work3.Work(w1)
	work4.Work(w1)
}
