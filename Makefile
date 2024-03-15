CMD = ./cmd

test:
	go test ${CMD}/search
	go test ${CMD}/sort
	go test ${CMD}/structures
	go test ${CMD}/pathfinding
	go test ${CMD}/trees
