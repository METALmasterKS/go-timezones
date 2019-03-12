Package helps to get valid timezone names

Now working only on *nix and Mac Os 

**Examples**


you can use `repository` if you want to read timezones once and store in memory for quick access  
```go
package main

import (
	"fmt"
	"time"
	"github.com/METALmasterKS/go-timezones"
)

func main() {
    tzRepository := go_timezones.NewRepository()
    tzlist := tzRepository.GetTimezones() 	
    
    location, err := time.LoadLocation(tzlist[0].String())
    if err != nil {
    	panic(err)
    }
    
    fmt.Print(time.Now().In(location))
}
```

if you don't need to store timezones in memory, you can use

```go
package main

import (
	"fmt"
	"github.com/METALmasterKS/go-timezones"
)

func main() {
    tzlist := go_timezones.GetTimezones()	
    
    fmt.Print(tzlist)
}
```