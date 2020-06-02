package beep

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
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
	onceBeep   sync.Once
	ch         chan float32
	max        int32
}

func NewBeepPlayer() (p *BeepPlayer, e error) {
	return &BeepPlayer{}, nil
}

func (p *BeepPlayer) getSinSrc(freq int) (ch chan float32, err error) {
	var max = int(p.sampleRate / float64(freq))
	err = nil
	if max == 0 {
		err = fmt.Errorf("freq is too big(>%f)", p.sampleRate)
		return
	}
	atomic.StoreInt32(&p.max, int32(max))
	p.onceBeep.Do(func() {
		p.ch = make(chan float32, 2)
		go func() {
			defer func() {
				recover()
			}()

			for {
				m := int(atomic.LoadInt32(&p.max))
				var step float64 = 0.9999 / float64(m)
				for i := 0; i < m; i++ {
					v := 0.9*0.5*math.Sin(2*math.Pi*float64(i)*step) + 0.5
					p.ch <- float32(v)
				}
			}
		}()
	})
	ch = p.ch
	err = nil
	return
}

func (p *BeepPlayer) Close() {
	defer func() {
		recover()
	}()
	portaudio.Terminate()
	close(p.ch)
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
