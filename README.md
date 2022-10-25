# SimpleChatRoom
The goal is to create a simple chat room that supports private messages between two users, allowing for many users to be connected to the server at the same time.

## Running Instruction
- To officiall start, first will run ```chatRoom.go``` to establish the server. Then in a seperate terminal, run ```client.go```. In order to see it is fully functional, it is suggested to run at least 2 ```client.go``` in 2 seperate terminals(so you can see this is a simulated function of the chatroom)

- Here are the specific instruction on running each of the file
   - To run the ChatRoom code, do   ```go build ChatRoom.go```, this will generated an executable file in the directory
   - To compile the Client code, do ```go build Client.go```, this will generate an executable file in the directory, likewise.

To run the ChatRoom, do ```./ChatRoom PORT```, where PORT is a valid port number. 

The server will then start up. To close the server, type "EXIT". In addition, typing ```Ctrl + C``` will also be read as terminating the server. 

To run the Client, do ```./Client ADDRESS PORT USERNAME```, where ADDRESS is the address of the server(the info address will be displayed if the Server runs successfully),  ```PORT``` is the port the server is on, and ```USERNAME``` is the desired username for the server. Please note ```USERNAME``` must be unique for each different client(if the user acciedently supplies with a username that's already registered at the server, the terminal will prompt the reminder). Termination is the similar to the server side: type ```EXIT``` at any point in the program, or ```Ctrl + C``` will close the connection, immediately. 

If another client is already connected to the server with the same username, the user will be prompted to put in another username. To send messages, the client is prompted for the username they want to send the message to, and then prompted again for the message they want to send. If there are any difficulties delivering the message, the server will send an error message back.

## Code Description
- ChatRoom.go: This is the chatroom server, that handles all the client connections and passes on messages. There are the following threads:
  - One thread reads command-line input waiting for the user to type exit EXIT
  - One thread listens for new connections, which generates one thread per connection
  - n threads, where n is the number of connections, which read new messages from a connection and passes them to a channel
  - One thread total sends messages to the clients
- Utils.go: This saves shared structs and functions
- Client.go: This is the client, which connects to the chatroom and sends and recieves messages from the chatroom.
  - First the client connects to the server
  - Then the client sends usernames to the server until it sends a username that is not in use
  - Then the client sends messages until it receives EXIT

## Explanation of Design
Here are a few particular aspects of our design we wanted to explain in more detail.
### Server Side
- We save the connections within a slice to keep track of them
- There is a mutex whenever we set the username, to prevent race conditions. We do not use the mutex whenever we delete connections, as the delete operation is idempotent.
- A new thread is created for each connection, as reads are blocking
### Client Side
- The client checks for a few reserved usernames, such as chatroom, and will print an error message if that occurs. This is in addition to the server-side checking, just to make it let the user know quicker their username does not work.
