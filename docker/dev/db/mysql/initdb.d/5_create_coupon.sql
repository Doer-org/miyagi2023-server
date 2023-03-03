CREATE TABLE coupons (
    `id` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `expiration_date` int(10) NOT NULL,
    `discount_rate` int(10) NOT NULL,
    `created_at` datetime default current_timestamp,
    `spot_id` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);
