# 新华词典(dictionary-of-chinese)

> golang 版

> Author: xieWei

1. proverb // 谚语
2. word // 词语
3. idiom // 成语

### 预期功能

1. 如何设计 key (table)
2. 根据搜索排序： 谚语、词语、成语 (页面展示)
3. 统计：图表展示（谚语个数、词语个数、成语个数）；搜索占比
4. 分页显示


### 1. 谚语

> hash 

|谜面|谜底|
|:---|:---|
|竹子长杈|节外生枝|
| key| value|


- proverb:hash
- 搜索：根据谜面得出谜底


> API

```
GET /v1/api/proverb/keys/:key // 根据 key 搜索
GET /v1/api/proverb/ids/:id // 根据 id 搜索
GET /v1/api/proverb/samples?number=10&name=hjd  // 随机获取
```



### 2. 词语

> hash 


|词语|解释|
|:---|:---|
|阿q正传|中篇小说。鲁迅作。1921年发表。阿q是未庄的雇农，一贫如洗，但靠着精神胜利法”的麻醉而怡然自得。辛亥革命爆发后，他也开始神往于革命，但却遭假洋鬼子斥骂。不久，因赵秀才诬告，阿q被当作抢劫犯枪毙。小说揭示了贫苦农民的落后和愚昧，表达了作者改造国民性”的思想观点。|

- word:hash


> API

```
GET /v1/api/words/names/:name
GET /v1/api/words/ids/:id
GET /v1/api/words/samples?number=5&name=jdah
```

### 3. 成语

> string

|成语|拼音|释义|出处|示例|
|:---|:---|:---|:---|:---|
|阿其所好|ē qí suǒ hào|阿曲从；其他的；好爱好。指为取得某人的好感而迎合他的爱好。|《孟子·公孙丑上》宰我、子贡、有若，智足以知圣人，污不至阿其所好。”|吾何能～为？★鲁迅《坟·摩罗诗力说》|


|key|value|
|:---|:---|
|idiom:id:1:name|阿其所好|
|idiom:id:1:pinyin|ē qí suǒ hào|
|idiom:id:1:explain|阿曲从；其他的；好爱好。指为取得某人的好感而迎合他的爱好。|
|idiom:id:1:from|《孟子·公孙丑上》宰我、子贡、有若，智足以知圣人，污不至阿其所好。”|
|idiom:id:1:example|吾何能～为？★鲁迅《坟·摩罗诗力说》|


> hash

|key|filed|value|
|:---|:---|:---|
|idiom:1|name| |
|idiom:1|pinyin| |
|idiom:1|explain| |
|idiom:1|from| |
|idiom:1|example| |


> API

``` 
GET /v1/api/idioms/names/:name
GET /v1/api/idioms/ids/:id
GET /v1/api/idioms/samples?number=5&name=da 

```


### 4. 统计

> string , incr, decr


- proverb:count
- word:count
- idiom:count


### 5. 其他版本

> Python  版

[Python版](https://github.com/pwxcoo/chinese-xinhua)

### 6. License
MIT ©xieWei





