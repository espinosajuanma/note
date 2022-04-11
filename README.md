# 🌳 Note - Zettelkasten Quick App

> Inspired in [rwxrob's zet repo](https://github.com/rwxrob/zet), his
> zet bash commands and new [bonzai](https://github.com/rwxrob/bonzai)
> project.

Note is a simple and quick tool to create and manage notes.

[![GoDoc](https://godoc.org/github.com/espinosajuanma/note?status.svg)](https://godoc.org/github.com/espinosajuanma/note)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

## Install

This command can be installed as a standalone program or composed into a
Bonzai command tree.

Standalone

```
go install github.com/espinosajuanma/note/note@latest
```

Composed

```go
package z

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/espinosajuanma/note"
)

var Cmd = &Z.Cmd{
	Name:     `z`,
	Commands: []*Z.Cmd{help.Cmd, note.Cmd},
}
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C note note
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.

## To do

- [ ] Config `zettelkasten` path
- [ ] Allow multicall
- [ ] Add `search`/`query` command
- [ ] Document everything
- [ ] Add tests
