// Создания индекса уникальности на email автора
CREATE CONSTRAINT unique_email
FOR (u:User)
REQUIRE u.email IS UNIQUE

// Создание автора
CREATE (u:User {
    firstName: "%s",
    lastName: "%s",
    patronymic: "%s",
    job: "%s",
    email: "%s",
    password: "%s"
})

// Создание статьи
CREATE (a:Article {
    title: "%s",
    subject: "%s",
    keywords: "%s",
    annotation: "%s",
    yearOfPublication: %d,
    sourceLink: "%s"
})

// Создание связи между автором и статьей
MATCH (u:User), (a:Article)
WHERE ID(u) = %d AND ID(a) = %d
CREATE (u)-[:WROTE]->(a)

// Поиск автора по id
MATCH (u:User)
WHERE ID(u) = %d
RETURN u

// Поиск автора по email
MATCH (u:User)
WHERE u.email = "%s"
RETURN u

// Поиск статей по названию
MATCH (a:Article)
WHERE a.title CONTAINS "%s"
RETURN article

// Поиск статей по ФИО автора
MATCH (u:User)-[:WROTE]->(a:Article)
WHERE u.firstName CONTAINS "%s"
    OR u.lastName CONTAINS "%s"
    OR u.patronymic CONTAINS "%s"
RETURN a

// Поиск статей статей автора
MATCH (u:User)-[:WROTE]->(a:Article)
WHERE ID(u) = %d
RETURN a

// Поиск статей по ключевым словам
MATCH (a:Article)
WHERE ANY(keyword IN a.keywords WHERE keyword CONTAINS "%s")
RETURN a

// Поиск статей по тематике
MATCH (a:Article)
WHERE a.subject = "%s"
RETURN a

// Поиск статей по году публикации
MATCH (a:Article)
WHERE a.yearOfPublication = %d
RETURN a
