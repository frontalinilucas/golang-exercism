package tree

import (
	"errors"
	"sort"
)

type (
	Record struct {
		ID, Parent int
	}
	Node struct {
		ID       int
		Children []*Node
	}
)

var (
	ErrNoRootNode        = errors.New("no root node")
	ErrRootNodeHasParent = errors.New("root node has parent")
	ErrNonContinuous     = errors.New("non continuous")
	ErrDuplicateNode     = errors.New("duplicate node")
)

func validate(records []Record) (*Node, bool, error) {
	if len(records) == 0 {
		return nil, false, nil
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	if records[0].ID != 0 {
		return nil, false, ErrNoRootNode
	}
	if records[0].Parent != 0 {
		return nil, false, ErrRootNodeHasParent
	}
	if len(records) == 1 {
		return &Node{ID: records[0].ID}, false, nil
	}
	length := len(records) - 1
	if records[length].ID != length {
		return nil, false, ErrNonContinuous
	}
	return nil, true, nil
}

func Build(records []Record) (*Node, error) {
	node, valid, err := validate(records)
	if !valid {
		return node, err
	}
	nodes := map[int]*Node{
		records[0].ID: {ID: records[0].ID},
	}
	for i := 1; i < len(records); i++ {
		r := records[i]
		parent, ok := nodes[r.Parent]
		if !ok {
			parent = &Node{ID: r.Parent}
			nodes[r.Parent] = parent
		}
		children, ok := nodes[r.ID]
		if !ok {
			children = &Node{ID: r.ID}
			nodes[r.ID] = children
		} else {
			return nil, ErrDuplicateNode
		}
		if parent.ID != children.ID {
			parent.Children = append(parent.Children, children)
		}
	}

	return nodes[0], nil
}
