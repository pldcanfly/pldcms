clean:
	rm -rf bin/*

run:
	air

build:
	go build -o bin/pldcms cmd/pldcms/pldcms.go


