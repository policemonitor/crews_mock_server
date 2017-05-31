package main

import (
	"fmt"
	"crews_mock_server/utils"
	"sync"
	"time"
	"strconv"
)

func main() {
	var crews = utils.ReadCrewsConfigs()

	fmt.Println("LOADED CREWS:" + strconv.Itoa(len(crews)))

	var wg sync.WaitGroup
	wg.Add(len(crews))

	for i := range crews  {
		var crew = crews[i]

		go func() {
			defer wg.Done()
			for true {
				fmt.Println("moved "+crew.VinNumber)

				crew.Move()
				utils.SendCoordinates(crew)

				time.Sleep(time.Second * 2)
			}
		} ()
	}
	fmt.Println("EMULATING STARTED")
	wg.Wait()
	fmt.Println("ALL ROUTINES HAVE BEEN STOPED")
}
