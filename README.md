# Blogpost
## It consists of commits that show step by step evolution of a REST API project in Golang.
### For my students and for others to learning purpose

Migrate DB:
<br>
```migrate -path ./storage/migrations -database 'postgres://admin:qwerty123@localhost:5432/article_service_db?sslmode=disable' up```