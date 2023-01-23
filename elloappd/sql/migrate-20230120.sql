-- --------------------------------------------------------

--
-- create `channels`
--

CREATE TABLE IF NOT EXISTS `channels` (
    `id` bigint(20) NOT NULL,
    `creator_user_id` bigint(20) NOT NULL,
    `access_hash` bigint(20) NOT NULL,
    `random_id` bigint(20) NOT NULL,
    `participant_count` int(11) NOT NULL,
    `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `about` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `photo_id` bigint(20) NOT NULL DEFAULT '0',
    `link` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `admins_enabled` tinyint(4) NOT NULL DEFAULT '0',
    `deactivated` tinyint(4) NOT NULL DEFAULT '0',
    `version` int(11) NOT NULL DEFAULT '1',
    `date` int(11) NOT NULL DEFAULT '0',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- create `channel_message_boxes`
--

CREATE TABLE IF NOT EXISTS `channel_message_boxes` (
    `id` bigint(20) NOT NULL,
    `sender_user_id` bigint(20) NOT NULL,
    `channel_id` bigint(20) NOT NULL,
    `channel_message_box_id` bigint(20) NOT NULL,
    `message_id` bigint(20) NOT NULL,
    `date` int(11) NOT NULL DEFAULT '0',
    `deleted` tinyint(4) NOT NULL DEFAULT '0',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- create `channel_participants`
--

CREATE TABLE IF NOT EXISTS `channel_participants` (
    `id` bigint(20) NOT NULL,
    `channel_id` bigint(20) NOT NULL,
    `user_id` bigint(20) NOT NULL,
    `participant_type` tinyint(4) DEFAULT '0',
    `inviter_user_id` bigint(20) NOT NULL DEFAULT '0',
    `invited_at` int(11) NOT NULL DEFAULT '0',
    `joined_at` int(11) NOT NULL DEFAULT '0',
    `state` tinyint(4) NOT NULL DEFAULT '0',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- create `channel_pts_updates`
--

CREATE TABLE IF NOT EXISTS `channel_pts_updates` (
    `id` bigint(20) NOT NULL,
    `channel_id` bigint(20) NOT NULL,
    `pts` int(11) NOT NULL,
    `pts_count` int(11) NOT NULL,
    `update_type` tinyint(4) NOT NULL DEFAULT '0',
    `update_data` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
    `date2` int(11) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------


--
-- Indexes for table `channels`
--
ALTER TABLE `channels`
    ADD PRIMARY KEY (`id`),
    ADD KEY `creator_user_id_3` (`creator_user_id`,`access_hash`);

--
-- Indexes for table `channel_message_boxes`
--
ALTER TABLE `channel_message_boxes`
    ADD PRIMARY KEY (`id`);

--
-- Indexes for table `channel_participants`
--
ALTER TABLE `channel_participants`
    ADD PRIMARY KEY (`id`),
    ADD KEY `chat_id` (`channel_id`);

--
-- Indexes for table `channel_pts_updates`
--
ALTER TABLE `channel_pts_updates`
    ADD PRIMARY KEY (`id`);
