# Cluster
Aware of the topology of the flechette cluster, on read it performs a scatter-gather query, and routes writes in a balanced way.

# Topology Options
The cluster need not be on the same machine(s) as the servers, but is by default. (Each server runs a cluster instance as well).

# Scatter-Gather
Done via a pure asynchronous query, a uuid'd callback url is generated and sent with each server query, which is immediately closed. When the query is complete, each server calls the callback url with the resulting data, which is interleaved into the return. When all servers have responded or the timeout has been reached, the data is returned to the client.

# External API (Port 4646)

### GET /

### POST /query

### POST /index

### POST /bulk

# Internal API (Port 4647)

### GET /status

### POST /callback/index

### POST /callback/query