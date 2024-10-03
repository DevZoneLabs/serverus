## This is the documentation referred to the WarCraft Logger Bot

The `WarCraft-Logger` bot will be connected to the `WarcraftLogs` website via webhooks, which will allow to listen for events an provide images related to the logs, which allows the Discord users to navigate into the website uncessarily. 

### Implementation

#### Events

1. When the bot is added.
    1. The bot will try to create two channels:
        * `warcraft-logger-admin`
            1. Private channel, which should be accessible | visible by guild admins (initially)
        * `warcraft-logger-bot-priv`
            1. This is a private channel to be used by the bot, ideally we will not want this channel to be visible by any other member of the guild.
    2. After the channels are created the bot will create a webhook associated with `warcraft-logger-bot-priv` and present the admin in the `warcraft-logger-admin` with instructions on how to setup the bot to start listening for.
    3. After the webhook is connected, no further action will be required from the user at the moment.
        1. The plan will be to have a set of functions the user will be able to do.
        
#### Storage - Relations
1. If this bot is to be used by multiple servers, it will be crucial to know the relationship between the incoming event, the channel that triggered it and where the result should be posted.
    1. The easier at efficient way at the moment will be to:
        1. Associate `guild.ID` with the output `channel.ID` OR
        2. Associate `priv-channel.ID` with the output `channel.ID`
        3. Keep this relationship in a persistent storage. Some simple database or file?

####

### Basic requirements from the user:
* A channel name for the bot to return the images
* Connect the provided webhook url to the `WarcraftLogs`



