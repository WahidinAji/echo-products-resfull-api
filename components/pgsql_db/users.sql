

CREATE TABLE users(
    id uuid not null,
    first_name varchar(100) not null,
    last_name varchar(100) not null,
    email varchar(50) not null unique ,
    phone_number varchar(14),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

insert into users(id, first_name, last_name, email, phone_number) VALUES
    ('73e12106-207e-4693-9c0d-3147d6ab606a','Wahidin','Aji','a17wahidin@gmail.com','123456789012'),
    ('44c22cb3-ff6c-4043-8c79-8a5506ce11e9','Tia','Ulul Putri','tiaulul.putri@mail.com','123123123123'),
    ('d90f8110-039a-47f4-a164-37d807f77ab5','Omoy','Bungsu Putri','omoy.putri@mail.com','987654321098');

select  * from users;
drop table users;

SELECT EXISTS (SELECT id FROM users WHERE id='73e12106-207e-4693-9c0d-3147d6ab606a');

select * from users where id='d90f8110-039a-47f4-a164-37d807f77ab5';
