# GO usb relay

![Imgur](https://i.imgur.com/zuNIMH9.jpg)

A driver for control usb relay module implement with Golang.
We access devices through karalabe/hid (with [karalabe/hid#13](https://github.com/karalabe/hid/pull/13) merged).

## Support
|    OS     |  Is supported |
|:---------:|:-------------:|
| MacOS     |  Yes          |
| Windows   |  Yes          |
| GNU/Linux |  Yes ( No test yet ) |

## Install
```sh
go get -v -u github.com/e61983/go-usb-relay
```
## Testing
```sh
$ make test
```

## Build
```sh
$ make
```

## Usage
```sh
# Trun On channel 1 relay.
$ release/darwin/amd64/go-usb-relay -n 1 -o

# Set SN
$ release/darwin/amd64/go-usb-relay -sn 12345
```
