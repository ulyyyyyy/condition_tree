package condition_tree

import "errors"

type queue []string

func (q *queue) isEmpty() (err error) {
	if len(*q) == 0 {
		return errors.New("queue size is 0")
	}
	return nil
}

func (q *queue) Push(str string) {
	*q = append(*q, str)
}

func (q *queue) Pop() (str string, err error) {
	err = q.isEmpty()
	if err != nil {
		return "", err
	}
	index := 0
	element := (*q)[index]
	*q = (*q)[1:]
	return element, nil
}

func (q *queue) getLast() (str string, err error) {
	if len(*q) == 0 {
		return "", errors.New("queue size is 0")
	}
	index := len(*q)
	return (*q)[index-1], nil
}
