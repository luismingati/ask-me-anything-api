# API Documentation

This documentation covers the available endpoints for the Go-based API that handles room creation, message management, and WebSocket subscriptions.

## Endpoints

### Subscribe to Room Messages (WebSocket)

- **GET** `/subscribe/{room_id}`: Subscribe to a room for real-time message updates.

### Room Management

- **POST** `/api/rooms`: Create a new room.
  - **Request Body**: `{ "theme": "Room theme" }`
  - **Response**: `{ "id": "room_id" }`

- **GET** `/api/rooms`: Retrieve all rooms.
  - **Response**: `[ { "id": "room_id", "theme": "Room theme" } ]`

### Message Management

- **POST** `/api/rooms/{room_id}/messages`: Create a new message in a room.
  - **Request Body**: `{ "message": "Message content" }`
  - **Response**: `{ "id": "message_id" }`

- **GET** `/api/rooms/{room_id}/messages`: Retrieve all messages in a room.
  - **Response**: `[ { "id": "message_id", "content": "Message content" } ]`

- **GET** `/api/rooms/{room_id}/messages/{message_id}`: Retrieve a specific message.
  - **Response**: `{ "id": "message_id", "content": "Message content" }`

- **PATCH** `/api/rooms/{room_id}/messages/{message_id}/react`: React to a specific message.
  - **Response**: `{ "count": "Reaction count" }`

- **DELETE** `/api/rooms/{room_id}/messages/{message_id}/react`: Remove a reaction from a specific message.
  - **Response**: `{ "count": "Reaction count" }`

- **PATCH** `/api/rooms/{room_id}/messages/{message_id}/answer`: Mark a message as answered.
  - **Response**: `HTTP 200 OK`

## WebSocket Messages

WebSocket endpoint `/subscribe/{room_id}` will push real-time updates:

- **Message Created**
  ```json
  { "kind": "message_created", "value": { "id": "message_id", "message": "Message content" } }
  ```
- **Message Reaction Increased**
  ```json
  { "kind": "message_reaction_increased", "value": { "id": "message_id", "count": "Reaction count" } }
  ```
- **Message Reaction Decreased**
  ```json 
  { "kind": "message_reaction_decreased", "value": { "id": "message_id", "count": "Reaction count" } }
  ```
- **Message Answered
  ```json
  { "kind": "message_answered", "value": { "id": "message_id" } }
  ```

## Technologies Used

- Router: chi
- WebSockets: gorilla/websocket
- Database: pgx (PostgreSQL driver)
- Query Builder: sqlc
- Migrations: tern


