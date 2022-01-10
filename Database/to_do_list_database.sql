create database to_do_lists;
create database to_do_lists_test;

use to_do_lists;

create table Users(
idUser int auto_increment,
name varchar(20),
password varchar(50),

primary key(idUser)
);

Create table Projects(
projectid int auto_increment,
idUser int,
proejctname varchar(50),

primary key(projectid),
foreign key(idUser) references Users(idUser)
);

create table to_do(
idto_do int auto_increment,
projectid int,
task varchar(30),
description varchar(200),

primary key(idto_do),
foreign key(projectid) references Projects(projectid)
);
