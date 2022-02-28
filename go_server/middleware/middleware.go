package middleware

import (
	"encoding/json"
	"net/http"
	"fmt" 
	"os"
	"io/ioutil"
	"bytes"
	"os/exec"
	"hash/fnv"
)

type QuestionAndAnswer struct {
	Context string `json:"context"`
	Question string `json:"question"`
}

func init() {
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func GetQuestionAndAnswerResult(w http.ResponseWriter, r *http.Request) {
	var data QuestionAndAnswer
	err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	
	filename := fmt.Sprint(hash(data.Context + data.Question)) + ".json"
	saveQuestionAndAnswerQuestion(filename, data)
	payload := runAndGetQuestionAndAnswerResult(filename)
	json.NewEncoder(w).Encode(payload)
}

func saveQuestionAndAnswerQuestion(filename string, data QuestionAndAnswer) {
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