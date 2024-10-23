# Documentation

## Links to resources from which the concept of clean architecture is taken:

[![Link_number_1](https://img.shields.io/badge/Link_number_1-FF5555)](https://golang-blog.blogspot.com/2021/04/basic-project-layout-go-application.html)
[![Link_number_2](https://img.shields.io/badge/Link_number_2-FF00FF)](https://github.com/golang-standards/project-layout)
[![Link_number_3](https://img.shields.io/badge/Link_number_3-800085)](https://github.com/olezhek28/clean-architecture/tree/main)
[![Link_number_4](https://img.shields.io/badge/Link_number_4-0000FF)](https://habr.com/ru/companies/inDrive/articles/690088/)

#### - Some materials are not in English, but if they are translated, they will be presented very well.

---

## Description :white_check_mark:

### API Part :star:

![executers](https://github.com/user-attachments/assets/77c8b9bb-4f8d-452d-ba4e-767299d67425)

Telegram supports 2 types of executers - Long Polling and Webhook.
Long Polling is sending requests to telegrams in order to receive
updates. And Webhook is the acceptance of requests from the telegram itself. 
We  will use Long Polling, since we will not have a huge
number of requests from users. Development using Long Polling goes a lot further faster, 
but the only negative is the load on your own application (in
in our case, a bot). But in our case, this minus does not play a big role.
So we will not create an API folder and describe our own API there for processing 
HTTP requests, as this will slow down development (so I believe that we are not 
violating the clean architecture). If you still have questions about this point, 
you can google it on the Internet, chat-GPT or go follow the link and 
read the article I proposed, translating it into English or any other language.

[![link_to_article](https://img.shields.io/badge/link_to_article-119812)](https://grammy.dev/ru/guide/deployment-types)
