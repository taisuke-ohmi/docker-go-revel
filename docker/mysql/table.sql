CREATE TABLE revel.message(
  id int NOT NULL AUTO_INCREMENT,
  content text,
  created_at datetime not null,
  updated_at datetime not null,
  PRIMARY KEY (id)
);
