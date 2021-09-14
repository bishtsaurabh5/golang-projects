# Quiz project spec
* Input a csv file that will contain questions ,answer and points per question
* csv file should have question,answers,points format
* the program should take csv as input via flag
* A time limit should be considered for timing out the questions
* Use time package and encoding/csv

## Key Takeaways
* Learnt how to use go interfaces
* Here the Reader interface was being implemented by the File struct
* So just passing the output of os.Open() which is a file pointer to the csv.ReadAll retrieved the complete csv
* Learnt how to use go buffered channels here(Read inline comments in code for more)
* flag package can be used to parse the flags using the --<flag_name>
