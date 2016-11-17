package oi

type Search struct {
	ID               int
	Question         string
	Answer           string
	QuestionSubjects []string
	AnswerSubjects   []string
	Suggestions      []string
}

// CONFIG
//
// {
//   "settings": {
//     "analysis": {
//       "filter": {
//         "brazilian_stop": {
//           "type":       "stop",
//           "stopwords":  "_brazilian_"
//         },
//         "brazilian_keywords": {
//           "type":       "keyword_marker",
//           "keywords":   []
//         },
//         "brazilian_stemmer": {
//           "type":       "stemmer",
//           "language":   "brazilian"
//         }
//       },
//       "analyzer": {
//         "brazilian": {
//           "tokenizer":  "standard",
//           "filter": [
//             "lowercase",
//             "brazilian_stop",
//             "brazilian_keywords",
//             "brazilian_stemmer"
//           ]
//         }
//       }
//     }
//   }
// }

// {
//   "mappings": {
//     "message" : {
//       "properties" : {
//         "question" : {
//             "type" :    "string",
//             "analyzer": "brazilian"
//         },
//         "questionSubjects": {
//             "type" :    "string",
//             "analyzer": "brazilian"
//         },
//         "answer" : {
//             "type":     "string",
//             "index":    "not_analyzed"
//         },
//         "answerSubjects": {
//             "type" :    "string",
//             "index":    "not_analyzed"
//         },
//         "suggestions" : {
//           "type" :   "string",
//           "index":    "not_analyzed"
//         }
//       }
//     }
//   }
// }

// {"question" : "oi",
// "questionSubjects": ["intro"],
// "answer" : "oi, tudo bem?",
// "answerSubjects": ["intro"],
// "suggestions" : ["sim", "não"]}

// {"question" : "fala ai",
// "questionSubjects": ["intro"],
// "answer" : "tranquilo, tudo bem?",
// "answerSubjects": ["intro"],
// "suggestions" : ["sim", "não"]}

// {
//   "query": {
//       "should" : [
//         { "match" : {
//             "questionSubjects" : "wow",
//             "boost": 3
//         } },
//         {
//             "match" : {
//                 "question" : "oi",
//                 "boost": 2
//         } }
//       ]
//     }
//   }
