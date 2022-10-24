# SimpleChatRoom
The goal is to create a simple chat room that supports private messages between two users, allowing for many users to be connected to the server at the same time.

## Running Instruction
- To officiall start, first will run ```chatRoom.go``` to establish the server. Then in a seperate terminal, run ```client.go```. In order to see it is fully functional, it is suggested to run at least 2 ```client.go``` in 2 seperate terminals(so you can see this is a simulated function of the chatroom)

- Here are the specific instruction on running each of the file
   - To run the ChatRoom code, do   ```go build ChatRoom.go```, this will generated an executable file in the directory
   - To compile the Client code, do ```go build Client.go```, this will generate an executable file in the directory, likewise.

To run the ChatRoom, do ```./ChatRoom PORT```, where PORT is a valid port number. 

The server will then start up. To close the server, type "EXIT". In addition, typing Ctrl-C will also be read as terminating the connection to the server. 

To run the Client, do ```./Client ADDRESS PORT USERNAME```, where ADDRESS is the address of the server(the info address will be displayed if the Server runs successfully),  ```PORT``` is the port the server is on, and ```USERNAME``` is the desired username for the server. 

If another client is already connected to the server with the same username, the user will be prompted to put in another username. To send messages, the client is prompted for the username they want to send the message to, and then prompted again for the message they want to send. If there are any difficulties delivering the message, the server will send an error message back.

## Code Description
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

## Explanation of Design
As the requirement suggested, there are two sides that we need to consider 
- Server Side
   - The server is designed to have the following algorithm for designing "algorithm"
      - 1: Like Last Homework, first will intialize the TCP Connection via the given port(Listening to TCP connection)
      - 2: It will accept a connection from any other client process by first checking the uniqueness of the name and make sure the name is validated(the goal of function getUsername and sendErr)
      - 3: The server will send the message to the desired destination specified in the message 
      - 4: The server will terminate if the user enter EXIT at any point when asked for userinput. 
- Client Side
    - Client Side works in a similar but opposite way compared to the Server side.
       - 1: The Client will input the tcp port and the username, such information is gathered to send to the Server. To validate this step, Server class is also equipped with a function to check the uniqueness
       - 2: If proper TCP Port provided, with custom username, then Client will connect to the chatroom via TCP
       - 3: The Client will keep dialing to TCP connection to send message, and will exit if the user prompt when asked.
       
