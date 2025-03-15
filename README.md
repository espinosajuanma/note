# Note - Zettelkasten Quick App

> [!IMPORTANT]
> I am re-writing this app [notes](https://github.com/espinosajuanma/notes).
> Soon it will be archived or removed.

Note is a simple and quick tool to create and manage notes.


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
