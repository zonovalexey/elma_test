package solutions

import (
	"encoding/json"
	"fmt"
	"log"
	"tasks/connect"
)

type task struct {
	data  []int
	shift int
}

func Task(task_name []string, num int, user_name string, c chan int) {
	data := connect.Connect(task_name[num])
	var res = make(map[string]interface{}, len(data)+1)
	var tmp = make(map[string][]interface{}, 2)
	res["user_name"] = user_name
	res["task"] = task_name[num]
	for _, v := range data {

		var t task
		var tmp_payload []interface{}

		switch task_name[num] {
		case task_name[0]:
			{
				for _, d := range v.([]interface{}) {
					switch vv := d.(type) {
					case float64:
						t.shift = int(vv)
					case []interface{}:
						for _, k := range d.([]interface{}) {
							t.data = append(t.data, int(k.(float64)))
						}
					}
				}
			}
		case task_name[1], task_name[2], task_name[3]:
			{
				for _, d := range v.([]interface{}) {
					for _, k := range d.([]interface{}) {
						t.data = append(t.data, int(k.(float64)))
					}
				}
			}
		}

		switch task_name[num] {
		case task_name[0]:
			{
				tmp_payload = append(tmp_payload, t.data, t.shift)
				sol := solutionTask1(t.data, t.shift)
				tmp["payload"] = append(tmp["payload"], tmp_payload)
				tmp["results"] = append(tmp["results"], sol)
			}
		case task_name[1]:
			{
				tmp_payload = append(tmp_payload, t.data)
				sol := solutionTask2(t.data)
				tmp["payload"] = append(tmp["payload"], tmp_payload)
				tmp["results"] = append(tmp["results"], sol)
			}
		case task_name[2]:
			{
				tmp_payload = append(tmp_payload, t.data)
				sol := solutionTask3(t.data)
				tmp["payload"] = append(tmp["payload"], tmp_payload)
				tmp["results"] = append(tmp["results"], sol)
			}
		case task_name[3]:
			{
				tmp_payload = append(tmp_payload, t.data)
				sol := solutionTask4(t.data)
				tmp["payload"] = append(tmp["payload"], tmp_payload)
				tmp["results"] = append(tmp["results"], sol)
			}

		}

	}

	res["results"] = tmp

	b, err := json.Marshal(res)

	if err != nil {
		fmt.Println("Ошибка преобразования результатов выполнения задания \"", task_name[num], "\" перед передачей данных на сервер")
		log.Fatal(err)
	}

	res_server := connect.Send(b, task_name[num])

	if int(res_server.Percent) == 100 {
		fmt.Println("Задача \""+task_name[num]+"\": решено", fmt.Sprint(res_server.Percent)+"%")
	} else {
		fmt.Println("Задача \""+task_name[num]+"\": решено", fmt.Sprint(res_server.Percent)+"%\nИнформация об ошибке:")
		fmt.Println("-------------------------------------------")
		for _, k := range res_server.Fails {
			fmt.Println("Номер блока данных:", k.DataSet)
			fmt.Println("Отправленный результат:", k.ExternalResult)
			fmt.Println("Результат сервера:", k.OriginalResult)
			fmt.Println("-------------------------------------------")
		}
	}

	c <- num
}
