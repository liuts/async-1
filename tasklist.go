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
//将一个任务添加到任务列表中
func (tlist *TaskList)Add(task *Task) *TaskList {
	tlist.list.PushFront(task);
	return tlist;
}
//批量添加任务到任务列表中
func (tlist *TaskList)AddRange(tasks ... *Task) *TaskList {
	for _,task := range tasks{
		tlist.list.PushFront(task);
	}
	return tlist;
}
//运行任务列表中的所有任务
func (tlist *TaskList)Run() *TaskList{
	fmt.Println(tlist.list.Len())
	for element := tlist.list.Front(); element != nil; element = element.Next() {
		if task ,ok:= element.Value.(*Task);ok && !task.IsCompleted{
			task.Run();
		}
	}
	return tlist;
}
//等待所有任务执行完成
func (tlist *TaskList)WaitAll()  {
	for element := tlist.list.Front(); element != nil; element = element.Next() {
		if task ,ok:= element.Value.(*Task);ok && !task.IsCompleted{
			task.wait.Wait();
		}
	}
}