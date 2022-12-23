create table user_feeds
(
    id      bigint unsigned auto_increment
        primary key,
    user_id bigint not null,
    chat_id bigint not null
);

create unique index user_feeds_user_id_chat_id_index
    on user_feeds (user_id, chat_id);

