package aefire

/*func MessageToToken(token, title, body, priority string, notRegisteredTokenHandler func(fcmToken FcmToken), datas ...string) error {
	data := StringMapOf(datas...)
	data["title"] = title
	data["body"] = body

	_, err := FCM.Send(context.Background(), &messaging.Message{
		Token: token,
		Android: &messaging.AndroidConfig{
			Priority: priority,
		},
		Data: data,
	})

	if err != nil &&
		messaging.IsRegistrationTokenNotRegistered(err) &&
		notRegisteredTokenHandler != nil {
		notRegisteredTokenHandler(FcmToken(token))
	}

	return err
}


func MessageToTokens(tokens []string, title, body, priority string, notRegisteredTokensHandler func([]FcmToken), datas ...string) map[string]error {
	result := map[string]error{}

	wg := sync.WaitGroup{}
	wg.Add(len(tokens))

	notRegisteredTokens := []FcmToken{}

	for _, token := range tokens {
		go func(token string) {
			defer wg.Done()
			result[token] = MessageToToken(token, title, body, priority, func(s FcmToken) {
				notRegisteredTokens = append(notRegisteredTokens, s)
			}, datas...)
		}(token)
	}

	wg.Wait()

	if notRegisteredTokensHandler != nil {
		notRegisteredTokensHandler(notRegisteredTokens)
	}

	return result
}

func MessageToTopic(topic, title, body, priority string, datas ...string) error {

	data := StringMapOf(datas...)
	data["title"] = title
	data["body"] = body

	_, err := FCM.Send(context.Background(), &messaging.Message{
		Topic: topic,
		Android: &messaging.AndroidConfig{
			Priority: priority,
		},
		Data: data,
	})

	return err
}
*/
