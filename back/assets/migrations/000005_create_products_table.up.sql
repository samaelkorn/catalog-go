CREATE TABLE products (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    price INTEGER NOT NULL,
    color_id INTEGER NOT NULL,
    status_id INTEGER NOT NULL,
    FOREIGN KEY (color_id) REFERENCES colors(id),
    FOREIGN KEY (status_id) REFERENCES statuses(id)
);