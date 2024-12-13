package main

import "fmt"

/*
*type ServerState int：
这行代码定义了一个新的类型 ServerState，它是 int 类型的别名。这意味着 ServerState 实际上是一个整数类型，但我们可以给它赋予更有意义的名字来表示服务器的不同状态。

const (：
这行代码开始了一个常量组的定义。在Go中，你可以使用圆括号将多个常量组合在一起。

StateIdle ServerState = iota：
这里定义了第一个常量 StateIdle，并且将其类型指定为 ServerState。iota 是Go语言的一个特殊常量生成器，它用于在常量组中自动生成递增的整数值。因此，StateIdle 将被赋予值 0。

StateConnected、StateError、StateRetrying：
接下来的三个常量 StateConnected、StateError 和 StateRetrying 并没有显式地赋值。由于它们紧跟在 iota 后面，并且处于同一个常量组中，Go语言会自动为它们赋予递增的整数值。因此，StateConnected 将被赋予值 1，StateError 被赋予值 2，StateRetrying 被赋予值 3。
*/

type Weekday int

const (
	Sunday    Weekday = iota // Sunday 将是 0
	Monday                   // Monday 将是 1
	Tuesday                  // Tuesday 将是 2
	Wednesday                // Wednesday 将是 3
	Thursday                 // Thursday 将是 4
	Friday                   // Friday 将是 5
	Saturday                 // Saturday 将是 6
)

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)

	fmt.Println(ServerState.String(1))
	fmt.Println(StateRetrying.String())
	fmt.Println(ServerState.String(StateRetrying))
	fmt.Println(Sunday)
	fmt.Println(StateRetrying)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:

		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
