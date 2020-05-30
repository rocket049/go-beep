package beep

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/gordonklaus/portaudio"
)

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

type BeepPlayer struct {
	sampleRate float64
	once       sync.Once
	param      portaudio.StreamParameters
}

func NewBeepPlayer() (p *BeepPlayer, e error) {
	return &BeepPlayer{}, nil
}

func (p *BeepPlayer) getSrc(freq int) (ch chan float32, err error) {
	var max = int(p.sampleRate / float64(freq) / 2)
	err = nil
	if max == 0 {
		err = fmt.Errorf("freq is too big(>%f)", p.sampleRate)
		return
	}
	ch = make(chan float32, 2)
	go func() {
		defer func() {
			recover()
		}()
		var top float32 = 0.999
		var base float32 = 0.001
		var step float32 = 0.998 / float32(max)
		for {
			for i := 0; i < max; i++ {
				v := base + step*float32(i)
				ch <- v
			}
			for i := 0; i < max; i++ {
				v := top - step*float32(i)
				ch <- v
			}
		}
	}()

	return
}

func (p *BeepPlayer) getSinSrc(freq int) (ch chan float32, err error) {
	var max = int(p.sampleRate / float64(freq))
	err = nil
	if max == 0 {
		err = fmt.Errorf("freq is too big(>%f)", p.sampleRate)
		return
	}
	ch = make(chan float32, 2)
	go func() {
		defer func() {
			recover()
		}()
		var step float64 = 0.9999 / float64(max)
		for {
			for i := 0; i < max; i++ {
				v := 0.7 * math.Sin(math.Pi*float64(i)*step)
				ch <- float32(v)
			}
		}
	}()

	return
}

func (p *BeepPlayer) Close() {
	portaudio.Terminate()
}

//freq 频率
//delay 毫秒
func (p *BeepPlayer) Beep(freq, delay int) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = r.(error)
		}
	}()
	p.once.Do(func() {
		portaudio.Initialize()
		h, err := portaudio.DefaultHostApi()
		chk(err)
		p.param = portaudio.HighLatencyParameters(nil, h.DefaultOutputDevice)
		p.param.Output.Channels = 1
		p.sampleRate = p.param.SampleRate

	})

	ch, err := p.getSinSrc(freq)
	chk(err)
	defer close(ch)
	stream, err := portaudio.OpenStream(p.param, func(out []float32) {
		for i := range out {
			out[i] = <-ch
		}
	})
	chk(err)
	chk(stream.Start())
	time.Sleep(time.Millisecond * time.Duration(delay))
	chk(stream.Stop())
	return nil
}
