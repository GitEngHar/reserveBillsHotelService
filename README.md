### ローカル検証用コマンド

#### pubsubにテストメッセージを送信

```gcloud
gcloud pubsub topics publish reserve  --message "Good Morning"
```

#### pubsubにjsonメッセージを送信
```gcloud
gcloud pubsub topics publish reserve --message "{\"id\": 0, \"is_cancel\": false, \"hotel_id\": 1, \"user_id\": 12345, \"reserved_datetime\": 1700000000, \"checkin_datetime\": 1700604800}"
```

