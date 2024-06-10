
create table users
(
    id          serial
        constraint users_id
            primary key,
    "firstName" varchar(100)                        not null,
    "lastName"  varchar(100)                        not null,
    created_at  timestamp default CURRENT_TIMESTAMP not null
);

create table posts
(
    id             serial
        constraint posts_id
            primary key,
    user_id        integer                             not null
        constraint posts_users__id
            references users,
    text           varchar(10000)                      not null,
    created_at     timestamp default CURRENT_TIMESTAMP not null,
    "allowComment" boolean   default true              not null
);


create table comments
(
    id             serial
        constraint id
            primary key,
    parent_comment integer,
    post_id        integer                             not null
        constraint comments_posts_id
            references posts,
    text           varchar(2000)                       not null,
    user_id        integer                             not null
        constraint comments_users_id
            references users,
    created_at     timestamp default CURRENT_TIMESTAMP not null
);





