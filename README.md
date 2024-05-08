# GoLang WebSocket Chat Server

This is a GoLang WebSocket chat server that allows users to join chat rooms, send, and receive voice data concurrently.

## Features

- Real-time voice data exchange using WebSocket.
- Multiple users can join different chat rooms concurrently.
- Efficient handling of data streams using goroutines.

## Installation

1. Make sure you have GoLang installed on your system. If not, you can download and install it from the official website: [https://golang.org/](https://golang.org/)

2. Clone this repository:
   ```bash
   git clone https://github.com/Purva2803/Varta-Lap.git

3. Navigate to the project directory:
   ```bash
   cd Varta-Lap

4. Install dependencies:
   ```bash
   go get github.com/gorilla/websocket


## Usage

1. Start the server:
```bash
   go run server.go

2. Connect to the WebSocket endpoint using a WebSocket client (e.g., browser, `wscat`, Postman).
Example WebSocket endpoint: `ws://localhost:8080/ws?room=your_room_name`

3. Join a chat room by specifying the room name in the query parameter (`room`). If the room doesn't exist, it will be created automatically.

4. Send and receive voice data in real-time with other users in the same chat room.

## Endpoints

- `/ws`: WebSocket endpoint for connecting to the chat server. Users must specify the room they want to join via the `room` query parameter.



