# Mastodon Importer

A simple script to start playing with mastodon importing toots. It's just a startpoint because it's not fully usable yet.

It uses a `sample.json` file as input (you can get this file in your mastodon instance server, under settings + export information).

## Guidelines

* Get your `sample.json` from your mastodon instance server
    * Prefrences
    * Import and Export
    * Request your data
    * Extract the zip and get the `outbox.json`
* Modify the code to match the path of your `outbox.json`
* Create a `.env` file with your postgresql credentials from the `.env.sample` file in the repo.
* Run and/or compile this script:
  * `go run`
  * `go build`. Take in account **where** you will run the binary. I mean I've developed it under macos darwin but I wanted to run it on my mastodon instance machine, so I compiled it with some flags in order to get it running: `GOOS=linux GOARCH=amd64 go build`
* Move the binary and the `outbox.json` to the server (ensure path matches) and profit.

### Final notes

Feel free to modify the script in order to fill your needs.
