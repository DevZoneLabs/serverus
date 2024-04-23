# Services

### Serverus\_Bot

This service is our main service and possibly the one that will require the most amount amount of logic/effort. Let's try to keep documentation up to date in regards to all the changes we make to it.

Purpose:
This service has two main purposes as it takes care of:
1. Receiving requests and interaction through the Discord App to later process them and interact with other services. 
2. Receives request from the external API to process them and send them through Discord.

#### package\_structure

- **serverus-bot**:
    - **cmd/**: Contains the main application executable.
        - **main.go**: Entry point of the application. It will start a DiscordBot Session as well as a server.
    - **bot/**:  Source code for bot implementation / handlers.
        - **actions.go**: Handles Customs Actions for Bot to Resolve in Discord
        - **bot.go**: Handles Bot instance, running, stopping, registering handlers.
        - **handlers.go**: Handlers event interaction from the Discord App
    - **api/**: Server Implementation
        - **server.go**: Server initialization and configuration
        - **routes.go**: Server Router configuration. Active endpoints, middlewares, routes, etc.
        - **handlers.go**: Endpoints implementation





