# gonroni: Negroni middleware for GoNR NewRelic plugin

gonroni is a simple package that creates Negroni middleware to display HTTP response time metrics in NewRelic via the GoNR NewRelic plugin.

# Installation

```go
go get github.com/neocortical/gonroni
```

# Use

```go 
import (
  "github.com/neocortical/gonr"
  "github.com/neocortical/gonroni"
)

func main() {
  gonrAgent := gonr.New(gonr.DefaultConfig().WithLicense("abc123"))
  
  n := negroni.New()
  n.Use(gonroni.AddNegroniHttpMetrics(gonrAgent))
  // ...
}
```

# Notes

GoNR and gonroni are very new, but will be in a production environment soon. I can't promise nothing will change at this point, but things are starting to look prety stable. 

If you would find some change or feature useful, I would love the feedback! Submit an issue or a pull request.

I am probably going to add error metrics of some kind to the middleware (e.g. rate of 404, 500, etc. responses). If you have any ideas, let me know!
