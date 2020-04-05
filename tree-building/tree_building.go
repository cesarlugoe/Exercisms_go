package tree

import "fmt"

// Record ...
type Record struct {
	ID     int
	Parent int
}

// Node ...
type Node struct {
	ID       int
	Children []*Node
}

// Build builds an ordered tree
func Build(records []Record) (*Node, error) {
	rl := len(records)
	if rl < 1 {
		return nil, nil
	}

	recordsOrder := make([]int, rl)
	for i, r := range records {
		if r.ID < 0 || r.ID >= rl {
			return nil, fmt.Errorf("invalid record ID %d", r.ID)
		}
		recordsOrder[r.ID] = i
	}

	nodes := make([]Node, rl)
	for i := range recordsOrder {
		r := records[recordsOrder[i]]

		if !((r.ID > r.Parent) || (r.ID == 0 && r.Parent == 0)) {
			return nil, fmt.Errorf("Invalid Parent on record %d", r.ID)
		}

		if r.ID != i {
			return nil, fmt.Errorf("IDs are not secuential")
		}

		nodes[i].ID = i
		if nodes[i].ID != 0 {
			p := &nodes[r.Parent]
			p.Children = append(p.Children, &nodes[i])
		}
	}

	return &nodes[0], nil
}
