# upgraded-broccoli
@solvaholic is learning to use Go

## Black Hat Go, from nostarch.com

> It's like C, but...

If you're tired of :point_up: that phrase, then check out [Black Hat Go](https://nostarch.com/blackhatgo) from no starch press.

Learn Go while you learn networking. What's not to love?!

See also [Effective Go](https://go.dev/doc/effectivego) and [Go Wiki: Go Code Review Comments](https://go.dev/wiki/CodeReviewComments).

## ./port-scanner
A simple port scanner.

```bash
go run port-scanner/main.go
```

1. port-scanner dials scanme.nmap.org TCP ports 1 through 100
1. And outputs `NN open` if port `NN` is open

## ./echo
A simple echo server.

1. Start the echo server
    ```bash
    go run echo/main.go
    ```
1. echo listens for text, echoes it back
    ```bash
    telnet 127.0.0.1 20080
    Hello World
    ```

## ./gaping-hole
An unsecure remote shell.

1. Start the gaping-hole server
    ```bash
    go run gaping-hole/main.go
    ```
1. gaping-hole listens for text, runs in the remote shell
    ```bash
    telnet 127.0.0.1 20080
    echo "Hello World"
    ```
