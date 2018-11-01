# yelp-graphql-client
Go (golang) GraphQL Client for Yelp's GraphQL API

To get an access token and learn more about the Yelp GraphQL API, See
https://www.yelp.com/developers/graphql/guides/intro



**Usage**
```go
import yelp "github.com/raja/yelp-graphql-client"

func main() {
    yelp := yelp.NewClient("your-api-token")

    // Lookup details of a business given a yelp id
    biz, err := yelp.Business("garaje-san-francisco")
    if err != nil {
        fmt.Panic(err)
    }
    fmt.Printf("we have %+v\n", biz)

    // Search for businesses based on a name and location
    businesses, err = yelp.Search("tacos", "Toronto", 10)
    if err != nil {
        fmt.Panic(err)
    }
    // Do something useful with the results
    for _, biz := range businesses {
        fmt.Printf("%+v\n", biz)
    }
}
```