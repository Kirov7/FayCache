# FayCache
## An attempt at a distributed cache scheme
## Non-production level projects !!!
## Just to learn and test my personal ideas !!!

Temporarily using the sdk
import by go mod
```shell
go get https://github.com/Kirov7/FayCache
go mod tidy
```
test open server
```go
func TestServer(t *testing.T) {
	c := cache.New(cache.STORAGE_INMEMORY)
	n, _ := cluster.NewCluster("127.0.0.1", "")
	tcp.NewTCPServer(c, n).Listen()
}
```
test client connection
```go
func TestClient(t *testing.T) {
	cacheClient := client.New(client.SERVER_TCP, "127.0.0.1")

	cmds := []*client.Cmd{}
	for i := 1; i <= 100; i++ {
		var op client.COMMAND_TYPE
		if rand.Intn(2) == 1 {
			op = client.COMMAND_SET
		} else {
			op = client.COMMAND_GET
		}
		op = client.COMMAND_GET
		tempCmd := &client.Cmd{
			Name:  op,
			Key:   fmt.Sprintf("hello%d", i),
			Value: fmt.Sprintf("word%d", i),
		}
		cmds = append(cmds, tempCmd)
		if i%5 == 0 {
			cacheClient.PipelinedRun(cmds)
			if tempCmd.Error != nil {
				fmt.Println("error:", tempCmd.Error)
			}
			for _, cmd := range cmds {
				fmt.Println("key: ", cmd.Key, "value: ", cmd.Value)
			}
			cmds = []*client.Cmd{}
		}
	}
}
```

test set
```go
func TestSet(t *testing.T) {
	cacheClient := client.New(client.SERVER_TCP, "127.0.0.1")
	cmds := []*client.Cmd{}
	for i := 1; i <= 100; i++ {
		var op client.COMMAND_TYPE
		if rand.Intn(2) == 1 {
			op = client.COMMAND_SET
		} else {
			op = client.COMMAND_GET
		}
		op = client.COMMAND_SET
		tempCmd := &client.Cmd{
			Name:  op,
			Key:   fmt.Sprintf("hello%d", i),
			Value: fmt.Sprintf("word%d", i),
		}
		cmds = append(cmds, tempCmd)
		if i%5 == 0 {
			cacheClient.PipelinedRun(cmds)
			if tempCmd.Error != nil {
				fmt.Println("error:", tempCmd.Error)
			} else {
				fmt.Println(tempCmd.Value)
			}
			cmds = []*client.Cmd{}
		}
	}
}
```

test get
```go
func TestGet(t *testing.T) {
	cacheClient := client.New(client.SERVER_TCP, "127.0.0.1")
	for i := 1; i <= 100; i++ {
		var op client.COMMAND_TYPE
		op = client.COMMAND_GET
		tempCmd := &client.Cmd{
			Name:  op,
			Key:   fmt.Sprintf("hello%d", i),
			Value: fmt.Sprintf("word%d", i),
		}

		cacheClient.Run(tempCmd)
		if tempCmd.Error != nil {
			fmt.Println("error:", tempCmd.Error)
		} else {
			fmt.Println(tempCmd.Value)
		}
	}
}

```
