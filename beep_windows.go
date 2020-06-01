package beep

import (
	"golang.org/x/sys/windows"
)

type BeepPlayer struct {
	h    *windows.DLL
	proc *windows.Proc
}

func NewBeepPlayer() (p *BeepPlayer, e error) {
	defer func() {
		if r := recover(); r != nil {
			p = nil
			e = r.(error)
		}
	}()
	h, err := windows.LoadDLL("kernel32.dll")
	chk(err)
	proc, err := h.FindProc("Beep")
	chk(err)
	return &BeepPlayer{h, proc}, nil
}

//freq 频率
//delay 毫秒
func (p *BeepPlayer) Beep(freq, delay int) error {
	_, _, err := p.proc.Call(uintptr(freq), uintptr(delay))
	return err
}

func (p *BeepPlayer) Close() {
	p.h.Release()
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
