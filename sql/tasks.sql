create database gotask;

\c gotask;

create table if not exists tasks(
  uid bigserial primary key,
  description varchar(255) not null,
  user_id varchar(100) not null,
  closed boolean not null default false,
  created_at timestamp default current_timestamp
);
