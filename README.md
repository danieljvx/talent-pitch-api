# talent-pitch-api @danieljvx

### `1 - Up conntainer Database:`
```bash
docker-compose up db -d --build
```

### `2 - Up conntainer Api:`
```bash
docker-compose up api -d --build
```

### `3 - Generate Doc Swagger:`
```bash
docker-compose exec api bash -c "swag init -g main.go --output docs"
```

### `4 - Import database:`
```bash
docker-compose exec db bash -c "mariadb -uroot -p123qwe talentpitch < /tmp/talentpitch.sql"
```

### `Local talent-pitch-api:` [http://localhost:3000](http://localhost:3000)
| Resource            | Type | Path                                                                                       | Body |
|---------------------|------|--------------------------------------------------------------------------------------------| ------ |
| Swagger             | GET  | [http://localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html)       | null |
| Base                | GET  | [http://localhost:3000](http://localhost:3000)                                             | null |
| GPT Data Migration  | GET  | [http://localhost:3000/gpt/migration](http://localhost:3000/gpt/migration)                 | null |
| User Get by Id      | GET  | [http://localhost:3000/user/:id](http://localhost:3000/user/:id)                           | null |
| User Get List       | GET  | [http://localhost:3000/users?page=0&perPage5](http://localhost:3000/users?page=0&perPage5) | null |
| User Create         | POST | [http://localhost:3000/user](http://localhost:3000/user)                                   | `{ "name": "Monica","email": "moniquin@gmail.com",	"image": ""}` |
| User Update         | PUT  | [http://localhost:3000/user/:id](http://localhost:3000/user/:id)                           | `{ "name": "Monica","email": "moniquin@gmail.com",	"image": ""}` |
| Program Get by Id   | GET  | [http://localhost:3000/program/:id](http://localhost:3000/program/:id)                        | null |
| Program Get List    | GET  | [http://localhost:3000/programs?page=0&perPage5](http://localhost:3000/programs?page=0&perPage5) | null |
| Program Create      | POST | [http://localhost:3000/program](http://localhost:3000/program)                                   | `{"title": "Program 3", "description": "program description", "start_date": "2024-06-12", "end_date": "2024-07-12", "user_id": 1 }` |
| Program Update      | PUT  | [http://localhost:3000/program/:id](http://localhost:3000/program/:id)                           | `{"title": "Program 3", "description": "program description", "start_date": "2024-06-12", "end_date": "2024-07-12", "user_id": 1 }` |
| Program Participant | PUT  | [http://localhost:3000/program/1/participant](http://localhost:3000/program/1/participant)                           | `{"program_id": 2, "challenge_id": 2, "company_id": 2 }` |
| Challenge Get by Id   | GET  | [http://localhost:3000/challenge/:id](http://localhost:3000/challenge/:id)                        | null |
| Challenge Get List  | GET  | [http://localhost:3000/challenges?page=0&perPage5](http://localhost:3000/challenges?page=0&perPage5) | null |
| Challenge Create      | POST | [http://localhost:3000/challenge](http://localhost:3000/challenge)                                   | `{"title": "Challenge 3", "description": "challenge description", "difficulty": 1, "user_id": 1 }` |
| Challenge Update      | PUT  | [http://localhost:3000/challenge/:id](http://localhost:3000/challenge/:id)                           | `{"title": "Challenge 3", "description": "challenge description", "difficulty": 1, "user_id": 1 }` |
| Company Get by Id   | GET  | [http://localhost:3000/company/:id](http://localhost:3000/company/:id)                        | null |
| Company Get List  | GET  | [http://localhost:3000/companies?page=0&perPage5](http://localhost:3000/companies?page=0&perPage5) | null |
| Company Create      | POST | [http://localhost:3000/company](http://localhost:3000/company)                                   | `{"title": "Challenge 3", "description": "challenge description", "difficulty": 1, "user_id": 1 }` |
| Company Update      | PUT  | [http://localhost:3000/company/:id](http://localhost:3000/company/:id)                           | `{"title": "Challenge 3", "description": "challenge description", "difficulty": 1, "user_id": 1 }` |


### `Test App:`
```bash
go test -v ./... -cover
```

## Daniel Villanueva

Email: [villanueva.danielx@gmail.com](mail://villanueva.danielx@gmail.com)

Github: [@danieljvx](https://github.com/danieljvx)
