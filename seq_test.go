package sequence

import (
	"sync"
	"testing"
	"time"
)

func TestGen(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 20000; i++ {
		wg.Add(1)

		// time.Sleep(time.Duration(rand.Int31n(5000))*time.Millisecond)
		go func() {
			base, seq, err := SecondSequence.Gen()
			if err != nil {
				time.Sleep(1 * time.Second)
				// time.Sleep(500 * time.Millisecond)
				base, seq, _ = SecondSequence.Gen()
			}
			t.Log(base, seq)
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkSeq_Gen(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SecondSequence.Gen()
	}
}
