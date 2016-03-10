CREATE DATABASE tablespacedb;
USE tablespacedb;

CREATE TABLE fooditem{
	id INT NOT NULL PRIMARY KEY,
	name varchar(50) NOT NULL,
	price money NOT NULL,
	discount money,
	description varchar(5000),
	isactive boolean NOT NULL,
	omnivoreid varchar(100),
	itempicture varchar(1000)
};

CREATE TABLE tag{
	id INT NOT NULL PRIMARY KEY,
	description varchar(1000),
	name varchar(100) NOT NULL
};

CREATE TABLE foodtags{
	foodid INT NOT NULL,
	tagid INT NOT NULL,
	CONSTRAINT pk_foodTag PRIMARY KEY (foodid,tagid)
};

CREATE TABLE restaurant{
	id INT NOT NULL PRIMARY KEY,
	name varchar(500) NOT NULL, 
	logo varchar(1000),
	address varchar(500),
	phonenumber varchar(20),
	email varchar(50) NOT NULL,
	hours varchar (500)
};

CREATE TABLE menu{
	id INT NOT NULL PRIMARY KEY,
	isactive boolean
};

CREATE TABLE menuitems{
	menuid INT NOT NULL,
	foodid INT NOT NULL,
	CONSTRAINT pk_menuitem PRIMARY KEY (menuid, foodid)
};

CREATE TABLE restaurantmenus{
	restaurantid INT NOT NULL,
	menuid INT NOT NULL,
	CONSTRAINT pk_restaurantmenu PRIMARY KEY (restaurantid, menuid)
};

CREATE TABLE users{
	id INT NOT NULL PRIMARY KEY,
	joined date NOT NULL,
	lastactive date NOT NULL,
	username varchar(100) NOT NULL,
	firstname varchar(50),
	lastname varchar(50),
	address varchar (100),
	email varchar(500) NOT NULL
};

CREATE TABLE passwords{
	userid INT NOT NULL PRIMARY KEY,
	hash varchar(50) NOT NULL,
	salt varchar(50) NOT NULL
};

CREATE TABLE restrictions{
	userid INT NOT NULL,
	tagid INT NOT NULL,
	CONSTRAINT pk_restriction PRIMARY KEY (userid, tagid)
};

CREATE TABLE favourites{
	userid INT NOT NULL,
	foodid INT NOT NULL,
	CONSTRAINT pk_favourite PRIMARY KEY (userid, foodid)
};

CREATE TABLE accesslevels{
	accessleveid INT NOT NULL PRIMARY KEY,
	description varchar(100),
	name varchar(50) NOT NULL
};

CREATE TABLE permissions{
	userid INT NOT NULL,
	restaurantid INT NOT NULL,
	accesslevelid INT NOT NULL,
	CONSTRAINT pk_permission PRIMARY KEY (userid, restaurantid)
};

CREATE TABLE paymenttypes{
	id INT NOT NULL PRIMARY KEY,
	description varchar(1000),
	name varchar(50) NOT NULL
};

CREATE TABLE userpaymenttypes{
	userid INT NOT NULL,
	typeid INT NOT NULL,
	CONSTRAINT pk_userpaymenttype PRIMARY KEY (userid, typeid)
};

CREATE TABLE transactions{
	id INT NOT NULL PRIMARY KEY,
	paymenttypeid INT NOT NULL,
	totalprice money NOT NULL,
	transactiondate date NOT NULL	
};

CREATE TABLE orders{
	id INT NOT NULL PRIMARY KEY,
	creationtime date NOT NULL,
	updatetime date,
	statusid INT NOT NULL,
	restaurantid INT NOT NULL,
	total money NOT NULL,
};

CREATE TABLE ordertransactions{
	orderid INT NOT NULL,
	transactionid INT NOT NULL,
	CONSTRAINT pk_ordertransaction PRIMARY KEY(orderid, transactionid)
};

CREATE TABLE statuses{
	id INT NOT NULL PRIMARY KEY,
	name varchar(50) NOT NULL
};

CREATE TABLE ordercontents{
	orderid INT NOT NULL,
	foodid INT NOT NULL,
	CONSTRAINT pk_ordercontents PRIMARY KEY (orderid, foodid)
};