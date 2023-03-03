CREATE TABLE coupon_statuses (
    `id` varchar(255) NOT NULL,
    `used_flg` boolean NOT NULL,
    `updated_at` datetime default current_timestamp,
    `created_at` datetime default current_timestamp,
    `coupon_id` varchar(255) NOT NULL,
    `user_id` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);
