# blog-server
Blog Api

## Usage

```bash
go install

go build
#or
go run main.go
```

URL: `localhost:8080`

**getAllPost**

```
method: GET
path:   /get_all
```

**getPublishPost**

```
method: GET
path:   /get_publish
```

**getPrivatePost**

```
method: GET
path:   /get_private
```

**getPostById**

```
method: GET
path:   /post?id=1
```

**Insert**

```
method: POST
path:   /create
params: {
    title:      'this is title',
    content:    'this is content',
    summary:    'this is summary',
    publish:    1
}
```

**Update**

```
method: POST
path:   /update
params: {
    id:         1,
    title:      'title changed',
    content:    'this content is changed too',
    summary:    'this is summary',
    publish:    0
}
```

**Delete**

```
method: POST
path:   /delete
params: {
    id:         1
}
```