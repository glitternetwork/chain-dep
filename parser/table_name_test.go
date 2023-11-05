package parser

import (
	"context"
	"fmt"
	"github.com/glitternetwork/chain-dep/parser/analyzer"
	"testing"
)

func Test_parseSql(t *testing.T) {
	//sql1 := "create database otherpay"
	sql1 := "CREATE TABLE IF NOT EXISTS index3.mirrorentry (\n        _id VARCHAR(255) PRIMARY KEY ,        title VARCHAR(2000) COMMENT 'article title',        title2 VARCHAR(2000) COMMENT 'article title',        body VARCHAR(1000000) COMMENT 'article body',        body2 VARCHAR(1000000) COMMENT 'article body',        cover_url VARCHAR(516) COMMENT 'article cover image url',        article_link VARCHAR(516) COMMENT 'article link',        language VARCHAR(32) COMMENT 'article language:en or zh',        published_time INT(11) COMMENT 'article published time',        category VARCHAR(1024) COMMENT 'article category',        author VARCHAR(255) COMMENT 'author ethereum address',        display_name VARCHAR(128) COMMENT 'author nickname',        avatar_url  VARCHAR(255) COMMENT 'author avatar image url',         ens VARCHAR(128) COMMENT 'ethereum name server',        domain VARCHAR(128) COMMENT 'mirror second domain',        author_url VARCHAR(516) COMMENT 'mirror author homepage',        status INT(11),        _tx_id VARCHAR(255) COMMENT 'glitter transaction id',        FULLTEXT(title) WITH PARSER jieba,        FULLTEXT(title2) WITH PARSER standard,        FULLTEXT(body) WITH PARSER jieba,        FULLTEXT(body2) WITH PARSER standard,        FULLTEXT(language) WITH PARSER keyword,        FULLTEXT(category) WITH PARSER standard,        FULLTEXT(author) WITH PARSER keyword,        FULLTEXT(display_name) WITH PARSER standard,        FULLTEXT(ens) WITH PARSER standard,        FULLTEXT(published_time) WITH PARSER keyword,         FULLTEXT(domain) WITH PARSER standard        ) ENGINE=full_text COMMENT 'mirror article table';"
	ctx := context.Background()
	ast, e1 := ParseOneSQL(ctx, sql1)
	if e1 != nil {
		fmt.Println(e1)
		return
	}
	s, e2 := analyzer.FindDBTablID(ast)
	fmt.Println(e2)
	fmt.Println(s.Database)
	fmt.Println(s.Table)
}
