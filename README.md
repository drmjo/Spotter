# Spotter
Here to help you bench.

##What exactly does it do?
Spotter will be a command-line tool to help you benchmark your web servers. 

Morally distinct from another existing tool... 

##What it can currently do:
`go run main.go -n 10 -c 5 -r "get" http://github.com`

##What it currently returns:
![example](./example.jpg)

##Planned flags:
-n : Request Number
-r: Request Type
-h: Request Header
-c: Concurrency Level