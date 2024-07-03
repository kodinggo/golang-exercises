
-- +migrate Up
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `fullname` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `biodata` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
);

-- +migrate Down
DROP TABLE IF EXISTS `users`;
