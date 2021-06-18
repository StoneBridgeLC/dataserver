# Data Server

- /news
    - Response is all news.
- /news/range?from={id}&to={id}
    - Response is news in range in id.
- /news/:id
    - Response is News that id is :id.
- /news/month
    - Response is News that updated in a month.
- /topic
    - Response is all topics
- /topic/range?from={id}&to={id}
    - Response is topics in range in id.
- /topic/:id
    - Response is topic that id is :id.
- /news/:id/topic
    - Response is topic of news that id is :id.
- /comment
    - Response is all comments.
- /comment/range?from={id}&to={id}
    - Response is all comments in range in id.
- /comment/:id
    - Response is comment that id is :id.
- /news/:id/comment
    - Response is comments of news that id is :id.   
  
- GET /comment/unlabeled
    - Get unlabeled comment list.
    - example
      - request
        ```json 
        /comment/unlabeled?limit=5
        ```
      - response
        ```json
        [
          {
            "id": 1,
            "nid": 6,
            "body": "이게 맞지 삼성이 가전과 반도체 둘 다 가진 회사라면 엘지는 가전에 에스케이는 반도체에 집중해서 삼성과 경쟁하며 시너지 효과를 내는게 좋다. 하나 더 바라는 점이 있다면 sk는 과거에도 sky폰을 운영해 본 경험이 있으니 lg의 휴대폰사업을 이어받았으면 한다. 국산폰이 삼성 하나로 독점체제로 나가는 것은 소비자 입장에서 바람직하지 않다",
            "pid": {
              "Int64": 0,
              "Valid": false
            },
            "is_pos": {
              "Int64": 0,
              "Valid": false
            },
            "create_time": "2021-04-14T10:37:44Z",
            "update_time": "2021-04-14T10:37:44Z"
          },
          {
            "id": 2,
            "nid": 6,
            "body": "오예~~",
            "pid": {
              "Int64": 0,
              "Valid": false
            },
            "is_pos": {
              "Int64": 0,
              "Valid": false
            },
            "create_time": "2021-04-14T07:18:45Z",
            "update_time": "2021-04-14T07:18:45Z"
          },
          {
            "id": 3,
            "nid": 7,
            "body": "그래도 얘네는 탈당이라도 빠르네. 윤미향이도 아직 버티는 더듬당과는 다르군. 손혜원이나 이상직도 더럽게 버티다 나갔는데.",
            "pid": {
              "Int64": 0,
              "Valid": false
            },
            "is_pos": {
              "Int64": 0,
              "Valid": false
            },
            "create_time": "2021-04-14T14:41:05Z",
            "update_time": "2021-04-14T14:41:05Z"
          },
          {
            "id": 4,
            "nid": 7,
            "body": "탈당을 왜 시켜? 탈당 못하게 하고 출당시켰어야지. 자기 발로 나간 게 아니라 잘랐어야지.",
            "pid": {
              "Int64": 0,
              "Valid": false
            },
            "is_pos": {
              "Int64": 0,
              "Valid": false
            },
            "create_time": "2021-04-14T14:40:53Z",
            "update_time": "2021-04-14T14:40:53Z"
          },
          {
            "id": 5,
            "nid": 7,
            "body": "이제 탈당은 했고 폭행에대해 형사처벌 받으셔야지요.",
            "pid": {
              "Int64": 0,
              "Valid": false
            },
            "is_pos": {
              "Int64": 0,
              "Valid": false
            },
            "create_time": "2021-04-14T14:35:40Z",
            "update_time": "2021-04-14T14:35:40Z"
          }
        ]
        ```
  
- PUT /comment/label
    - Update comment label(is_pos)
    - example
      - request
      ```json
      /comment/label
      
      [
        {
          "id": 2566,
          "is_pos": 1
        },
        {
          "id": 2567,
          "is_pos": 0
        }
      ]
      ```
      - response : empty response body
      