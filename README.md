# Google AI Hackathon (maybe others)

## How to run everything
```
// need to be in this directory
cd ./services/

// Builds all your packages and updates Docker configs
docker compose build

// Spins up your containers: backend, frontend, and proxy
docker compose up
```

## Go backend (without Docker)
### How to run (locally)
```
// Go to the directory
cd google_ai_hackathon/chat-backend

// Install dependencies
go mod download

// Run server on port 1122 or whatever is specified in main.go
go run main.go
```