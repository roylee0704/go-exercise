## fetchURls

Concurrently fetches many URLs, so that the process will take no longer 
than the longest fetch rather than the sum of all fetch times.

### To Run
```sh
$ go run main.go https://golang.org https://gopl.io https://godoc.org

0.82s    8447 http://golang.org
0.82s    6927 http://godoc.org
1.17s    4154 https://gopl.io
1.17s elapsed
```