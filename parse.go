package condition_tree

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

var (
	condition_field_list = []string{"field", "op", "value"}
)

// Parse string to conditionTree
func Parse(str string) (tree ConditionTree, err error) {

	return ConditionTree{}, err
}

// BuildTree test
func BuildTree(str string) {
	//var logicStack stack
	//var conditionQueue conditionQueue

	var strList []string
	str = strings.ToUpper(str)
	split := strings.Split(str, " AND ")
	for _, s := range split {
		i := strings.Split(s, " OR ")
		for _, s2 := range i {
			strList = append(strList, s2)
		}
	}
	log.Println(strList)
	for _, str := range strList {
		if str[0] != '(' && str[len(str)-1] != ')' {

		}
	}
}

// parseCondition
func parseCondition(str string) (condition, error) {
	var opString string
	var field string
	var index int

	str = strings.ReplaceAll(str, ", ", ",")
	for i, ele := range strings.Split(str, " ") {
		index = i
		if checkOp(ele) {
			opString = ele
		} else if opString != "" {
			if strings.Contains(ele, "(") || strings.Contains(ele, ")") {
				ele = ele[1 : len(ele)-1]
			}
			return *newCondition(field, opString, ele), nil
		} else {
			field = ele
		}
	}
	return condition{}, errors.New(fmt.Sprintf("parse condition error. [%s] mustn't be nil", condition_field_list[index+1]))
}
