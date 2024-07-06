package errd

import (
	"fmt"
	"os"

	"github.com/NuruProgramming/Nuru/token"
)

type MakosaSintaksia struct {
	Ujumbe       string
	Muktadha string
	Info token.Token
}

func (s *MakosaSintaksia) Kosa() string {
	return fmt.Sprintf("Kosa la kisintaksia: %s:%d:%d\n%s\n\n%s",
		s.Info.Filename, s.Info.Line.Start.Line, s.Info.Line.Start.Column, s.Ujumbe, s.Muktadha)
}

func (s *MakosaSintaksia) Onyesha() {
	fmt.Fprintf(os.Stderr, "%s\n", s.Kosa())
}

func (s *MakosaSintaksia) Hatari() {
	s.Onyesha()
	os.Exit(1)
}
