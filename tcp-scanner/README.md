# TCP Scanner

Scan all 65,535 ports on localhost using Goroutines, a Worker Pool, and multichannel communication.

A Worker Pool is used to avoid inconsistencies due to network or system limitations when scanning a large number of ports simultaneously.

This is from the lesson in the **Black Hat Go** book.

## Usage

Run in Terminal

```
go run .
```
