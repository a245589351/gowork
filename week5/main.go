package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type SlideWindow interface {
	Total() int64      // 获取窗口内总请求量
	Max() int64        // 获取窗口内时间切片中最大的请求量
	IsAllow() bool     // 判断是否允许请求
	Avg() float64      // 求窗口内均值
	increment()        // 时间切片自增
	localIndex() int64 // 获取当前位置
	cleanTimeSlice()   // 清空队列中窗口以外的计数
}

type SlideOption func(s *slideWindow)

func WindowSize(ws int64) SlideOption {
	return func(s *slideWindow) {
		s.windowSize = ws
	}
}

func TimeMillisPerSlice(tmps int64) SlideOption {
	return func(s *slideWindow) {
		s.timeMillisPerSlice = tmps
	}
}

func MaxCnt(mc int64) SlideOption {
	return func(s *slideWindow) {
		s.maxCnt = mc
	}
}

type slideWindow struct {
	window             []int64 // 记录切片内的请求数量
	windowSize         int64   // 窗口大小
	timeMillisPerSlice int64   // 每个时间切片大小
	timeSliceSize      int64   // 队列总长度
	maxCnt             int64   // 允许的最大请求数量
	index              int64   // 当前位置
}

func NewSlideWindow(opts ...SlideOption) SlideWindow {
	s := &slideWindow{
		windowSize:         10,
		timeMillisPerSlice: 10,
		maxCnt:             10,
	}
	for _, o := range opts {
		o(s)
	}

	// 保证总长度超过窗口
	s.timeSliceSize = s.windowSize*2 + 1

	// 初始化窗口内请求数
	s.window = make([]int64, s.timeSliceSize)

	go func() {
		for {
			s.cleanTimeSlice()
			time.Sleep(time.Duration(s.timeMillisPerSlice) * time.Millisecond)
		}
	}()

	return s
}

func (s *slideWindow) cleanTimeSlice() {
	// 获取当前位置
	s.index = s.localIndex()

	var i int64
	for i = 1; i <= s.timeSliceSize-s.windowSize; i++ {
		s.window[(s.index+i+s.timeSliceSize)%s.timeSliceSize] = 0
	}

	fmt.Println(s.window)
}

func (s *slideWindow) Total() int64 {
	var sum, i int64
	sum = 0
	for i = 0; i < s.windowSize; i++ {
		sum += s.window[(s.index-i+s.timeSliceSize)%s.timeSliceSize]
	}

	return sum
}

func (s *slideWindow) Max() int64 {
	var max, i int64
	max = 0
	for i = 0; i < s.windowSize; i++ {
		cnt := s.window[(s.index-i+s.timeSliceSize)%s.timeSliceSize]
		if max < cnt {
			max = cnt
		}
	}

	return max
}

func (s *slideWindow) IsAllow() bool {
	// 获取当前位置
	s.index = s.localIndex()

	// 获取窗口总请求数
	sum := s.Total()

	// 阈值判断，若未超限则自增
	if sum < s.maxCnt {
		s.increment()

		return true
	}

	return false
}

func (s *slideWindow) Avg() float64 {
	return float64(s.Total()) / float64(s.windowSize)
}

func (s *slideWindow) increment() {
	// 自增
	atomic.AddInt64(&s.window[s.index], 1)
}

func (s *slideWindow) localIndex() int64 {
	milliTime := time.Now().UnixNano() / 1e6
	return (milliTime / s.timeMillisPerSlice) % s.timeSliceSize
}

func main() {
	sw := NewSlideWindow(WindowSize(5), TimeMillisPerSlice(1000), MaxCnt(300))

	for _, val := range []int64{1, 2, 13, 44, 99, 103, 499, 16} {
		val := val
		var j int64
		for j = 0; j < val; j++ {
			go func() {
				sw.IsAllow()
			}()
		}

		time.Sleep(1 * time.Second)
	}

	time.Sleep(10 * time.Second)
}
