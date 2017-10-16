# Gounity
Dell EMC Unity package that provides API bindings for Go.

# Examples
## Initialize a new Unity session

```Golang
server := "myunity.example.com"
insecure := true
username := "myuser"
password := "mypassword"

session, err := gounity.NewSession(server, insecure, username, password)
if err != nil {
    log.Fatal(err)
}
```

## Get Unity's DNS configuration

```Golang
err = session.GetDNSServer()
if err != nil {
    log.Fatal(err)
}
```

## Create a realtime metric's query

```Golang
// metric paths
paths := []string{
    "sp.*.cpu.summary.busyTicks",
    "sp.*.cpu.uptime",
    "sp.*.storage.pool.*.sizeFree",
    "sp.*.storage.pool.*.sizeSubscribed",
    "sp.*.storage.pool.*.sizeTotal",
    "sp.*.storage.pool.*.sizeUsed",
    "sp.*.storage.pool.*.sizeUsedBlocks",
    "sp.*.memory.summary.totalBytes",
    "sp.*.memory.summary.totalUsedBytes",
}

// metric interval in second
var interval uint32
interval = 60

Metric, err := session.NewMetricRealTimeQuery(paths, interval)
if err != nil {
    log.Fatal(err)
}
```

# Author

**Erwan Quélin**
- <https://github.com/equelin>
- <https://twitter.com/erwanquelin>

# License

Copyright 2017 Erwan Quelin and the community.

Licensed under the MIT License.
