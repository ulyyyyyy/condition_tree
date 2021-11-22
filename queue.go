package condition_tree

import "errors"

type conditionQueue []condition

func (q *conditionQueue) isEmpty() (err error) {
	if len(*q) == 0 {
		return errors.New("conditionQueue size is 0")
	}
	return nil
}

func (q *conditionQueue) Push(con condition) {
	*q = append(*q, con)
}

func (q *conditionQueue) Pop() (con condition, err error) {
	err = q.isEmpty()
	if err != nil {
		return condition{}, err
	}
	index := 0
	element := (*q)[index]
	*q = (*q)[1:]
	return element, nil
}

func (q *conditionQueue) getLast() (con condition, err error) {
	if len(*q) == 0 {
		return condition{}, errors.New("conditionQueue size is 0")
	}
	index := len(*q)
	return (*q)[index-1], nil
}
