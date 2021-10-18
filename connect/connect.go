package connect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Connect(task string) []interface{} {

	Tasks_path := "http://116.203.203.76:3000/tasks/" + strings.ReplaceAll(task, " ", "%20")

	resp, err := http.Get(Tasks_path)

	if err != nil {
		fmt.Println("Ошибка соединения с сервером при получении задания \"" + task + "\"")
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Ошибка чтения задания \"" + task + "\" с сервер")
		log.Fatal(err)
	}

	var f interface{}
	err = json.Unmarshal(body, &f)

	if err != nil {
		fmt.Println("Ошибка преобразования данных задания \"" + task + "\" в JSON формат")
		log.Fatal(err)
	}

	m := f.([]interface{})

	return m
}

func Send(b []byte, task_name string) map[string]interface{} {
	r := bytes.NewReader([]byte(b))
	resp, err := http.Post("http://116.203.203.76:3000/tasks/solution", "application/json", r)

	// resp, err := http.Post("https://enk17fim6reyhvz.m.pipedream.net", "application/json", r)

	if err != nil {
		fmt.Println("Ошибка передачи результатов выполнения задания \"" + task_name + "\" на сервер")
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Ошибка чтения ответа на решение задания \"" + task_name + "\" с сервер")
		log.Fatal(err)
	}

	var f interface{}
	err = json.Unmarshal(body, &f)

	if err != nil {
		fmt.Println("Ошибка преобразования результатов решения задачи \"" + task_name + "\" в формат JSON")
		log.Fatal(err)
	}

	m := f.(map[string]interface{})

	return m
}
