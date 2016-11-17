package oi

type Question struct{
    Message string
    ChatID int
}

type Answer struct{
    Message string
    Suggestions []string
    Subject string
    Conversation []Message
}

