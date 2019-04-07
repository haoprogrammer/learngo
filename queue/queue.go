package queue

//interface{}指示任意类型
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)

}

//弹出
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	(*q) = (*q)[1:]
	return head
}

//只限定int
func (q *Queue) PopInt() int {
	head := (*q)[0]
	(*q) = (*q)[1:]
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
