# Flechette

A read-optimized distributed linear-scaling in-memory Geospatial Database/Index.

Search for complex geometries.

# Goals
Query 1 million complex geometries, restricting and sorting on properties, and return the top 100 in ~.1 sec from a 3-node cluster.



# Architecture
Auto-sharded multi-master patterned after elasticsearch's shard/index architecture. It uses scatter-gather to query each node.


# Capabilities


# Server
[link](server/)

# Index
[link](index/)

# Cluster
[link](cluster/)