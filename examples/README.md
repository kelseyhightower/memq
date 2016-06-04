# Example Usage

## Create a new work queue

```
curl -X POST 127.0.0.1:8000/queues/work
```

## Add messages to the work queue

```
curl -X POST 127.0.0.1:8000/queues/work/messages -d "message 1"
curl -X POST 127.0.0.1:8000/queues/work/messages -d "message 2"
curl -X POST 127.0.0.1:8000/queues/work/messages -d "message 3"
```

## Get the broker stats

```
curl 127.0.0.1:8000/stats
```

## Read messages from the work queue.

```
curl 127.0.0.1:8000/queues/works/messages
curl 127.0.0.1:8000/queues/works/messages
curl 127.0.0.1:8000/queues/works/messages
```

## Add messages to the work queue and drain them.

```
curl -X POST 127.0.0.1:8000/queues/work/messages -d "message 1"
curl -X POST 127.0.0.1:8000/queues/work/messages -d "message 2"
curl -X POST 127.0.0.1:8000/queues/work/messages -d "message 3"
```

```
curl -X POST 127.0.0.1:8000/queues/work/drain
curl 127.0.0.1:8000/stats
```

## Delete the work queue
curl -X DELETE 127.0.0.1:8000/queues/work
