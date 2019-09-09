unitest:
	go test -v -failfast ./graph/*.go

example:
	go test -v -failfast ./sensor/*.go
