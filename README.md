Package helps to get valid timezone names

Now working only on *nix and Mac Os 

**Examples**


you can use `repository` if you want to read timezones once and store in memory for quick access  
```go
package main

import "fmt"

func main() {
    tzRepository := NewRepository()
    tzlist := tzRepository.GetTimezones() 	
    
    fmt.Print(tzlist)
}
```

if you don't need to store timezones in memory, you can use

```go
tzlist := GetTimezones()
```