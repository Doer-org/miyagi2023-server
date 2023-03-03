CREATE TABLE users (
    `id` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `age` int(10) NOT NULL,
    `gender` varchar(255) NOT NULL,
    `birthday` datetime NOT NULL,
    `address` varchar(255) NOT NULL,
    `profile_img` varchar(255),
    `prefecture` varchar(255) NOT NULL,
    `created_at` datetime default current_timestamp,
    PRIMARY KEY (`id`)
);
