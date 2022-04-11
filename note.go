package note

import (
	"strings"

	note "github.com/espinosajuanma/note/pkg"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/config"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
)

var Cmd = &Z.Cmd{
	Name:     `note`,
	Summary:  `taking quick notes`,
	Version:  `v0.0.1`,
	Commands: []*Z.Cmd{help.Cmd, config.Cmd, List, New, Edit, Remove, Latest, Push},
}

var New = &Z.Cmd{
	Name:     `new`,
	Aliases:  []string{"add"},
	Summary:  `add a new note to the current directory`,
	Usage:    `[title]`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		title := "Title's note"
		if len(args) != 0 {
			title = strings.Join(args, " ")
		}
		n, err := note.New(title)
		if err != nil {
			return err
		}
		n.Init()
		n.Edit()
		return nil
	},
}

var List = &Z.Cmd{
	Name:     `list`,
	Aliases:  []string{"ls"},
	Summary:  `list all valid notes in the current directory`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		list, err := note.List()
		if err != nil {
			return err
		}
		for _, v := range list {
			v.Print()
		}
		return nil
	},
}

var Edit = &Z.Cmd{
	Name:     `edit`,
	Aliases:  []string{"e", "ed"},
	Usage:    `<id>`,
	Summary:  `open a specific note in your editor`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		n, err := note.Latest()
		if len(args) != 0 {
			n, err = note.GetById(args[0])
		}
		if err != nil {
			return err
		}
		n.Edit()
		return nil
	},
}

var Push = &Z.Cmd{
	Name:     `push`,
	Usage:    `<id>`,
	Aliases:  []string{"commit"},
	Summary:  `commit and push to git`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		n, err := note.Latest()
		if len(args) != 0 {
			n, err = note.GetById(args[0])
		}
		if err != nil {
			return err
		}
		Z.Exec("git", "add", n.Path)
		Z.Exec("git", "commit", "-m", "'"+n.Title+"'")
		Z.Exec("git", "push")
		return nil
	},
}

var Remove = &Z.Cmd{
	Name:     `remove`,
	Aliases:  []string{"rm"},
	Usage:    `<id>`,
	Summary:  `remove a specific note`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		n, err := note.Latest()
		if len(args) != 0 {
			n, err = note.GetById(args[0])
		}
		if err != nil {
			return err
		}
		n.Print()
		if yesOrNo("Are you sure you want to remove this note?") {
			err := n.Remove()
			if err != nil {
				return err
			}
		}
		return nil
	},
}

var Latest = &Z.Cmd{
	Name:     `latest`,
	Aliases:  []string{"last"},
	Summary:  `prints latest note`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		n, err := note.Latest()
		if err != nil {
			return err
		}
		n.Print()
		return nil
	},
}

func yesOrNo(q string) bool {
	y := term.Prompt(q + " (y/N) ")
	return y == "Y" || y == "y" || y == "yes"
}
