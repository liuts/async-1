package tasks

import (
	"container/list"
	"fmt"
)

type TaskList struct {
	list *list.List
}

func NewTaskList()*TaskList{

	return &TaskList{
		list : list.New(),
	}
}
func (tlist *TaskList)Add(task *Task) *TaskList {
	tlist.list.PushFront(task);
	return tlist;
}
func (tlist *TaskList)AddRange(tasks ... *Task) *TaskList {
	for _,task := range tasks{
		tlist.list.PushFront(task);
	}
	return tlist;
}

func (tlist *TaskList)Run() *TaskList{
	fmt.Println(tlist.list.Len())
	for element := tlist.list.Front(); element != nil; element = element.Next() {
		if task ,ok:= element.Value.(*Task);ok && !task.IsCompleted{
			task.Run();
		}
	}
	return tlist;
}

func (tlist *TaskList)WaitAll()  {
	for element := tlist.list.Front(); element != nil; element = element.Next() {
		if task ,ok:= element.Value.(*Task);ok && !task.IsCompleted{
			task.wait.Wait();
		}
	}
}