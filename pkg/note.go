package note

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/rwxrob/fs"
	"github.com/rwxrob/fs/dir"
	"github.com/rwxrob/fs/file"
	uniq "github.com/rwxrob/uniq/pkg"
)

const FILE string = "README.md"

type Note struct {
	Id    string
	Path  string
	Title string
}

func New(title string) (*Note, error) {
	id := uniq.Isosec()
	n := &Note{
		Id:    id,
		Title: title,
		Path:  id + "/" + FILE,
	}
	return n, nil
}

func (n *Note) Init() error {
	dir.Create(n.Id)
	f, err := os.Create(n.Path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString("# " + n.Title + "\n")
	return nil
}

func (n *Note) Print() { fmt.Printf("[%s] %s\n", n.Id, n.Title) }

func (n *Note) Edit() { file.Edit(n.Path) }

func (n *Note) Remove() error { return os.RemoveAll(n.Id) }

func List() ([]*Note, error) {
	list := dir.Entries(".")
	notes := []*Note{}
	for _, v := range list {
		if IsValid(v) && fs.IsDir(v) && fs.Exists(v+"/"+FILE) {
			t, _ := GetTitle(v)
			notes = append(notes, &Note{Id: v, Title: t, Path: v + "/" + FILE})
		}
	}
	if len(notes) == 0 {
		return notes, fmt.Errorf("There is no notes here")
	}
	return notes, nil
}

func IsValid(id string) bool {
	return regexp.MustCompile(`\d{14}`).MatchString(id)
}

func GetTitle(id string) (string, error) {
	f, err := os.Open(id + "/" + FILE)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	return strings.Trim(scanner.Text(), "# "), nil
}

func Latest() (*Note, error) {
	list, err := List()
	if err != nil {
		return new(Note), err
	}
	return list[len(list)-1], nil
}

func GetById(id string) (*Note, error) {
	if IsValid(id) && fs.IsDir(id) && fs.Exists(id+"/"+FILE) {
		t, _ := GetTitle(id)
		return &Note{Id: id, Title: t, Path: id + "/" + FILE}, nil
	}
	return new(Note), fmt.Errorf("[%s] is not a valid note", id)
}
