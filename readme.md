# Go learning

Playing around with Go. 

```bash
wget https://go.dev/dl/go1.22.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz


# in .bashrc
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$(go env GOPATH)/bin

# test with go version
go version

# seems like a lot of work just to get the newest version. 
# sudo apt install golang-go would give an older version 
# that doesn't have the latest features.
```

## air

the idea is to do live reloads to make life easier. 


SEe this pagehttps://github.com/air-verse/air


```bash
air init
```

which will make `.air.toml` file

```bash
air -c .air.toml
```

I had to move a1, a2, A3 to a folder called old.

## run

now that everything is setup I can just run the server with

```bash
air 
```


OR

```bash
go run A4_gin_website.go
# // or build and then run
go build -o A4_gin_website A4_gin_website.go
 ./A4_gin_website
```