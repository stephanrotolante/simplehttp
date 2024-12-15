Stupid simple http client wrapping the go lang http package


```
import "github.com/stephanrotolante/simplehttp"

func main() {
	httpRequest := simplehttp.CreateHttpRequest("https://google.com")

	httpRequest
		.Post()
		.Body([]byte{})
		.Execute()
}

```