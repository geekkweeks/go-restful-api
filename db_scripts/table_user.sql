CREATE TABLE user (
  id int not null primary key auto_increment,
  username varchar(200) not null,
  firstName varchar(200) not null,
  lastName varchar(200) not null,
  phone varchar(50) not null

) engine = innodb