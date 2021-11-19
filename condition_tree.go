package condition_tree

// ConditionTree it's a complete binary tree
type ConditionTree struct {
	Val   condition
	Left  *ConditionTree
	Right *ConditionTree
}

// a condition is used to check whether it's existed in the map.
type condition struct {
	Logic string

	F  string
	OP string
	V  string
}

func newCondition(f string, OP string, v string) *condition {
	return &condition{F: f, OP: OP, V: v}
}

// ToString convert the tree to String.
func (tree ConditionTree) ToString() (ret string, err error) {

	return "", nil
}

// GetResult Get the final boolean result from data map via tree.
func (tree ConditionTree) GetResult(json map[string]string) (result bool, err error) {

	return false, err
}
