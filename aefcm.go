package aefire

import (
	"firebase.google.com/go/messaging"
	"strings"
	"sync"
)

type MessageBuilder messaging.Message

func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{
		Data:         nil,
		Notification: &messaging.Notification{},
		Android: &messaging.AndroidConfig{
			CollapseKey:           "",
			Priority:              "",
			TTL:                   nil,
			RestrictedPackageName: "",
			Data:                  StringMapOf(),
			Notification:          nil,
			FCMOptions:            nil,
		},
		Webpush: nil,
		APNS: &messaging.APNSConfig{
			Headers: StringMapOf(),
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					AlertString: "",
					Alert: &messaging.ApsAlert{
						Title:           "",
						SubTitle:        "",
						Body:            "",
						LocKey:          "",
						LocArgs:         nil,
						TitleLocKey:     "",
						TitleLocArgs:    nil,
						SubTitleLocKey:  "",
						SubTitleLocArgs: nil,
						ActionLocKey:    "",
						LaunchImage:     "",
					},
					Badge:            nil,
					Sound:            "",
					CriticalSound:    nil,
					ContentAvailable: false,
					MutableContent:   false,
					Category:         "",
					ThreadID:         "",
					CustomData:       nil,
				},
				CustomData: MapOf(),
			},
			FCMOptions: nil,
		},
		FCMOptions: nil,
		Token:      "",
		Topic:      "",
		Condition:  "",
	}
}

func (m *MessageBuilder) AsMessage() *messaging.Message {
	msg := messaging.Message(*m)
	return &msg
}

func (m *MessageBuilder) SetToken(t string) *MessageBuilder {
	m.Token = t

	return m
}

func (m *MessageBuilder) Title(title string) *MessageBuilder {
	m.Android.Data["title"] = title
	if m.Android.Notification != nil {
		m.Android.Notification.Title = title
	}
	m.APNS.Payload.Aps.Alert.Title = title

	return m
}

func (m *MessageBuilder) Channel(channel string) *MessageBuilder {
	if m.Android.Notification != nil {
		m.Android.Notification.ChannelID = channel
	}
	m.Android.Data["channel"] = channel

	return m
}

func (m *MessageBuilder) Body(body string) *MessageBuilder {
	m.Android.Data["body"] = body
	if m.Android.Notification != nil {
		m.Android.Notification.Body = body
	}
	m.APNS.Payload.Aps.Alert.Body = body

	return m
}

func (m *MessageBuilder) SetTopic(topic string) *MessageBuilder {
	m.Topic = topic

	return m
}

func (m *MessageBuilder) PutData(key, value string) *MessageBuilder {
	m.Android.Data[key] = value
	m.APNS.Payload.CustomData[key] = value

	return m
}

func (m *MessageBuilder) Priority(p string) *MessageBuilder {
	m.Android.Priority = p

	return m
}

func (a *AEFire) MessageToTokens(tokens []string, msg *messaging.Message) map[string]error {
	result := map[string]error{}

	wg := sync.WaitGroup{}
	wg.Add(len(tokens))

	m := sync.Mutex{}

	for _, token := range tokens {
		go func(token string, msg messaging.Message) {
			defer wg.Done()
			msg.Token = token
			_, err := a.FCM.Send(a, &msg)

			if err != nil {
				m.Lock()
				result[token] = err
				m.Unlock()
			}

		}(token, *msg)
	}

	wg.Wait()

	return result
}

func TokenIID(token string) string {
	if strings.Index(token, ":") > 0 {
		return token[:strings.Index(token, ":")]
	}

	return ""
}

func (a *AEFire) MessageToQueryResults(msg *messaging.Message, collectionName, tokenFieldName, queryFieldName, queryValue string) map[string]error {
	tokens := a.QueryStringField(a.Col(collectionName).Where(queryFieldName, "==", queryValue), tokenFieldName)

	return a.MessageToTokens(StringMapValuesToSlice(tokens), msg)
}
