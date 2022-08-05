package pattern

import "fmt"

// интерфейс некоторого работника
type _Worker interface {
	GetWorkLevel() int
	GetWorkName() string
	Next() _Worker
	Work(_Work) bool
}

// конкретный работник
type StockWorker struct {
	WorkLevel   int
	name        string
	_nextWorker *StockWorker
}

// реализуем интерфейс
func (sw *StockWorker) GetWorkLevel() int {
	return sw.WorkLevel
}

func (sw *StockWorker) GetWorkName() string {
	return sw.name
}

// напишем для удобства фуекцию которая будет вставлять сама дальше в цепочку ниже воркера
func (sw *StockWorker) SetNext(next *StockWorker) {
	if sw._nextWorker != nil {
		nx := sw._nextWorker
		for nx._nextWorker != nil {
			nx = nx._nextWorker
		}
		nx._nextWorker = next
		return
	}
	sw._nextWorker = next
}

// сама работа аля смотрим сложность работы и перебираем по цепочке кто это может взять
func (sw *StockWorker) Work(w _Work) bool {
	if sw.WorkLevel < w.GetWorkСomplexity() {
		swnext := sw._nextWorker
		for swnext != nil {
			if swnext.WorkLevel >= w.GetWorkСomplexity() {
				fmt.Println(sw.name, " Worked in stock by pattern ", w.GetPattern(), " work level ", w.GetWorkСomplexity(), " worker level ", swnext.WorkLevel)
				return true
			}
			swnext = swnext._nextWorker
		}
	} else {
		fmt.Println(sw.name, " Worked in stock by pattern ", w.GetPattern(), " work level ", w.GetWorkСomplexity(), " worker level ", sw.WorkLevel)
		return true
	}
	fmt.Println("no worker for this work")
	return false
}

// безопастный (нулевой интерфейс не равень интерфейсу нулевого объекта) получатель след воркеров
func (sw *StockWorker) Next() _Worker {
	if sw._nextWorker == nil {
		return nil
	} else {
		return sw._nextWorker
	}
}

func NewStockWorker(name string, worklevel int) *StockWorker {
	return &StockWorker{name: name, WorkLevel: worklevel}
}

// некая работа
type _Work interface {
	GetWorkСomplexity() int
	Work(_Worker) bool
	GetPattern() string
}

type WorkinStock struct {
	WorkСomplexity int
	patt           string
	//...
}

// реализуем интерфейс
func (ws *WorkinStock) GetWorkСomplexity() int {
	return ws.WorkСomplexity
}

func (ws *WorkinStock) GetPattern() string {
	return ws.patt
}

// работа может работать во чудо ... при передаче ей работника тут где то логика потерялась но интересно же...
// аналогично что и нормальный способ просто хотелось проверить обратную связь
func (ws *WorkinStock) Work(w _Worker) bool {
	if w.GetWorkLevel() < ws.WorkСomplexity {
		w = w.Next()
		for w != nil {
			if w.GetWorkLevel() >= ws.WorkСomplexity {
				fmt.Println(w.GetWorkName(), " Worked in stock by pattern ", ws.patt, " work level ", ws.WorkСomplexity, " worker level ", w.GetWorkLevel())
				return true
			}
			w = w.Next()
		}
	} else {
		fmt.Println(w.GetWorkName(), " Worked in stock by pattern ", ws.patt, " work level ", ws.WorkСomplexity, " worker level ", w.GetWorkLevel())
		return true
	}
	fmt.Println("no worker for this work")
	return false
}

func NewWorkinStock(workcomplexity int, pattern string) *WorkinStock {
	return &WorkinStock{WorkСomplexity: workcomplexity, patt: pattern}
}

// паттерн цепочка обязанностей я тут сделал и обратный ход для работы (работа сама переберет worker(ов))
// плюсики удобно обработчики сами управляют передать ли или обработать задачку, гибкая реализация позволяет не сильно упарываться в 100500 строк когда в классе итд
// работа может быть не выполнена по той или иной причине :D.

// w1 := pattern.NewStockWorker("IGOOOOOOR", 100)
// w2 := pattern.NewStockWorker("ILLYAA", 150)
// w3 := pattern.NewStockWorker("Misha", 155) // просто миша
// w1.SetNext(w2)
// w2.SetNext(w3)
// work1 := pattern.NewWorkinStock(80, "1 kg ch")
// work2 := pattern.NewWorkinStock(120, "2 kg ch")
// work3 := pattern.NewWorkinStock(152, "3 kg ch")
// work4 := pattern.NewWorkinStock(9999, "spawn lightning")
// w1.Work(work1)
// w1.Work(work2)
// w1.Work(work3)
// w1.Work(work4)
// fmt.Println("\n")
// work1.Work(w1)
// work2.Work(w1)
// work3.Work(w1)
// work4.Work(w1)
// OUT:
// IGOOOOOOR  Worked in stock by pattern  1 kg ch  work level  80  worker level  100
// IGOOOOOOR  Worked in stock by pattern  2 kg ch  work level  120  worker level  150
// IGOOOOOOR  Worked in stock by pattern  3 kg ch  work level  152  worker level  155
// no worker for this work

// IGOOOOOOR  Worked in stock by pattern  1 kg ch  work level  80  worker level  100
// ILLYAA  Worked in stock by pattern  2 kg ch  work level  120  worker level  150
// Misha  Worked in stock by pattern  3 kg ch  work level  152  worker level  155
// no worker for this work
