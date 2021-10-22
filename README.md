
```bash

docker run --rm --name binchbank -v $(pwd)/_pg:/var/lib/postgresql/data -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:13

```

