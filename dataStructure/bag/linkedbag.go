package bag

// Bag implemented using a singly linked list.
type Bag struct {
	first *node
	n     int
}

type node struct {
	item interface{}
	next *node
}

// New initialize a Bag.
func New() *Bag {
	return &Bag{}
}

// Add adds the item to this bag.
func (b *Bag) Add(element interface{}) {
	var oldfirst *node = b.first
	b.first = &node{}
	b.first.item = element
	b.first.next = oldfirst
	b.n++
}

// IsEmpty check whether the bag is empty.
func (b *Bag) IsEmpty() bool {
	return b.first == nil
}

// Size returns the number of items in this bag.
func (b *Bag) Size() int {
	return b.n
}

// Values return all of the values in the Bag as a slice.
func (b *Bag) Values() (v []interface{}) {
	for x := b.first; x != nil; x = x.next {
		v = append(v, x.item)
	}
	return v
}
