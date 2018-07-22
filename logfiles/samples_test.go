package logfiles

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	ser1 := New("Ser1")
	ser1.show()
}

func TestAdd(t *testing.T) {
	ser2 := New("Ser2")
	for i := 0; i < 100; i++ {
		newsamp := Sample{time.Now(), float32(i)}
		ser2.Add(newsamp)
		time.Sleep(1e8)
	}
	ser2.show()
}

func TestRange(t *testing.T) {
	ser3 := New("Ser3")
	for i := 0; i < 100; i++ {
		newsamp := Sample{time.Now(), rand.Float32()}
		ser3.Add(newsamp)
		time.Sleep(1e8)
	}
	min, max := ser3.Range()
	fmt.Printf("Min %f Max %f\n", min, max)
}
