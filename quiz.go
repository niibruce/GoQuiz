package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)


var filename_flag *string
var timelimit_flag *int
var shuffle_flag *bool

func init(){
	//parse the command line flags
	parseCLFlags()
}

func main(){


	var questions = loadQuestions(*filename_flag, *shuffle_flag)

	//set the timer
	timer := time.NewTimer(time.Duration(*timelimit_flag) * time.Second)

	question_p := &questions //creating an array to the questions struct

	runQuiz(question_p, timer) //start the quiz

	scoreQuiz(question_p) // provide a score of the quiz

}


func checkerror(e error){
	if e != nil{
		panic(e)
	}
}

func printQuestions(q QuestionCollection){
	for _, i := range q.Questions {
		fmt.Println(i.question + " " + strconv.FormatBool(i.isCorrect))
	}
}

func loadQuestions(filename string, shuffle bool) QuestionCollection{
	//read the csv file
	csvFile, err := os.Open(filename)
	checkerror(err)
	defer csvFile.Close()

	lines, err := csv.NewReader(csvFile).ReadAll()
	checkerror(err)

	var questions QuestionCollection

	for _, line := range lines {
		data := Question{
			question: string(line[0]),
			answer: string(line[1]),
			response: "",
			isCorrect: false,
		}
		questions.Questions = append(questions.Questions, data)
	}

	fmt.Println(fmt.Sprintf("%d questions loaded.", len(lines)))

	//Shuffling the questions
	if shuffle == true{
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for n:=len(questions.Questions);n>0; n--{
			randIndex := r.Intn(n)
			questions.Questions[n-1], questions.Questions[randIndex] = questions.Questions[randIndex], questions.Questions[n-1]
		}
	}

	return questions
}

func parseCLFlags(){
	filename_flag = flag.String("csv", "problems_test.csv", "a string")
	timelimit_flag = flag.Int("limit", 5, "the time limit for the quiz in seconds")
	shuffle_flag = flag.Bool("shuffle", false, "Shuffle the deck of questions or not")
	flag.Parse()
}

func runQuiz(q *QuestionCollection, t *time.Timer){
	reader := bufio.NewReader(os.Stdin)
	var response string

	answerCh := make(chan *Question)

	for i := 0; i < len(q.Questions); i++ {
		p := &q.Questions[i]

		fmt.Println(p.question)

		go func() {
			response, _ = reader.ReadString('\n')
			p.response = strings.Trim(response, "\n")
			answerCh <- p
		}()

		select {
		case <- t.C:
			fmt.Println("Time is up!")
			return

		case answer := <-answerCh:
			if answer.response == p.answer {
				p.isCorrect = true
			}
		}

		//response, _ = reader.ReadString('\n')
		//p.response = strings.Trim(response, "\n")

/*
		if p.response == p.answer {
			p.isCorrect = true
		}
*/
	} //end for
}

func scoreQuiz(q *QuestionCollection){
	var correct_answers = 0

	for _, v := range q.Questions{
		if v.isCorrect == true {
			correct_answers = correct_answers + 1
		}
	}

	fmt.Println(fmt.Sprintf("You scored a total of %d correct out of %d questions", correct_answers, len(q.Questions)))
}

func printQuestionCollection(q *QuestionCollection){
	for _, v := range q.Questions {
		fmt.Println(fmt.Sprintf("%s, %s, %s, %b", v.question, v.answer, v.response, v.isCorrect))
	}

}
