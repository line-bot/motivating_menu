SET SESSION FOREIGN_KEY_CHECKS=0;

/* Drop Tables */

DROP TABLE IF EXISTS recipe;
DROP TABLE IF EXISTS category;

/* Create Tables */

CREATE TABLE category
(
	id bigint unsigned NOT NULL AUTO_INCREMENT,
	name varchar(30) NOT NULL,
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at timestamp NULL DEFAULT NULL
	PRIMARY KEY (id)
);


CREATE TABLE recipe
(
	id bigint unsigned NOT NULL AUTO_INCREMENT,
	category_id bigint unsigned NOT NULL,
	name varchar(30) NOT NULL,
	count1 bigint unsigned DEFAULT 0 NOT NULL,
	count2 bigint unsigned DEFAULT 0 NOT NULL,
	count3 bigint unsigned DEFAULT 0 NOT NULL,
	count4 bigint unsigned DEFAULT 0 NOT NULL,
	count5 bigint unsigned DEFAULT 0 NOT NULL,
	count6 bigint unsigned NOT NULL,
	count7 bigint unsigned DEFAULT 0 NOT NULL,
	count8 bigint unsigned DEFAULT 0 NOT NULL,
	count9 bigint unsigned DEFAULT 0 NOT NULL,
	count10 bigint unsigned DEFAULT 0 NOT NULL,
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at timestamp NULL DEFAULT NULL
	PRIMARY KEY (id)
);



/* Create Foreign Keys */

ALTER TABLE recipe
	ADD FOREIGN KEY (category_id)
	REFERENCES category (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;



