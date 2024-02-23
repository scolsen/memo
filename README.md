# memo

<small>A utility for remembering things on the command line.</small> 

`memo` is a command line utility that captures and stores the previous argument
passed to it. When invoked without an argument, it prints this previously stored
value:

```sh
$ memo "remember: have a great day" 
# ... a bit later on...
$ memo
$ "remember: have a great day"
```

If you pass multiple, space delimited strings to `memo`, it recalls it all as
one "memory":

```sh
$ memo have a great day
$ # ... a bit later on...
$ memo
$ have a great day
```

`memo` can only remember one thing. If you give it something else to "remember"
it will forget whatever you stored before:

```sh
$ echo "$(./memo)" | sed s/great/bad/g | memo
$ # ... a bit later on...
$ memo
$ have a bad day
```

All memo really does is write whatever you give it to a temporary file,
`/tmp/.memo`. Later, it reads that file. If you need to, you can always get rid
of your memo file using other tools like `rm`. If your system doesn't have a
`/tmp` directory, memo won't work. Luckily, the storage location is easy to
change: clone this repository, edit memo.go, and run `go build`.

This utility is mostly useful to support scripting. For example, you can store
the output of one command with `memo` then use process substitution to recover
it in another script or context:

```sh
$ memo "remember: have a great day"
$ echo "Here's what I $(memo)"
"Here's what I remember: have a great day"
```

Note: If you're using an older shell, it may not support the `$(...)` process
substitution syntax.

I tend to use `memo` to save the output of a long pipeline for later. In
particular, I use it to store the output of
[`fzf`](https://github.com/junegunn/fzf) to make interactive data selection
pipelines:

```sh
$ ...some data generation pipeline | fzf | memo
# ... some time later
echo "I selected datum: $(memo)"
```

That's all that `memo` has to offer. If you want to use it, either:

1. Run `go install github.com/scolsen/memo` to install the `memo` binary.
2. Clone this repository, tweak `memo.go` as much as you'd like, and run `go
   build` to build the binary yourself.

In the future, `memo` might be updated to accept a few flags that slightly alter
its behaviors, but its basic functionality won't change.

