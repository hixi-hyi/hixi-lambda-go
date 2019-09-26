
tests:
	go test ./... -v -count=1

test-%:
	go test ./${*} -v -count=1
