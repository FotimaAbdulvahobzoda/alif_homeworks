CREATE TABLE IF NOT EXISTS books(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    price INT NOT NULL,
    pages INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
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
    ('Death Note Volume 8', 'Tsugumi Ohba and Takeshi Obata', 7, 220),
    ('Project Hail Mary', 'Andy Weir', 101, 340),
    ('Pride and Prejudice', 'Jane Austen', 80, 296),
    ('The Housemaid', 'Freida McFadden', 56, 456),
    ('To Kill a Mockingbird', 'Harper Lee', 59, 234),
    ('Fourth Wing', 'Rebecca Yarros', 93, 210);

SELECT DISTINCT * FROM books;


select title, author from books;

select price, pages from books;

select title, author from books where author='Robert Martin';

SELECT title, author  FROM books WHERE price > 100;

select distinct title, price from books where price < 50;

select distinct title, pages from books where pages > 500;

select distinct title, price from books where price between 5 and 20;

select distinct title, price, pages from books where price > 100 and pages > 300;

select distinct title, price, pages from books where price < 100 or pages < 200;

select distinct title, author, price from books where author='Theodor Dreiser' and price > 150;

select distinct title, price from books order by price desc;

select distinct title, price from books order by price ASC;

select distinct title, pages from books order by pages desc;

select distinct title from books order by title;

select distinct title, price from books order by price limit 5 ;

select distinct title, pages from books order by pages limit 3;

update books set price =+ 100 where price < 100;

delete from books where id=1;

select distinct * from books;

delete from books where price < 50;

delete from books where pages < 100;

select distinct title, price, pages from books;

select distinct count(*) from books ;

select distinct count(*) from books where price > 100;

select distinct count(*) from books where author='Robert Martin';

select  distinct title from books where title like 'Go%';

select distinct title, author from books where author like '%Dazai';

select distinct title from books where price < 100;

update books set price = 250 where price < 100;

select distinct title, price from books;

select distinct title, price from books limit 5;


















