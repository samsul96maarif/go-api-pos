/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-07-03 21:29:23
 * @modify date 2022-07-03 21:29:23
 * @desc [description]
 */

-- +migrate Up
CREATE TABLE item_prices (
    id serial PRIMARY KEY,
    item_id INT NOT NULL,
    name varchar(255),
    code text,
    price int,
    is_default bool DEFAULT FALSE,
    created_by int,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT item_id_fkey FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE item_prices;