package tasks

import (
	"testing"
	"fmt"
	"time"
)

func TestNewTask(t *testing.T) {
	fmt.Println("==================================TestNewTask=================================")
	handler := func() string {
		fmt.Println("aaaaaaa");
		return "完成了";
	}
	task := NewTask(handler);
	task.Run();
	task.Wait();

	fmt.Println(task.Result)
}

func TestWaitAll(t *testing.T) {
	fmt.Println("==================================TestWaitAll=================================")
	handler1 := func(){
		fmt.Println(time.Now())
		fmt.Println("handler1","我在等待指定的时间后执行");
	}
	param2 := "aaaaaaaaaaaaaaaaaaaaaa";
	handler2 := func(p string) {
		fmt.Println("handler2",time.Now())
		fmt.Println(p);
	}
	param3 := "bbbbbbbbbbbbbbbbbbbbbbbbbb";
	handler3 := func(p string) string {
		fmt.Println(p);
		return p+"111111111111111";
	}

	task1 := NewTask(handler1).ContinueWith(func(result TaskResult){
		fmt.Println("我在task1执行后执行。");
	}).ContinueWith(func(result TaskResult){
		fmt.Println("我在task1执行后第二次执行。");
	}).Delay(5*time.Second).Run();
	task2 := NewTask(handler2,param2).Run();
	task3 := NewTask(handler3,param3).Run();

	WaitAll(task1,task2,task3);
	fmt.Println(task3.Result)
}