// Code generated DO NOT EDIT

package cmds

import "testing"

func hyperloglog0(s Builder) {
	s.Pfadd().Key("1").Element("1").Element("1").Build()
	s.Pfadd().Key("1").Build()
	s.Pfcount().Key("1").Key("1").Build()
	s.Pfmerge().Destkey("1").Sourcekey("1").Sourcekey("1").Build()
	s.Pfmerge().Destkey("1").Build()
}

func TestCommand_InitSlot_hyperloglog(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { hyperloglog0(s) })
}

func TestCommand_NoSlot_hyperloglog(t *testing.T) {
	var s = NewBuilder(NoSlot)
	t.Run("0", func(t *testing.T) { hyperloglog0(s) })
}
