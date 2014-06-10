# Server
A single index server, it is not cluster aware and runs on port 4648.

# API

### POST /query?cb

````javascript
{
  filter: {
    and:[{
      op: eq,
      term1: {
        source:
        val:
      },
      term2: {
        source:
        val:
      }
    },{

    }]
    or
  }
  wkt: ""
  id: 'the value to return'
  order: []
  limit: 10
  offset: 0
}
````

### POST /index?cb

````javascript
{
  wkt
  properties
}
````