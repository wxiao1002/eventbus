package pkg_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	. "github.com/wxiao1002/eventbus/pkg"
)

const (
	create = "create"
	update = "update"
	delete = "delete"
)

type EventData struct {
	data int64
}

func InitResource() {
	f  := func(data any) {
		fmt.Print(data)
	}
	SubScribe(create, f)
	SubScribe(create, func(data any) {

		fmt.Print("123")
	})

	SubScribe(delete, func(data any) {
		eventDeleteData := data.(EventData)
		deleteResource(eventDeleteData.data)
	})
}

func TestEventBus(t *testing.T) {
	InitResource()

	var wg sync.WaitGroup

	wg.Add(4)
	for i := 0; i < 4; i++ {
		i64 := int64(i)
		go func() {
			PubLish(create, EventData{
				data: i64,
			})
			wg.Done()
		}()
	}

	wg.Wait()
	time.Sleep(time.Second)

	fmt.Println("========================================")

	wg.Add(3)
	for i := 0; i < 3; i++ {
		i64 := int64(i)
		go func() {
			PubLish(update, EventData{
				data: i64,
			})
			wg.Done()
		}()
	}

	wg.Wait()
	time.Sleep(time.Second)

	fmt.Println("========================================")

	wg.Add(2)
	for i := 0; i < 2; i++ {
		i64 := int64(i)
		go func() {
			PubLish(delete, EventData{
				data: i64,
			})
			wg.Done()
		}()
	}

	wg.Wait()
	time.Sleep(time.Second)
}

func createResource(id int64) {
	fmt.Printf("Create Resource, id = %v\n", id)
}

func deleteResource(id int64) {
	fmt.Printf("Delete Resource, id = %v\n", id)
}
