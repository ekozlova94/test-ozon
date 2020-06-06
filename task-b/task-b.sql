-- для составления sql запроса созданы и заполнены тестовыми данными указанные в задании таблицы

create table goods
(
    id   INTEGER,
    name TEXT
);

create table tags
(
    id   INTEGER,
    name TEXT
);

create table tags_goods
(
    tag_id   INTEGER,
    goods_id INTEGER,
    UNIQUE (tag_id, goods_id)
);

insert into goods (id, name)
VALUES (1, 'a'),
       (2, 'b'),
       (3, 'c'),
       (4, 'd');

insert into tags (id, name)
VALUES (1, 't1'),
       (2, 't2'),
       (3, 't3'),
       (4, 't4');

insert into tags_goods (tag_id, goods_id)
VALUES (1, 1),
       (2, 1),
       (3, 1),
       (4, 1),
       (1, 4),
       (2, 4),
       (3, 4),
       (4, 4),
       (1, 2),
       (2, 2),
       (3, 3),
       (4, 3);

select *
from goods as g
-- в выборку попадают те товары, у которых количество собственных тегов совпадает с общим
where
-- выбрано количество тегов принадлежащих товару g
(select count(*) from tags_goods as t where t.goods_id = g.id)
    =
-- выбрано количество тегов всего
(select count(*) from tags);