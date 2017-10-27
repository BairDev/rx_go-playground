## RxGo Playground

Here I'm playing with [reactive programming](https://gist.github.com/staltz/868e7e9bc2a7b8c1f7549) and use Go for doing this.

My current Go version is 1.9, for installing Go have a look at [their Getting Started page](https://golang.org/doc/install).

Using the RxGo library was my first guess how to start, but it might be a [bad descision](https://news.ycombinator.com/item?id=13562538). I will see this in the future.

[RxGo](https://github.com/ReactiveX/RxGo) is apparently still work in progress, I hope that it is and will be a vivid project.

### Compiling

You can compile the most simple test program only for now. It is more or less C&P from the introduction to *RxGo*.
For compiling I'm currently using this command (in `rx_go/src/test-rx/`):

    go build -o /your/path/to/rx_go/bin/rx1 rx1.go

### Executing

    /your/path/to/rx_go/bin/rx1
    # -> 20