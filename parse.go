package condition_tree

import "strings"

// Parse string to conditionTree
func Parse(str string) (tree ConditionTree, err error) {

	return ConditionTree{}, err
}

// BuildTree test
func BuildTree(str string) {
	var logicStack stack

	var opStack stack
	var conditionQueue queue

	str = strings.ToUpper(str)
	str = strings.ReplaceAll(str, ", ", ",")
	for _, ele := range strings.Split(str, " ") {
		if ele == "AND" || ele == "OR" {
			logicStack.Push(ele)
		} else if checkOp(ele) {
			opStack.Push(ele)
		} else {
			last, _ := conditionQueue.getLast()
			if checkOp(last) {
				f, _ := conditionQueue.Pop()
				newCondition(f, last, ele)
			}
		}
	}
}
