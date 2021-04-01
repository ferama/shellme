# ShellMe

Shellme is a full web terminal server micro service.

It supports escape sequences so you can run vim too :). It is written in golang

![Image of Shellme](https://raw.githubusercontent.com/ferama/shellme/master/docs/shellme.gif)

## How to run

You can build a binary executable cloning this repo and running the **build.sh** script.
You will need:

* go 1.16
* node 14

For a quick tryout you can run an instance inside an ubuntu based container using:

```
$ docker run -p 8000:8000 ferama/shellme
```

## Security Warning

Be warned: launching shellme locally as is will expose on 8000 a full shell with your user privileges. It is highly recommendend that you put a security authenticated proxy on front of it