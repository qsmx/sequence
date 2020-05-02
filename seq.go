package sequence

import (
	"errors"
	"sync"
	"time"
)

// SeqHandler sequence handler
type SeqHandler interface {
	Gen() (base interface{}, seq int, err error)
}

type tSeq struct {
	seq    int // sequence id
	maxSeq int // max sequence id

	base   interface{}        // base
	baseFn func() interface{} // base function

	lock sync.Mutex
}

// ErrOverflow overflow
var ErrOverflow = errors.New("out of range")

// preview defined
var (
	GlobalSequence = New(-1, func() interface{} { return 0 })                 // 全局序列
	SecondSequence = New(-1, func() interface{} { return time.Now().Unix() }) // 秒级序列
)

// New new sequence handler
func New(maxSeq int, fn func() interface{}) SeqHandler {
	return &tSeq{
		seq:    0,
		maxSeq: maxSeq,
		base:   fn(),
		baseFn: fn,
	}
}

// Gen generate sequence
func (s *tSeq) Gen() (interface{}, int, error) {
	base := s.baseFn()

	s.lock.Lock()

	if s.base != base {
		s.seq = 0
		s.base = base
	}

	s.seq++

	if s.maxSeq != -1 && s.seq >= s.maxSeq {
		s.lock.Unlock()
		return nil, 0, ErrOverflow
	}

	defer s.lock.Unlock()
	return s.base, s.seq, nil
}
