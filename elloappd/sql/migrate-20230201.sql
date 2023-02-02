-- --------------------------------------------------------

--
-- add columns for `users_ello`
--

ALTER TABLE `users_ello` ADD COLUMN `type` ENUM('personal', 'business') NOT NULL;
ALTER TABLE `users_ello` ADD COLUMN `kind` ENUM('public', 'private') NOT NULL;
ALTER TABLE `users_ello` MODIFY COLUMN `date_of_birth` DATETIME;

