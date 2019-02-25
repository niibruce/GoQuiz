# GoQuiz

Dead Simple Command line application to run timed, arithmetic quizzes.


go run ./GoQuiz -csv {csv filename} -time {time limit in seconds} -shuffle {shuffle}

### Arguments
* filename - name of the CSV file containing the questions. One question per line in the format: Question, Answer. Default value is _problems_test.csv_.
* time - time in seconds available to answer all questions. Quiz stops after the time is reached. Default value is _5_ seconds.
* shuffle - optional flag to shuffle the questions
