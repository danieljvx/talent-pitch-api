# talent-pitch-api @danieljvx

### `Deploy talent-pitch-api:` [https://danieljvx-meli-morse.herokuapp.com](https://danieljvx-meli-morse.herokuapp.com)
| Resource | Type | Path | Body |
| ------ | ------ | ------ | ------ |
| Swagger | GET | [https://danieljvx-meli-morse.herokuapp.com/swagger/index.html](https://danieljvx-meli-morse.herokuapp.com/swagger/index.html) | null |

### `Up conntainer:`
```bash
docker-compose up -d --build
```

### `Import database:`
```bash
docker-compose exec db bash -c "mariadb -uroot -p123qwe talentpitch < /tmp/talentpitch.sql"
```

### `Local talent-pitch-api:` [http://localhost:3000](http://localhost:3000)
| Resource | Type | Path | Body |
| ------ | ------ | ------ | ------ |
| Swagger | GET | [http://localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html) | null |

### `Test App:`
```bash
go test -v ./... -cover
```

## Daniel Villanueva

Email: [villanueva.danielx@gmail.com](mail://villanueva.danielx@gmail.com)

Github: [@danieljvx](https://github.com/danieljvx)
