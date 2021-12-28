package react

// Define reactor, cell and canceler types here.
// These types will implement the Reactor, Cell and Canceler interfaces, respectively.
type (
	cellID int
	cell   struct {
		id        cellID
		value     int
		reactor   *reactor
		callbacks []*func(int)
		computes1 func(int) int
		computes2 func(int, int) int
	}
	canceler struct {
		callback *func(int)
	}
	reactor struct {
		lastCellID cellID
		sons       map[cellID][]*cell
		parents    map[cellID][]*cell
	}
)

func (c *canceler) Cancel() {
	*c.callback = nil
}

func (c *cell) Value() int {
	return c.value
}

func (c *cell) SetValue(value int) {
	c.setValue(value)
	c.reactor.onChange(c)
}

func (c *cell) setValue(value int) {
	if c.value == value {
		return
	}
	c.value = value
	c.runCallbacks(value)
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	c.callbacks = append(c.callbacks, &callback)
	return &canceler{callback: &callback}
}

func (c *cell) runCallbacks(value int) {
	for _, callback := range c.callbacks {
		if *callback != nil {
			(*callback)(value)
		}
	}
}

func New() Reactor {
	return &reactor{
		sons:    map[cellID][]*cell{},
		parents: map[cellID][]*cell{},
	}
}

func (r *reactor) getNextID() cellID {
	r.lastCellID++
	return r.lastCellID
}

func (r *reactor) CreateInput(initial int) InputCell {
	return &cell{id: r.getNextID(), value: initial, reactor: r}
}

func (r *reactor) addParent(son, parent *cell) {
	r.parents[son.id] = append(r.parents[son.id], parent)
}

func (r *reactor) addSons(son, parent *cell) {
	r.addSon(son, parent)
	for i := range r.parents[parent.id] {
		r.addSons(son, r.parents[parent.id][i])
	}
}

func (r *reactor) addSon(son, parent *cell) {
	for i := range r.sons[parent.id] {
		if r.sons[parent.id][i] == son {
			return
		}
	}
	r.sons[parent.id] = append(r.sons[parent.id], son)
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	c := &cell{id: r.getNextID(), value: compute(dep.Value()), reactor: r, computes1: compute}
	parent := dep.(*cell)

	r.addParent(c, parent)
	r.addSons(c, parent)

	return c
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	c := &cell{id: r.getNextID(), value: compute(dep1.Value(), dep2.Value()), reactor: r, computes2: compute}
	parent1 := dep1.(*cell)
	parent2 := dep2.(*cell)

	r.addParent(c, parent1)
	r.addParent(c, parent2)

	r.addSons(c, parent1)
	r.addSons(c, parent2)

	return c
}

func (r *reactor) onChange(cell *cell) {
	for _, c := range r.sons[cell.id] {
		if c.computes1 != nil {
			parent := r.parents[c.id][0]
			c.setValue(c.computes1(parent.value))
		} else if c.computes2 != nil {
			parent1 := r.parents[c.id][0]
			parent2 := r.parents[c.id][1]
			c.setValue(c.computes2(parent1.value, parent2.value))
		}
	}
}
