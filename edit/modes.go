package edit

import (
	"github.com/elves/elvish/edit/completion"
	"github.com/elves/elvish/edit/history"
	"github.com/elves/elvish/edit/lastcmd"
	"github.com/elves/elvish/edit/location"
	"github.com/elves/elvish/eval"
)

func init() {
	atEditorInit(func(ed *editor, ns eval.Ns) {
		location.Init(ed, ns)
		lastcmd.Init(ed, ns)
		history.Init(ed, ns)
		completion.Init(ed, ns)
	})
}
