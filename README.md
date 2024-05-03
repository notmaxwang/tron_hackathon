# ✨ KeyFi - Your AI-powered Home Purchasing Platform

## Elevator pitch
This is KeyFi, a real-estate website that incorporates Generative AI and Blockchain technology to democratize home buying process and reduce overall transaction fees from 5% down to 1%.

## Try it out!!!
At the time of writing this README, the Google ChatSession API is broken (because they added a new field in their API response which broke our protobufs). However, it may be fixed by the time you click on it. So try it out!
http://ec2-34-236-81-43.compute-1.amazonaws.com

But it worked the morning before I recorded the demo. Here is a terminal log from right before the API response change:
![working chat](https://github.com/buzzcrackle/google_ai_hackathon/blob/main/working_chat.png?raw=true)

## Note for Google AI Hackathon Judges
The Gemini code exists in ```./services/keyfi-backend/chat/ws_interface.go```

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
