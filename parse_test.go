package condition_tree

import "testing"

func Test_buildTree(t *testing.T) {
	testString := `project = LTD and (assignee = "高志威" or assignee = "QA_robot")`
	BuildTree(testString)
}
