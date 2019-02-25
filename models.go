package main

type Question struct{
	question, answer, response string
	isCorrect bool
}

type QuestionCollection struct{
	Questions []Question
}