all: build

build: saucy.go
	go build .

parser: Saucy.g4
	java -Xmx500M -cp "/usr/local/bin/antlr-4.7.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool -Dlanguage=Go -o parser Saucy.g4
