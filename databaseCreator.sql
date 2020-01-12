create table users(
    username varchar(256) primary key,
    firstname varchar(256) not null,
    lastname varchar(256) not null,
    email varchar(256) not null,
    phone varchar(20) not null,
    address varchar(256) not null,
    password varchar(256) not null,
    imagepath text default 'defaultuser.jpg',
    account numeric default 200000
);
create table admin(
    username varchar(256) primary key,
    firstname varchar(256) not null,
    lastname varchar(256) not null,
    email varchar(256) not null,
    password varchar(256) not null,
    imagepath text default 'defaultadmin.jpg'
);
create table authors (
    id  integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name varchar(256) not null,
    email varchar(256) not null,  
    phone varchar(256) not null,
    address varchar(256) not null,
    description varchar,
    rating int default 0,
    imagepath text default 'defaultcompany.jpg',
    password varchar(256) not null,
    account numeric default 200000,
    activated bool default false
);

create table comments (
    id  integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    username varchar references users(username) on delete cascade,
    cid integer references companies(id) on delete cascade,  
    messages varchar(256) not null,
    placedat varchar(256) not null
);


create table admin (
    username varchar(256) primary key,
    firstname varchar(256) not null,
    lastname varchar(256) not null,
    email varchar(256) not null,
    password varchar(1000) not null,
);

create table books(
    id  integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name varchar(256) not null,
    owner int references authors(id) on delete cascade,
    priceperday numeric default 0,
    ondiscount boolean default false,
    discount numeric default 0,
    onsale boolean default false,
    imagepath text default 'defaultmaterial.jpg'
);