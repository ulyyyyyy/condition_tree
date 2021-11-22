package condition_tree

import (
	"log"
	"testing"
)

func TestBuildTree(t *testing.T) {
	testString := `project = LTD and (assignee = "高志威" or assignee = "QA_robot")`
	BuildTree(testString)
}

func Test_parseCondition(t *testing.T) {
	condition1, err := parseCondition("project = LTD")
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(condition1)
	condition2, err := parseCondition(`issuetype in (故障, 缺陷, bug)`)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(condition2)
	condition3, err := parseCondition(`issuetype in`)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(condition3)
}
