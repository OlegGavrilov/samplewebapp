# Sample web app

This is an example of realtime web chat app written in Go. 

I used Gin to serve static files and provide middleware for websocket handler, and a simple goroutine to handle messages that was posted by WS handler. 

# How to build and run

`$ go get github.com/OlegGavrilov/samplewebapp`

Navigate to %GOPATH/src/github.com/OlegGavrilov/samplewebapp and then

`$ go build`

This produces command file (executable), which could be run 

`$ simplewebapp`

To access UI, open browser on http://localhost:8080/
