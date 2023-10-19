以下クエリで取得し下記レスポンスに加工

```
SELECT *  FROM environment_variables WHERE tenant_id = "t1" ORDER BY env_key, default_flg desc
```

レスポンス<br>
※1 envList は、一覧に出すデータ(デフォルト値)<br>
※2 envDetailMap は、envKey を元に参照(デフォルト値が一番上)

```
{
"data": {
  "envList": [
    {
      "id": "1",
      "envKey": "ek1",
      "envValue": "./ddd/ddd"
    },
    {
      "id": "4",
      "envKey": "ek2",
      "envValue": "{ddd}"
    }
  ],
  "envDetailMap": {
    "ek1": [
      {
        "id": "1",
        "tenantID": "t1",
        "envKey": "ek1",
        "stageID": "default",
        "envValue": "./ddd/ddd",
        "defaultFLG": 1,
        "updatedAt": "2023-10-17T09:26:28Z",
        "createdAt": "2023-10-19T09:26:28Z"
      },
      {
        "id": "2",
        "tenantID": "t1",
        "envKey": "ek1",
        "stageID": "stage_a",
        "envValue": "./aaa/aaa",
        "defaultFLG": null,
        "updatedAt": "2023-10-19T09:26:52Z",
        "createdAt": "2023-10-19T09:26:52Z"
      },
      {
        "id": "3",
        "tenantID": "t1",
        "envKey": "ek1",
        "stageID": "stage_b",
        "envValue": "./bbb/bbb",
        "defaultFLG": null,
        "updatedAt": "2023-10-19T09:28:05Z",
        "createdAt": "2023-10-19T09:28:05Z"
      }
    ],
    "ek2": [
      {
        "id": "4",
        "tenantID": "t1",
        "envKey": "ek2",
        "stageID": "default",
        "envValue": "{ddd}",
        "defaultFLG": 1,
        "updatedAt": "2023-10-18T09:28:05Z",
        "createdAt": "2023-10-19T09:28:05Z"
      },
      {
        "id": "5",
        "tenantID": "t1",
        "envKey": "ek2",
        "stageID": "stage_a",
        "envValue": "{aaa}",
        "defaultFLG": null,
        "updatedAt": "2023-10-19T09:28:05Z",
        "createdAt": "2023-10-19T09:28:05Z"
      }
    ]
  }
}
}
```
