# Services

### Serverus\_Bot

This service is out main service and possibly the one that will require the most amount amount of logic/effort. Let's try to keep documentation up to date in regards to all the changes we make to it.

Purpose:
This service has two main purposes as it takes care of:
1. Receiving requests and interaction through the Discord App to later process them and interact with other services. 
2. Receives request from the external API to process them and send them through Discord.

#### package\_structure

- **cmd/server/**: Contains the main application executable and any supporting files.
    - **main.go**: Entry point of the application. It will start a DiscordBot Session as well as a server.
    - **routes.go**: Contains the code responsible for defining the routes and corresponding handlers for the HTTP endpoints of the application.
- **internal/**: Contains packages that internal to the project and are not intended to be used outside the project.
    - **bot/**: Source code for bot implementation / handlers.





