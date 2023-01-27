-- --------------------------------------------------------

--
-- create `channels`
--

CREATE TABLE `channels`
(
    `id`                bigint(20) NOT NULL,
    `creator_user_id`   bigint(20) NOT NULL,
    `access_hash`       bigint(20) NOT NULL,
    `random_id`         bigint(20) NOT NULL,
    `participant_count` int(11) NOT NULL,
    `title`             varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `about`             varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `photo_id`          bigint(20) NOT NULL DEFAULT '0',
    `link`              varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `username`          varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `admins_enabled`    tinyint(1) NOT NULL DEFAULT '0',
    `deactivated`       tinyint(1) NOT NULL DEFAULT '0',
    `version`           int(11) NOT NULL DEFAULT '1',
    `date`              int(11)  NOT NULL DEFAULT '0',
    `created_at`        timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- create `channel_participants`
--

CREATE TABLE `channel_participants`
(
    `id`               bigint(20) NOT NULL,
    `channel_id`       bigint(20) NOT NULL,
    `user_id`          bigint(20) NOT NULL,
    `participant_type` int(11) DEFAULT '0',
    `admin_rights`     json NOT NULL,
    `inviter_user_id`  bigint(20) NOT NULL DEFAULT '0',
    `invited_at`       int(11) NOT NULL DEFAULT '0',
    `kicked_at`        int(11) NOT NULL DEFAULT '0',
    `left_at`          int(11) NOT NULL DEFAULT '0',
    `joined_at`        int(11) NOT NULL DEFAULT '0',
    `state`            tinyint(4) NOT NULL DEFAULT '0',
    `created_at`       timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`       timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Indexes for table `channels`
--
ALTER TABLE `channels`
    ADD PRIMARY KEY (`id`),
    ADD KEY `creator_id` (`creator_user_id`);
ALTER TABLE `channels`
    MODIFY id bigint auto_increment;

--
-- Indexes for table `channel_participants`
--
ALTER TABLE `channel_participants`
    ADD PRIMARY KEY (`id`),
    ADD KEY `chat_id` (`channel_id`),
    ADD KEY `participant_id` (`user_id`);
ALTER TABLE `channel_participants`
    MODIFY id bigint auto_increment;
