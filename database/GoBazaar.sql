DROP TABLE IF EXISTS user;

CREATE TABLE user (
  id         INT AUTO_INCREMENT NOT NULL,
  first_name      VARCHAR(128) NOT NULL,
  last_name     VARCHAR(128) NOT NULL,
  email     VARCHAR(255) NOT NULL,
  contact VARCHAR(255) NOT NULL,
  city VARCHAR(128) NOT NULL,

  wallet_balance  DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS userCreds;

CREATE TABLE userCreds (
  id         INT NOT NULL,
  pass      VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS merchant;

CREATE TABLE merchant (
  id         INT AUTO_INCREMENT NOT NULL,
  company_name      VARCHAR(128) NOT NULL,
  email     VARCHAR(255) NOT NULL,
  merchant_address VARCHAR(255) NOT NULL,
  discount_offered DECIMAL(5,2) NOT NULL, 
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS merchCreds;
CREATE TABLE merchCreds (
  id         INT NOT NULL,
  pass      VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`)
);


DROP TABLE IF EXISTS product;

CREATE TABLE product (
  id         INT AUTO_INCREMENT NOT NULL,
  merchantID INT NOT NULL,
  namel     VARCHAR(255) NOT NULL,
  product_description VARCHAR(255) NOT NULL,
  price DECIMAL(5,2) NOT NULL, 
  stock INT NOT NULL,
  PRIMARY KEY (`id`)
);
