# This is lib and tool for autogenerate i18n


- Build binary tool in [i18gen](./cmd/i18ngen/)
- To test gen data file use in [test](./test/)

# Structure of lang folder 
- All language file place in lang folder 
- Language file must be json file with format [lang].json
- In each lang file have multi messsage id with lang 
```
{
    "$message_id_1" : {
        "description" : "descrition of message id",
        "other" : "default message id content",
        "one" :"this is content for singular object ex: {{.Name}} comment on your post",
        "two" :"this is for double object, ex : {.Name}} and {{.Name2}} comment to your post",
        "many" : "{{.Name}}, {{.Name2}} and {{.total}} other comment on your post",
    },
     "$message_id_2" : {
        "description" : "descrition of message id",
        "other" : "default message id content",
        "one" :"this is content for singular object ex: {{.Name}} comment on your post",
        "two" :"this is for double object, ex : {.Name}} and {{.Name2}} comment to your post",
        "many" : "{{.Name}}, {{.Name2}} and {{.total}} other comment on your post",
    }
    .....
}
```
- You can omit fields  "one","two" and "many" but field "other" must required


# About tool gen i18n
- Generate and validate all MesssageID and all language in lang folder
- Generate test file to check valid message result