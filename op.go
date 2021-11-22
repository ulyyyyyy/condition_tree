package condition_tree

var (
	opList = []string{"=", "==", "!=", ">", "<", ">=", "<=", "~", "between", "in", "not in"}
	op     map[string]interface{}
)

func checkOp(ele string) bool {
	if len(op) == 0 {
		initOP()
	}
	if _, ok := op[ele]; ok {
		return ok
	}
	return false
}

func initOP() {
	op = make(map[string]interface{})
	for _, o := range opList {
		op[o] = nil
	}
}
