package handling

// Parts represents the pieces of a file path.
type Parts []string

// NighAll provides all but the last element.
func (p Parts) NighAll() Parts {
	if p.IsEmpty() {
		return Parts{}
	} else {
		return p[:p.Len()-1]
	}
}

// Last returns the first part or an empty string if the Parts is empty.
func (p Parts) Last() string {
	n := p.Len()
	if n > 0 {
		return p[n-1]
	} else {
		return ""
	}
}

// First returns the first part in the Parts slice.
func (p Parts) First() string {
	if p.Len() > 0 {
		return p[0]
	} else {
		return ""
	}
}

// IsEmpty returns true if Parts is nil or empty.
func (p Parts) IsEmpty() bool {
	return p == nil || len(p) == 0
}

// Len returns the number of elements in the Parts slice.
func (p Parts) Len() int {
	if p == nil {
		return 0
	} else {
		return len(p)
	}
}
