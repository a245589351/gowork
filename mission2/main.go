package main

import "gowork/mission2/core"

func main() {
	if err := core.Run(); err != nil {
		panic(err)
	}
}
