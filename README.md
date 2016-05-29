= Install

```
$ go get -u github.com/reinerRubin/loli-file-server/cmd/loli_file_server
$ go get -u github.com/reinerRubin/loli-file-server/cmd/loli_file_client
```

= Usage

```
$ md5sum ~/Images/todo/1447547695583.jpg
7f87243fe8a03ceacb6f24793b196fbc  /home/kotya/Images/todo/1447547695583.jpg
$ loli_file_server ~/Images/todo
2016/06/03 02:36:21 start listening 50051
2016/06/03 02:36:21 start file dir reading
2016/06/03 02:36:21 file cnt 31
2016/06/03 02:36:21 start server
...

$ loli_file_client ls localhost:50051
1444251597626.jpg
1447542912513.jpg
1447547695583.jpg
1451170244218.webm

$ loli_file_client cp localhost:50051 1447547695583.jpg /tmp
$ md5sum /tmp/1447547695583.jpg
7f87243fe8a03ceacb6f24793b196fbc  /tmp/1447547695583.jpg

$ # cp behavior
$ mkdir /tmp/wow
$ loli_file_client cp localhost:50051 1447547695583.jpg 1451170244218.webm 1444251597626.jpg /tmp/wow
$ ls /tmp/wow
1444251597626.jpg  1447547695583.jpg  1451170244218.webm

$ ls /tmp/k.jpg
ls: cannot access '/tmp/k.jpg': No such file or directory
$ loli_file_client cp localhost:50051 1447547695583.jpg /tmp/k.jpg
$ ls /tmp/k.jpg
/tmp/k.jpg

```

= build
```
$ # generate protobuf files
$ protoc -I ./proto ./proto/service.proto  --go_out=plugins=grpc:pb
$ # build
$ go build -o bin/loli-file-client ./cmd/loli_file_client/main.go
$ go build -o bin/loli-file-server ./cmd/loli_file_server/main.go
```

= TODO

* Изменит контекст клиента (убрать context.Background())

* Полученный файлы писать сначала во временные файлы и в случае успеха делать mv

* Добавить в cp команду ключ -r и поддержку масок в путях

* Само собой всё переписать с "0"