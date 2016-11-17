package main

import (
    "github.com/jinzhu/gorm"
)

func main()  {
    // conecta com o repositório das mensagens
    db, err := gorm.Open("postgres", name)
	if err != nil {
		panic("postgres panic")
	}

    defer db.Close()

    db.AutoMigrate(&Chat{})
    db.AutoMigrate(&Message{})

    //conecta com a base das mensagens
    client, err := elastic.NewClient()
    if err != nil {
        panic("elastic panic")
    }

    _, err = client.CreateIndex("oi").Do()
    if err != nil {
        panic(err)
    }

    //geração de uma pergunta qualquer
    question := Question{
        Message: "oi",
        ChatID: 1,
    }
    
    //verifica se ja tem algum chat já criado com aquele ID
    chat := Chat{
        ID: ChatID,
    }

    db.FirstOrCreate(&chat)

    search := Search{
        Question: question.Message,
    }
    
    //se ja tiver chat ele vai pegar a ultima mensagem para saber do assunto tratado(contexto)
    if len(chat.Messages) > 0 {
        search.QuestionSubject = chat.Messages[len(chat.Messages) - 1].Subject
    }

    //pesquisa por uma mensagem correspondente com o assunto e a mensagem
    searchResult, _ := client.Search().
        Index("oi").
        Query(elastic.NewTermQuery("question", question.Message)).
        Query(elastic.NewTermQuery("questionSubject", question.QuestionSubject)).
        From(0).Size(1).
        Do()

    
    //transforma em uma mensagem apartir da resposta do elasticsearch
    answerMessage := Message{
        Message: answerSearch.Message,
        Subject: answerSearch.AnswerSubject,
        ChatID: chat.ID,
    }

    //inseri essa mensagem no chat
    db.Create(&answerMessage)

    //pega todo o chat(essa parte não é realmente necessária, 
    //pode dar um append da obj mensagem em cima com o obj chat)
    db.First(&chat)

    //tranforma toda resposta anterior para uma resposta do usuário
    answerUser := Answer{
        Message: answerMessage.Message,
        Suggestions: searchResult.Suggestions,
        Subject: answerMessage.AnswerSubject,
        Conversation: chat.Messages,
    }
}