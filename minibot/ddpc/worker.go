package main

import (
	"fmt"
	"time"

	// "github.com/docker/go-connections/nat"
	"github.com/cjkao/Rocket.Chat.Go.SDK/models"
	rt "github.com/cjkao/Rocket.Chat.Go.SDK/realtime"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type State int

const (
	Pending State = iota
	Scheduled
	Completed
	Running
	Failed
)

type Task struct {
	ID            uuid.UUID
	Name          string
	State         State
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}
type TaskEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}
type Worker struct {
	me       *models.User
	recPipe  <-chan map[string]interface{}
	pushChan chan *models.Message
	cli      *rt.Client
	// Queue     queue.Queue
	Db        map[uuid.UUID]Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("I will collect stats")
}

func (w *Worker) RunTask() {
	fmt.Println("I will start or stop a task")
}

const MAX_MSG = 1000

func (w *Worker) StartTask(msgChan <-chan map[string]interface{}, done <-chan interface{}) {
	msgCnt := 0
	for {
		if msgCnt > MAX_MSG {
			break
		}
		println(msgCnt)
		msgCnt++
		select {
		case <-done:
			return
		case v, ok := <-msgChan:
			PrettyPrint(v)
			if !ok {
				return
			}
			// mdec:=json.Unmarshal(v,)
			// if !v.UnRead {
			// 	print("not new msg")
			// 	continue
			// }
			// if v.RoomID == "" || v.Msg == "" {
			// 	continue
			// }
			// if v.User.ID == w.me.ID {
			// 	println("my self-->mesg")
			// 	continue
			// }

			// msg := rtClient.NewMessage(v.Channel, "i hear you"+v.Msg)
			// msg := new(models.Message)
			println("new msg")
			// switch v.Msg {
			// case "ls":
			// 	msg := w.cli.NewMessage(&models.Channel{ID: v.ID}, "AA_"+v.Msg)
			// 	msg.RoomID = v.RoomID
			// 	w.pushChan <- msg
			// }
			// msg.Msg = "AAAAAAA"
		}
		// var msg models.Message
		// msg, err := <-w.Channel
		// if err {
		// 	break
		// }
	}
	fmt.Println("I will start a task")
}

func (w *Worker) StartPush(done <-chan interface{}) {
	for {
		msg := <-w.pushChan
		m, err := w.cli.SendMessage(msg)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("send out --->%#v\n", m)
	}
}
func (w *Worker) StartUserNotify(notify <-chan string, done <-chan interface{}) {
	msgCnt := 0
	for {
		// if msgCnt > MAX_MSG {
		// 	break
		// }
		println(msgCnt)
		msgCnt++
		select {
		case <-done:
			return
		case v, ok := <-notify:
			fmt.Printf("~~~>notify: %s\n", v)
			if !ok {
				return
			}
		}
	}
	// fmt.Println("I will start a task")
}
func (w *Worker) StopTask() {
	fmt.Println("I will stop a task")
}
