# Code Examples from Black Hat Go 

## The book is Black Hat Go by Tom Steele: [https://nostarch.com/blackhatgo][1]

## To run:
- Assumes you have Golang installed: [https://golang.org/doc/install][2]
- Note: these are basically stand alone scripts and as such they each have their own main function. I may change this in the future because it is not in keeping with Go conventions.
- In this directory simply run `go run <name of go file>`. E.g., `go run tcp_scanner_slow.go`
- Alternatively you can run `go build <name of go file>` and then `./<name of executable>`. E.g., `go build tcp_scanner_slow.go` and then `./tcp_scanner_slow`.
- Similarly, for tests simply run `go test -v <name of go test file>`. There are no tests right now.

## The examples:

- `tcp_scanner_slow` scans scanme.nmap.org:* where * is the port numbers 1 through 100. Ports are scanned sequentially in a simple `for` loop. A simple print statement writes the status of the port (either open or closed) to standard output. 
- `tcp_scanner_fast` does the same as above with one notable exception: ports are scanned simultaneously using `goroutines` which means go will automatically open up as many `goroutines` as the system can handle at one time. This significantly reduces the amount of time it takes to scan a range or ports. I also increased the port range to 1024 and removed the print statement for closed ports so the output is more readable.
- `tcp_scanner_pool` does the same as above with one notable exception: a worker pool is created and channels implemented.  This helps ensure more consistent results by preventing too many ports being scanned simultaneously.  Now you can set it to whatever value you like: 100 workers is hard coded, but you can change it before running it.  This way, assuming you choose a balanced value, network or system limitations won't skew your results.

## Notes: 
- When port scanning I open up another terminal and ran `nettop` to watch the open TCP connections. That is one way to confirm the behavior of the tcp scanner scripts.
- I may at times modify the code examples. 
- These are for education and fun. 
- Please don't run the code here without knowing what you are doing first. For example, running un-authorized port scans of networks that you don't own can get you into legal trouble.

[1]:https://nostarch.com/blackhatgo
[2]:https://golang.org/doc/install