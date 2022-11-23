/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-07-03 21:29:23
 * @modify date 2022-07-03 21:29:23
 * @desc [description]
 */

-- +migrate Up
CREATE TABLE items (
    id serial PRIMARY KEY,
    name varchar(255),
    description text,
    qty int,
    created_by int,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE items;