package printer

import "fmt"

type Datastore interface {
	ListAll(traversalType string)
	Insert(data int64)
}

type Printer struct {
	ds Datastore
}

func New(ds Datastore) *Printer {
	return &Printer{ds: ds}
}

func (p *Printer) BFS() {
	fmt.Printf("\n Printing BFS \n")
	p.ds.ListAll("bfs")
	fmt.Printf("\n Done Printing BFS \n")
}

func (p *Printer) DFS() {
	fmt.Printf("\n Printing DFS \n")
	p.ds.ListAll("dfs")
	fmt.Printf("\n Done Printing DFS \n")
}

func (p *Printer) Build() {

	p.ds.Insert(50)
	p.ds.Insert(45)
	p.ds.Insert(12)
	p.ds.Insert(5)
	p.ds.Insert(33)
	p.ds.Insert(25)
	p.ds.Insert(55)
	p.ds.Insert(100)
	p.ds.Insert(98)
	p.ds.Insert(72)
	p.ds.Insert(61)

}
