# the `.dev` folder

utilites to help with developing and provisioning the on call bot.
these are not "first class" citizens of the project yet and may not work reliably on your machine.

### main.go

CLI utility to help generate a Slack app manifest with the proper scheme and host of the domain the you would like to host the bot at.

### init_db.sh

a basic shell script which ensures the local docker compose db is setup