drop table if exists users cascade;

create table users (
	user_id serial not null primary key,
	full_name character varying(64) not null,
	user_name character varying(32) not null,
	phone_number character varying(9) not null,
	created_at timestamptz default current_timestamp 
);

--MOCK DATA

insert into 
	users (full_name, user_name, phone_number) 
values ('Ali Aliyev', 'aziz', '954566545'),
		('Vali Valiyev', 'vali', '954116545'),
		('G''ani G''aniyev', 'g''ani', '952266545');