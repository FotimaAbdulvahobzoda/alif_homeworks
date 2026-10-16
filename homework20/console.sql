CREATE TABLE IF NOT EXISTS books(
    id SERIAL PRIMARY KEY,
    title TEXT,
    author TEXT,
    price INT,
    pages int,
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO books (title, author, price, pages) VALUES
    ('Harry Potter: The Order of the Phoenix', 'J.K.Rolling', 128, 837),
    ('Diary of Whimpy Kid', 'Jeff Keffley', 99, 157),
    ('No Longer Human', 'Osamu Dazai', 199, 170),
    ('Rashomon', 'Ryunosuke Akutagawa', 200, 320),
    ('Crime and Punishment', 'Fyodor Dostoevski', 152, 167),
    ('Iron Gold', 'Pierce Brown', 217, 601),
    ('Love without limits', 'Nick and Kanae Vujicic', 205, 347),
    ('The Financier', 'Theodor Dreiser', 164, 736),
    ('The Girl You Left Behind', 'Jojo Moyes', 151, 456),
    ('Death Note Volume 8', 'Tsugumi Ohba and Takeshi Obata', 7, 220);

select * from books;

select distinct title, author from books;

select distinct title from books where id=5;

update books set price = 350 where id=3;

update books set price = price + 100;

update books set pages = 500 where id=6;

select distinct title, pages from books;

delete from books where id=8;

select distinct title from books;

delete from books where price < 250;

select distinct title, price from books;

select distinct title, price from books where price > 400;

select distinct title, pages from books where pages > 300;

select distinct title, price from books order by price desc;

select title from books order by title limit 5;