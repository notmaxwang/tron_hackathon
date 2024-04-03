# Google AI Hackathon (maybe others)

## Go backend
### How to run (locally)
```
// Go to the directory
cd google_ai_hackathon/chat-backend

// Install dependencies
go mod download

// Run server on port 1122 or whatever is specified in main.go
go run main.go
```

### Common issues
You may see```protoc-gen-grpc-web: program not found or is not executable``` or something similar

```
// these are the ones I remember
brew install protoc-gen-grpc-web

npm install -g protoc-gen-js
```
