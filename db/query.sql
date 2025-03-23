-- name: CreateRedWord :execresult
insert into red_word (name) values ($1);

-- name: CreateGreenWord :execresult
insert into green_word (name) values ($1);

