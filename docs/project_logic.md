# Documentation

## Links to resources from which the concept of clean architecture is taken:

[![Link_number_1](https://img.shields.io/badge/Link_number_1-FF5555)](https://golang-blog.blogspot.com/2021/04/basic-project-layout-go-application.html)
[![Link_number_2](https://img.shields.io/badge/Link_number_2-FF00FF)](https://github.com/golang-standards/project-layout)
[![Link_number_3](https://img.shields.io/badge/Link_number_3-800085)](https://github.com/olezhek28/clean-architecture/tree/main)
[![Link_number_4](https://img.shields.io/badge/Link_number_4-0000FF)](https://habr.com/ru/companies/inDrive/articles/690088/)

#### - Some materials are not in English, but if they are translated, they will be presented very well.

---

## Description :white_check_mark:

Before describing what logic was implemented at each level, I would like to highlight one of 
the SOLID principles that I adhered to - the principle of "dependency inversion", which states 
that classes should depend on abstractions, not on specific details. In other words, this means 
that top-level abstractions should not depend on lower-level abstractions, usually the top level 
is the API level, it should not depend on the specific implementation of the repository level 
and the service level, that is, the base levels. And they should all be connected exclusively 
by abstractions, not by concrete implementations, but by abstractions. The same applies to the 
service level, it should not depend on the repository level.

--- 
<br>

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

## Services Part :fire:

![telegram bot logic](https://github.com/user-attachments/assets/ad39202f-0859-4878-96cd-90bedcf9fa63)

In the service layer, we have implemented all the basic logic of the bot's telegrams, in this layer the bot's operation is described in the message repeat mode and in the timer (alarm) mode of messages. The logic of these two tasks is fully implemented in the service layer, since I decided that accessing the database does not make sense for them, perhaps except for the timer (alarm) mode of messages, but I initially set myself the goal of making this bot functionality without resorting to the database in order to train and learn new material.

The only logic that required the repository layer (database accesses) This is the function of saving the links sent by the user and deleting these links by the user himself if desired. To implement this idea, the use of a database is ideal, I called this logic "link mode".

Although initially reviewing the code, it seems that the "link mode" mode in the bot required more lines of code than the "timer mode" mode, in fact, the code in the "link mode" mode looks more beautiful and more readable, SQL queries to the database are immediately visible there and this gives better readability.
