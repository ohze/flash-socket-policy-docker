# flash-socket-policy-docker
[![Build Status](https://travis-ci.org/ohze/flash-socket-policy-docker.svg?branch=master)](https://travis-ci.org/ohze/flash-socket-policy-docker)

Flash [Socket Policy files](http://www.adobe.com/devnet/flashplayer/articles/socket_policy_files.html) server implement in Go.

## Usages

```bash
docker run -d -p 843:843 sandinh/flash-socket-policy
```
or
```bash
docker run -d -p 843:843 -v `pwd`/crossdomain.xml:/crossdomain.xml sandinh/flash-socket-policy gofsp -file=/crossdomain.xml
```

## License

This project is free to use and licensed under LGPL.
