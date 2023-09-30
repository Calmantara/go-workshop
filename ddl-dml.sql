-- membuat table category
create table categories (
	id serial primary key,
	name varchar(200) not null,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp
);

alter table categories add column description text;

insert into categories(name, description, created_at, updated_at)
values ('ELECTRONIC', 'electronic stuff', now(), now());

insert into categories(name, description, created_at, updated_at)
values ('SCHOOL', 'school stuff', now(), now()),
	   ('WORK', 'work stuff', now(), now());
	   
select * from categories;

-- update
update categories
  set description = 'amazing school stuff'
where name = 'SCHOOL';
select * from categories;

-- delete
delete from categories
where name = 'WORK';
select * from categories;

-- select with filter
select * from categories
where name = 'ELECTRONIC';

select * from categories
where lower(name) = lower('electronic');