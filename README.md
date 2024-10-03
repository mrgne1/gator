# Gator

This is a blog aggregator project created from a Boot.Dev course

Installation
---
You'll need [PostGres](https://www.postgresql.org/download/) and [Go](https://go.dev/doc/install) Installed to use this.

Once those are installed, you can navigate to the root directory and type `go install gator` to install this aggregator.

Setup
---
Type `gator` at the command line once to create the initial configuration file in your home directory.

Navigate to your home directory and edit the `.gatorconfig.json` file
Change the "connection_string" entry to be the connection string for your PostGres database.
The format should be postgres://{{user name}}:{{user password}}@{{server}}:{{port}}/gator

Commands
---
register: Create a new user in gator. 

login: Switch to a user already registerd in gator

users: List users. Will indicate current user

addfeed: Add a feed to gator, current user will automatically follow

feeds: List feeds in gator

follow: Add feed to the current user's follows

following: List feed follows of current user

unfollow: Stop following a feed

agg: continuously download feeds.

browse: look at downloaded feeds
