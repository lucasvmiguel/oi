package main

import (
	"bytes"
	"context"
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
)

type Search struct {
	ID               int
	Question         string
	Answer           string
	QuestionSubjects []string
	AnswerSubjects   []string
	Suggestions      []string
	Identity         string
}

func main() {
	// conecta com o repositório das mensagens
	// db, err := gorm.Open("postgres", name)
	// if err != nil {
	// 	panic("postgres panic")
	// }

	// defer db.Close()

	// db.AutoMigrate(&Chat{})
	// db.AutoMigrate(&Message{})

	//conecta com a base das mensagens
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		panic("elastic panic")
	}

	var lw bytes.Buffer
	lout := log.New(&lw, "LOGGER ", log.LstdFlags)

	var tw bytes.Buffer
	tout := log.New(&tw, "TRACER ", log.LstdFlags)

	elastic.SetInfoLog(lout)
	elastic.SetTraceLog(tout)

	client.CreateIndex("oi").Do(context.TODO())

	client.PutMapping().Index("oi").Type("message").BodyJson(
		map[string]interface{}{
			"properties": map[string]interface{}{
				"Question": map[string]interface{}{
					"type":     "string",
					"analyzer": "brazilian",
				},
				"QuestionSubjects": map[string]interface{}{
					"type":     "string",
					"analyzer": "brazilian",
				},
				"Answer": map[string]interface{}{
					"type":  "string",
					"index": "not_analyzed",
				},
				"AnswerSubjects": map[string]interface{}{
					"type":  "string",
					"index": "not_analyzed",
				},
				"Suggestions": map[string]interface{}{
					"type":  "string",
					"index": "not_analyzed",
				},
				"Identity": map[string]interface{}{
					"type":  "string",
					"index": "not_analyzed",
				},
			},
		}).Do(context.TODO())

	// client.Index().
	// 	Index("oi").
	// 	Type("message").
	// 	BodyJson(Search{
	// 		Question:         "oi",
	// 		Answer:           "oi, tudo bem?",
	// 		QuestionSubjects: []string{"intro"},
	// 		AnswerSubjects:   []string{"intro"},
	// 		Suggestions:      []string{"sim", "não"},
	// 		Identity:      	  "oi_simples",
	// 	}).
	// 	Do(context.TODO())

	// client.Index().
	// 	Index("oi").
	// 	Type("message").
	// 	BodyJson(Search{
	// 		Question:         "olá",
	// 		Answer:           "oi, tudo bem?",
	// 		QuestionSubjects: []string{"intro"},
	// 		AnswerSubjects:   []string{"intro"},
	// 		Suggestions:      []string{"sim", "não"},
	// 		Identity:      	  "ola_simples",
	// 	}).
	// 	Do(context.TODO())

	//geração de uma pergunta qualquer
	// question := Question{
	// 	Message: "oi",
	// 	ChatID:  1,
	// }

	// //verifica se ja tem algum chat já criado com aquele ID
	// chat := Chat{
	// 	ID: ChatID,
	// }

	// db.FirstOrCreate(&chat)

	// search := Search{
	// 	Question: question.Message,
	// }

	// //se ja tiver chat ele vai pegar a ultima mensagem para saber do assunto tratado(contexto)
	// if len(chat.Messages) > 0 {
	// 	search.QuestionSubject = chat.Messages[len(chat.Messages)-1].Subject
	// }

	// //pesquisa por uma mensagem correspondente com o assunto e a mensagem

	searchResult, err := client.Search().
		Query(
			elastic.NewBoolQuery().Should(
				elastic.NewMatchQuery("Question", "oi").Boost(100),
				elastic.NewMatchQuery("QuestionSubjects", "intro").Boost(100),
			)).
		From(0).
		Size(1).
		Do(context.TODO())
	// searchResult, _ := client.Search().
	// 	Index("oi").
	// 	Query(elastic.NewTermQuery("question", question.Message)).
	// 	Query(elastic.NewTermQuery("questionSubject", question.QuestionSubject)).
	// 	From(0).Size(1).
	// 	Do()

	// //transforma em uma mensagem apartir da resposta do elasticsearch
	// answerMessage := Message{
	// 	Message: answerSearch.Message,
	// 	Subject: answerSearch.AnswerSubject,
	// 	ChatID:  chat.ID,
	// }

	// //inseri essa mensagem no chat
	// db.Create(&answerMessage)

	// //pega todo o chat(essa parte não é realmente necessária,
	// //pode dar um append da obj mensagem em cima com o obj chat)
	// db.First(&chat)

	// //tranforma toda resposta anterior para uma resposta do usuário
	// answerUser := Answer{
	// 	Message:      answerMessage.Message,
	// 	Suggestions:  searchResult.Suggestions,
	// 	Subject:      answerMessage.AnswerSubject,
	// 	Conversation: chat.Messages,
	// }
}
