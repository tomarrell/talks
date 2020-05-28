build:
	go build race.go

run:
	go run race.go

# To show it actually does sometimes print things...
runlots: build
	for i in {1..1000}; do ./race; done

# Let's try with a WaitGroup
wg:
	go run wg.go
	
# So what's the problem?
detect:
	go run -race wg.go

# Now let's fix it!

