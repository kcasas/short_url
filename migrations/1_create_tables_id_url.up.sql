CREATE TABLE `ids` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE `urls` (
  `short` varchar(20) NOT NULL DEFAULT '',
  `long_url` varchar(2083) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `exp_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`short`),
  KEY `exp_at` (`exp_at`)
) ENGINE=InnoDB;
