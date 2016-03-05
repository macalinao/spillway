# spillway

Rate limit console output.

Sometimes you're `k8read`ing a bunch of JSON or you're tailing some crazy logs. `spillway` allows you to limit how often lines get written to the terminal.

## Instaling

```bash
$ go get github.com/simplyianm/spillway
```

## Usage

```bash
# Print a random line every second
cat /dev/urandom | spillway -r 1s
```

## License

MIT
