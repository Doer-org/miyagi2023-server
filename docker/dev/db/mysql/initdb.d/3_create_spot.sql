CREATE TABLE spots (
    `id` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `detail` varchar(255) NOT NULL,
    `like` int(10) NOT NULL,
    `img_url` varchar(255) NOT NULL,
    `address` varchar(255) NOT NULL,
    `created_at` datetime default current_timestamp,
    PRIMARY KEY (`id`)
);
