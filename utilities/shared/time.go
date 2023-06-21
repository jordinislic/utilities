package shared

import (
	"fmt"
	"time"
)

func ChronoMilli(t1 int64) {
	t2 := time.Now().UnixMilli()
	fmt.Println("Trascorsi ", t2-t1, "Millisecondi")
}

func ChronoMicro(t1 int64) {
	t2 := time.Now().UnixMicro()
	fmt.Println("Trascorsi ", t2-t1, "Microsecondi")
}

func ChronoNano(t1 int64) {
	t2 := time.Now().UnixNano()
	fmt.Println("Trascorsi ", t2-t1, "Nanosecondi")
}

func Chrono(t1 int64) {
	t2 := time.Now().Unix()
	fmt.Println("Trascorsi ", t2-t1, "Secondi")
}
