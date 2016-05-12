package tasks

import (
	"testing"
	"fmt"
	"time"
)

func TestTaskList_Add(t *testing.T) {
	fmt.Println("==================================TestTaskList_Add=================================")
	handler1 := func(){
		fmt.Println(time.Now())
		fmt.Println("TestTaskList_Add: handler1");
	}
	param2 := "TestTaskList_Add:aaaaaaaaaaaaaaaaaaaaaa";
	handler2 := func(p string) {
		fmt.Println("TestTaskList_Add:handler2",time.Now())
		fmt.Println(p);
	}
	param3 := "TestTaskList_Add: bbbbbbbbbbbbbbbbbbbbbbbbbb";
	handler3 := func(p string) string {
		fmt.Println(p);
		return p+"111111111111111";
	}

	task1 := NewTask(handler1);
	task2 := NewTask(handler2,param2);
	task3 := NewTask(handler3,param3);

	taskList := NewTaskList();

	taskList.AddRange(task1,task2,task3).Run().WaitAll();
}
