### API-Network Explanation
This module provides an API to allow nodes to connect to a backend and outsource API calls through the nodes to distribute API key usage and can allow community incentives.

## Usage

# Options
```
var newOptions helpers.OptionsType
options.Set(newOptions)
```

# Fiber Health Check Controller
e.g. In your routing
```
app.Post("/api/healthcheck", api_fiber.NodeHealthCheck)
```
This allows nodes to add themselves to the backend and show they're active

# Making a request
use the nodes.Request Method
example:
```
nodes.Request("http://127.0.0.1:4000/ping", helpers.RequestMethod.GET, nil, func(response *http.Response, err error) {

})
```

# caching
To Do