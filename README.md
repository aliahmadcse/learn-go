# Learning GoLang

Let's get our hands dirty with Go 🚀

## Configuring GVM (GVM is not actively maintained and has issues, instead use Goland SDK manager feature)

- I am using Go Version Manager (GVM) to install and manage multiple version of Golang
- Install it using the following commands
``` 
$ sudo apt-get install bison

$ bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```
- Now that the GVM is installed on your machine, you can verify it using the following command
- `gvm version`

## Install Go

You can use GVM to install any or multiple versions of GoLang. I am using version `1.23.1` and it can be installed using

1. `gvm install go1.23.1 -B`
2. Set it the default version using `gvm use go1.23.1 --default`


## Run a file
- Use `go run main.go` to execute a particular file
