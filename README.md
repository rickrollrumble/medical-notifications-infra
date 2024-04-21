## Plan
- frontend-backend containers in one pod
- either replicas in different nodes OR
- multiple replicas in a single node, and that repeats across multiple nodes.
- all of this is in a cluster with a single node but

## Terms
- `pod` - a single container or multiple containers that together form an app
- `node` - is the server or VM on which a pod runs
- `cluster` - contains multiple such nodes
