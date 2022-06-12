# Estoque de Roupas

CRUD de um estoque de roupas em go para a disciplina de desenvolvimento WEB

## Estrutura de diretórios

```
estoque_roupas/
   ├─ controllers/
   |   └─ products.go
   ├─ db/
   |   └─ db.go
   ├─ models/
   |   └─ products.go
   ├─ routes/
   |   └─ routes.go
   ├─ templates/
   |   ├─ _head.html
   |   ├─ _menu.html
   |   ├─ edit.html 
   |   ├─ index.html
   |   └─ new.html   
   ├─ go.mod
   ├─ go.sum
   ├─ main.go
   └─ database.sql
 
```

- :file_folder: [controllers/](controllers): Dir que contém os endpoints da API
- :file_folder: [db/](db) Dir que contém os recursos para realizar a conexão com o banco de dados
- :file_folder: [models/](models) Dir que contém os recursos para interagir com o banco de dados
- :file_folder: [routes/](routes) Dir que contém os recursos para as rotas
- :file_folder: [templates/](templates) Dir que contém os arquivos em html 
- :page_with_curl: [database.sql](database.sql): Arquivo com o banco de dados.

## Imagens 


### Home
![Home](https://imgur.com/edvhPug)


