create table news_article.articles (
  id int unique not null auto_increment,
  title varchar(256),
  description text,
  content text,
  article_url varchar(1024),
  image_url varchar(1024),
  resource_name varchar(256),
  created_at datetime,
  updated_at datetime
);
