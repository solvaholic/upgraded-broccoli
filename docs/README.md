# upgraded-broccoli
@solvaholic is learning to use Go

## Black Hat Go, from nostarch.com

> It's like C, but...

If you're tired of :point_up: that phrase, then check out [Black Hat Go](https://nostarch.com/blackhatgo) from no starch press.

Learn Go while you learn networking. What's not to love?!

## ./port-scanner
A simple port scanner.

```bash
go run port-scanner/main.go
```

1. Dial scanme.nmap.org TCP ports 1 through 100
1. Output `NN open` if port `NN` is open

## ./tcp-proxy
A simple echo server.

1. Start the server
    ```bash
    go run tcp-proxy/main.go
    ```
1. Send text, receive echo
    ```bash
    telnet 127.0.0.1 20080
    Hello World
    ```

## ./gaping-hole
An unsecure remote shell.

1. Start the server
    ```bash
    go run gaping-hole/main.go
    ```
1. Run commands in the remote shell
    ```bash
    telnet 127.0.0.1 20080
    echo "Hello World"
    ```
