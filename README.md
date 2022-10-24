# SimpleChatRoom
The goal is to create a simple chat room that supports private messages between two users, allowing for many users to be connected to the server at the same time.

## How to Use
To compile the ChatRoom code, do "go build ChatRoom.go". To compile the Client code, do "go build Client.go"
To run the ChatRoom, do "./ChatRoom PORT", where PORT is a valid port number. The server will then start up. To close the server, type "EXIT"
To run the Client, do "./Client ADDRESS PORT USERNAME", where ADDRESS is the address of the server, PORT is the port the server is on, and USERNAME is the desired username for the server. If another client is already connected to the server with the same username, the user will be prompted to put in another username. To send messages, the client is prompted for the username they want to send the message to, and then prompted again for the message they want to send. If there are any difficulties delivering the message, the server will send an error message back.

## code description
- ChatRoom.go: This is the chatroom server, that handles all the client connections and passes on messages. There are the following threads:
  - One thread reads command-line input waiting for the user to type exit EXIT
  - One thread listens for new connections, which generates one thread per connection
  - One thread per connection, which reads new messages and passes them to a channel
  - One thread total sends messages to the clients
- Utils.go: This saves shared structs and functions
- Client.go: This is the client, which connects to the chatroom and sends and recieves messages from the chatroom.
  - First the client connects to the server
  - Then the client sends usernames to the server until it sends a username that is not in use
  - Then the client sends messages until it receives EXIT
