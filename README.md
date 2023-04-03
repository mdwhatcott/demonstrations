# httpreq

What's wrong with the following function call?

```
request, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
```

Passing an explicit `nil` to a function is a design smell. But at this point, any change to that function signature
would break the [Go compatability promise](https://go.dev/doc/go1compat).

So, herein I present an approach (based on Dave Cheney's excellent ["Functional Options for Friendly APIs"](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis))
to library design that follows the [open-closed principle](https://en.wikipedia.org/wiki/Open%E2%80%93closed_principle):

"Modules should be closed for modifification, open for extension."

Enjoy!