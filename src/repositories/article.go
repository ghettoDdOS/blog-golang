package repositories

import (
	"blog/src/config"
	"blog/src/models"
	"blog/src/utils"
	"fmt"

	"github.com/Epritka/morpheus/builder/entity"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

type ArticleRepository struct {
	ctx   *gin.Context
	alias string
}

func NewArticleRepository(ctx *gin.Context) *ArticleRepository {
	return &ArticleRepository{ctx: ctx, alias: "a"}
}

func (r *ArticleRepository) Create(article *models.Article, userId int64) error {
	db := config.GetDatabaseConnection()
	conn := db.NewExecuterWithContext(r.ctx.Request.Context())

	query := `CREATE (%s:Article {
    title: "%s",
    subject: "%s",
    keywords: [%s],
    annotation: "%s",
    yearOfPublication: %d,
    sourceLink: "%s"
	})
	WITH %s
	MATCH (u:User)
	WHERE ID(u) = %d
	CREATE (u)-[:WROTE]->(%s)`

	cypher := fmt.Sprintf(query,
		r.alias,
		article.Title,
		article.Subject,
		utils.ListToCypher(article.Keywords),
		article.Annotation,
		article.YearOfPublication,
		article.SourceLink,
		r.alias,
		userId,
		r.alias,
	)

	_, err := conn.DoQuery(cypher)

	conn.CloseWithContext(r.ctx.Request.Context())
	return err
}

func (r *ArticleRepository) parseRecords(records []*db.Record) []*models.Article {
	var articles []*models.Article
	for _, record := range records {
		node, found := record.Get(r.alias)
		if !found {
			continue
		}
		articleData := node.(dbtype.Node).Props
		keywords := make([]string, len(articleData["keywords"].([]interface{})))
		for i, keyword := range articleData["keywords"].([]interface{}) {
			keywords[i] = keyword.(string)
		}
		article := models.Article{
			BaseModel:         models.BaseModel{Id: node.(dbtype.Node).Id},
			Title:             articleData["title"].(string),
			Subject:           articleData["subject"].(string),
			Keywords:          keywords,
			Annotation:        articleData["annotation"].(string),
			YearOfPublication: articleData["yearOfPublication"].(int64),
			SourceLink:        articleData["sourceLink"].(string),
		}
		articles = append(articles, &article)
	}
	return articles
}

func (r *ArticleRepository) getBy(query string) ([]*models.Article, error) {
	db := config.GetDatabaseConnection()
	conn := db.NewExecuterWithContext(r.ctx.Request.Context())

	conn.
		Match(entity.NewNode(r.alias).SetLables("Article")).
		Where(query).
		Return(entity.NewAlias(r.alias))

	records, err := conn.Do()
	if err != nil {
		return nil, err
	}
	articles := r.parseRecords(records)
	conn.CloseWithContext(r.ctx.Request.Context())
	return articles, nil
}

func (r *ArticleRepository) GetByTitle(title string) ([]*models.Article, error) {
	return r.getBy(fmt.Sprintf(`%s.title CONTAINS "%s"`, r.alias, title))
}

func (r *ArticleRepository) GetBySubject(subject string) ([]*models.Article, error) {
	return r.getBy(fmt.Sprintf(`%s.subject CONTAINS "%s"`, r.alias, subject))
}

func (r *ArticleRepository) GetByYearOfPublication(yearOfPublication int) ([]*models.Article, error) {
	return r.getBy(fmt.Sprintf("%s.yearOfPublication = %d", r.alias, yearOfPublication))
}

func (r *ArticleRepository) GetByFullName(fullName string) ([]*models.Article, error) {
	db := config.GetDatabaseConnection()
	conn := db.NewExecuterWithContext(r.ctx.Request.Context())

	firstNameFilter := fmt.Sprintf(`"%s" CONTAINS u.firstName`, fullName)
	lastNameFilter := fmt.Sprintf(`"%s" CONTAINS u.lastName`, fullName)
	patronymicFilter := fmt.Sprintf(`"%s" CONTAINS u.patronymic`, fullName)

	conn.
		Match(entity.
			NewPatternList(entity.NewNode("u").SetLables("User")).
			Related(entity.NewRelationship("").SetLables("WROTE")).
			To(entity.NewNode(r.alias).SetLables("Article")),
		).
		Where(fmt.Sprintf("%s OR %s OR %s", firstNameFilter, lastNameFilter, patronymicFilter)).
		Return(entity.NewAlias(r.alias))

	records, err := conn.Do()
	if err != nil {
		return nil, err
	}
	articles := r.parseRecords(records)
	conn.CloseWithContext(r.ctx.Request.Context())
	return articles, nil
}
func (r *ArticleRepository) GetByUserId(userId int64) ([]*models.Article, error) {
	db := config.GetDatabaseConnection()
	conn := db.NewExecuterWithContext(r.ctx.Request.Context())

	query := `MATCH (u:User)-[:WROTE]->(%s:Article)
	WHERE ID(u) = %d
	RETURN %s`

	cypher := fmt.Sprintf(query, r.alias, userId, r.alias)

	records, err := conn.DoQuery(cypher)
	if err != nil {
		return nil, err
	}

	articles := r.parseRecords(records)
	conn.CloseWithContext(r.ctx.Request.Context())
	return articles, nil
}

func (r *ArticleRepository) GetByKeywords(keywords []string) ([]*models.Article, error) {
	db := config.GetDatabaseConnection()
	conn := db.NewExecuterWithContext(r.ctx.Request.Context())

	query := `MATCH (%s:Article)
	WHERE ANY(keyword IN [%s] WHERE keyword IN %s.keywords)
	RETURN %s`

	cypher := fmt.Sprintf(query, r.alias, utils.ListToCypher(keywords), r.alias, r.alias)

	records, err := conn.DoQuery(cypher)
	if err != nil {
		return nil, err
	}

	articles := r.parseRecords(records)
	conn.CloseWithContext(r.ctx.Request.Context())
	return articles, nil
}

func (r *ArticleRepository) Get() ([]*models.Article, error) {
	db := config.GetDatabaseConnection()
	conn := db.NewExecuterWithContext(r.ctx.Request.Context())
	conn.
		Match(entity.NewNode(r.alias).SetLables("Article")).
		Return(entity.NewAlias(r.alias))

	records, err := conn.Do()
	if err != nil {
		return nil, err
	}
	articles := r.parseRecords(records)
	conn.CloseWithContext(r.ctx.Request.Context())
	return articles, nil
}

func (r *ArticleRepository) GetAviableKeywords() ([]string, error) {
	db := config.GetDatabaseConnection()
	conn := db.NewExecuterWithContext(r.ctx.Request.Context())

	query := `MATCH (%s:Article)
	RETURN REDUCE(acc = [], value IN COLLECT(DISTINCT %s.keywords) | acc + value)
	AS keywords`

	cypher := fmt.Sprintf(query, r.alias, r.alias)

	records, err := conn.DoQuery(cypher)
	if err != nil {
		return nil, err
	}

	if len(records) < 1 {
		return nil, nil
	}

	node, found := records[0].Get("keywords")
	if !found {
		return nil, nil
	}
	keywords := make([]string, len(node.([]interface{})))
	for i, value := range node.([]interface{}) {
		keywords[i] = value.(string)
	}

	conn.CloseWithContext(r.ctx.Request.Context())
	return keywords, nil
}

func (r *ArticleRepository) DeleteById(id int) error {
	db := config.GetDatabaseConnection()
	conn := db.NewExecuterWithContext(r.ctx.Request.Context())

	conn.
		Match(entity.NewNode(r.alias).SetLables("Article")).
		Where(fmt.Sprintf("ID(%s) = %d", r.alias, id)).
		DetachDelete(entity.NewAlias(r.alias))

	_, err := conn.Do()
	if err != nil {
		return err
	}

	conn.CloseWithContext(r.ctx.Request.Context())
	return nil
}
