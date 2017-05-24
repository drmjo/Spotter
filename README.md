# Spotter
Work In Progress.

## What Exactly Does Spotter Do?
Spotter is a tool to help you benchmark web servers. 

## Requirements
Built using Go 1.7

## Want To Try It Out?
- Assuming you have Go 1.7 installed on your system and the proper environment variables setup ($GOBIN), 
run `go install github.com/hunterel/spotter`
- Spotter implements the following command-line options
* `-n` Number of HTTP requests to make.
* `-c` Number of workers to make HTTP requests.
* `-r` Request method to use (GET, POST, PUT, DELETE, OPTIONS, HEAD)
* `-o` A File path location for an output JSON file.
* `-d` Request body for the HTTP Request (Can use files instead by specifying "@relativeFileLocation")
* `-h` Request Headers. Repeatable command.
- Run a benchmark, `spotter -n 100 -c 10 -r "POST" -d "{"test":"test"}" -h "access-token: superSecureToken" www.MyDomainToTest.com/input`
- Output is currently bucketed by status code and returned as a formatted string, but will eventually be returned as JSON.

## Comments / Suggestions?
Please make an issue or pull-request. I'm making this tool to help me with work, but if there is a feature you would like or problem you have, I would love to help.
