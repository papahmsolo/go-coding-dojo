
DROP TABLE IF EXISTS `employees`;

CREATE TABLE `employees` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` text DEFAULT NULL,
  `email` text DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `employees` (`id`, `name`, `email`)
VALUES
	(1,'George B','gb@email.com'),
	(2,'Johny B','john@email.com'),
	(3,'Oleg Pupkin','obleg@email.com'),
	(4,'Bohdan','b@email.com');


DROP TABLE IF EXISTS `products`;

CREATE TABLE `products` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `category` text NOT NULL,
  `active` tinyint(1) NOT NULL DEFAULT 1,
  `on_sale` tinyint(1) NOT NULL DEFAULT 0,
  `empl_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `products` (`id`, `name`, `category`, `active`, `on_sale`, `empl_id`)
VALUES
	(2,'amstel','BEER',1,0,1),
	(3,'lidskae','BEER',1,1,1),
	(4,'zewa','PAPER',1,0,1),
	(5,'nazhdacka','PAPER',1,0,2),
	(6,'princessa nuri','TEA',1,0,1),
	(7,'slon','TEA',1,0,2),
	(8,'karkade','TEA',1,1,2),
	(9,'zolotaya chashka','TEA',0,0,3),
	(10,'domik v sele','MILK',1,0,3),
	(11,'prostokvashino','MILK',1,0,1),
	(12,'minskya charka','MILK',1,0,2),
	(13,'snowskae','MILK',1,0,2),
	(14,'kleckae','MILK',1,1,3),
	(15,'kavun','FRUIT',1,0,1);
