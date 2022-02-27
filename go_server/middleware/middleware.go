package middleware
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"fmt" 
	"github.com/gorilla/mux"
	"os"
	"io/ioutil"
	"bytes"
	"os/exec"
	"time"
)

type QuestionAndAnswer struct {
	Context string `json:"context"`
	Question string `json:"question"`
}

func init() {
}

func GetQuestionAndAnswerResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)

	currentTime := time.Now()
	filename = currentTime.Format("2006-01-02T15:04:05") + ".json"
	saveQuestionAndAnswerQuestion(filename, params["inputContext"], params["inputQuestion"])
	payload := getQuestionAndAnswerResult(filename)
	json.NewEncoder(w).Encode(payload)
}

func saveQuestionAndAnswerQuestion(filename string, inputContext string, inputQuestion string) {
	data := &QuestionAndAnswer{
		Context: inputContext,
		Question: inputQuestion,
	}

	buf, err1 := json.Marshal(data)
	if err1 !=nil {
		panic(err1)
	}
	inputFilePath := "/home/ubuntu/server_inputs/" + filename
	err2 := ioutil.WriteFile(inputFilePath, buf, 0644)
	if err2 !=nil {
		panic(err2)
	}
}

func runAndGetQuestionAndAnswerResult(filename string) (string) {
	cmd := exec.Command("python3", "/home/ubuntu/gogogo/python_backend/run_bert_answers_questions.py", "-f", filename, "-o", filename)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err1 := cmd.Run()
	if err1 != nil {
		fmt.Println(fmt.Sprint(err1) + ": " + stderr.String())
		return ""
	}

	outputFilePath := "/home/ubuntu/server_outputs/" + filename
	jsonFile, err2 := os.Open(outputFilePath)
    // if we os.Open returns an error then handle it
    if err2 != nil {
        fmt.Println(err2)
		return ""
    }
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    var answers map[string]string
    json.Unmarshal([]byte(byteValue), &answers)
	return answers["1"]
}