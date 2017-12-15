USER_TEST=tests/user_model_test.go
CLASS_TEST=tests/class_model_test.go
ASSIGN_TEST=tests/assignment_model_test.go
SUBMIT_TEST=tests/submission_model_test.go
default:
	./start.sh

test:
	go test $(USER_TEST)
	go test $(CLASS_TEST)
	go test $(ASSIGN_TEST)
	go test $(SUBMIT_TEST)
