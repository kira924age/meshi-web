package main

import (
	"encoding/json"
	"log"
)

type Meshi struct {
	Res         string `json:"res"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsClose     bool   `json:"isClose"`
}

func getMeshiByte(meshiStr string) []byte {
	mp := map[string]*Meshi{
		"shokujin": &Meshi{
			Res:         meshiStr,
			Title:       "食神餃子王",
			Description: "食神餃子王. 安くて量が多い中華料理店",
			IsClose:     false,
		},
		"ootoya": &Meshi{
			Res:         meshiStr,
			Title:       "大戸屋ごはん処 調布北口店",
			Description: "大戸屋ごはん処. 西友の近くにある定食屋",
			IsClose:     false,
		},
		"soramame": &Meshi{
			Res:         meshiStr,
			Title:       "そらまめらぁめん本舗",
			Description: "そらまめらぁめん本舗. ラーメン屋, 丼ものもある. 学生証提示で大盛り無料だった気がする",
			IsClose:     false,
		},
		"lunchtime": &Meshi{
			Res:         meshiStr,
			Title:       "らんちたいむ",
			Description: "らんちたいむ. 弁当屋さん(閉業)",
			IsClose:     true,
		},
		"gakushoku": &Meshi{
			Res:         meshiStr,
			Title:       "電気通信大学 生協食堂",
			Description: "電気通信大学 生協食堂. 電通大の学食. コスパはあまりよくない",
			IsClose:     false,
		},
		"royal": &Meshi{
			Res:         meshiStr,
			Title:       "ロイヤルホスト 調布店",
			Description: "ロイヤルホスト 調布店. ファミレス",
			IsClose:     true,
		},
		"nisishoku-den": &Meshi{
			Res:         meshiStr,
			Title:       "西しょく den",
			Description: "西しょく den. 電通大西地区の食堂. 運営は生協とは別の業者",
			IsClose:     false,
		},
		"jiro-yaen": &Meshi{
			Res:         meshiStr,
			Title:       "二郎野猿",
			Description: "二郎野猿. 八王子にあるらしい",
			IsClose:     false,
		},
		"jiro-fuchu": &Meshi{
			Res:         meshiStr,
			Title:       "ラーメン二郎 府中店",
			Description: "ラーメン二郎 府中店",
			IsClose:     false,
		},
		"jiro-sengawa": &Meshi{
			Res:         meshiStr,
			Title:       "ラーメン二郎 仙川店",
			Description: "ラーメン二郎 仙川店",
			IsClose:     false,
		},
		"saburou": &Meshi{
			Res:         meshiStr,
			Title:       "郎郎郎 調布店",
			Description: "郎郎郎 調布店. 味が濃いラーメン屋. 量が多い",
			IsClose:     false,
		},
		"mac": &Meshi{
			Res:         meshiStr,
			Title:       "マクドナルド ２０号調布店",
			Description: "マクナル.",
			IsClose:     false,
		},
		"tenya": &Meshi{
			Res:         meshiStr,
			Title:       "天丼てんや 調布とうきゅう店",
			Description: "天丼てんや 調布とうきゅう店. 天丼屋さん",
			IsClose:     false,
		},
		"sutadon": &Meshi{
			Res:         meshiStr,
			Title:       "伝説のすた丼屋 調布店",
			Description: "伝説のすた丼屋 調布店",
			IsClose:     false,
		},
		"matsuya": &Meshi{
			Res:         meshiStr,
			Title:       "松屋 調布駅前店",
			Description: "松屋 調布駅前店",
			IsClose:     false,
		},
		"sennen": &Meshi{
			Res:         meshiStr,
			Title:       "千年ラーメン",
			Description: "千年ラーメン. 九州系のラーメン屋さん",
			IsClose:     false,
		},
		"saize": &Meshi{
			Res:         meshiStr,
			Title:       "サイゼリヤ 調布駅前店",
			Description: "サイゼリヤ 調布駅前店",
			IsClose:     false,
		},
		"koizushi": &Meshi{
			Res:         meshiStr,
			Title:       "鯉寿司 調布店",
			Description: "鯉寿司 調布店. お寿司屋さん",
			IsClose:     false,
		},
		"bigboy": &Meshi{
			Res:         meshiStr,
			Title:       "ビッグボーイ 調布店",
			Description: "ビッグボーイ 調布店. ステーキ屋さん",
			IsClose:     false,
		},
		"gasuto": &Meshi{
			Res:         meshiStr,
			Title:       "ガスト 調布駅前店",
			Description: "ガスト 調布駅前店. 2016年11月23日（水）閉店",
			IsClose:     true,
		},
		"jonasan": &Meshi{
			Res:         meshiStr,
			Title:       "ジョナサン 調布駅前店",
			Description: "ジョナサン 調布駅前店",
			IsClose:     false,
		},
		"karendo": &Meshi{
			Res:         meshiStr,
			Title:       "かれんど",
			Description: "かれんど. カレー屋さん",
			IsClose:     false,
		},
		"raja": &Meshi{
			Res:         meshiStr,
			Title:       "インド・ネパール料理 Raja 調布店",
			Description: "インド・ネパール料理店 Raja 調布店",
			IsClose:     false,
		},
		"coco1": &Meshi{
			Res:         meshiStr,
			Title:       "CoCo壱番屋 京王調布駅北口店",
			Description: "CoCo壱番屋 京王調布駅北口店. カレー屋さん",
			IsClose:     false,
		},
		"C&C": &Meshi{
			Res:         meshiStr,
			Title:       "カレーショップ C&C 調布南口店",
			Description: "カレーショップ C&C 調布南口店. カレー屋さん",
			IsClose:     false,
		},
		"misdo": &Meshi{
			Res:         meshiStr,
			Title:       "ミスタードーナツ 調布パルコショップ",
			Description: "ミスタードーナツ 調布パルコショップ. ドーナツ屋さん",
			IsClose:     false,
		},
		"sukekaku": &Meshi{
			Res:         meshiStr,
			Title:       "助格家 調布店",
			Description: "助格家 調布店. ラーメン",
			IsClose:     false,
		},
		"tonsuke": &Meshi{
			Res:         meshiStr,
			Title:       "とん助",
			Description: "とん助. ラーメン屋さん",
			IsClose:     true,
		},
		"takeya": &Meshi{
			Res:         meshiStr,
			Title:       "竹屋 調布銀座店",
			Description: "竹屋 調布銀座店. つけ麺やさん",
			IsClose:     false,
		},
		"tatsumi": &Meshi{
			Res:         meshiStr,
			Title:       "たつみ",
			Description: "たつみ. ラーメン屋さん",
			IsClose:     false,
		},
		"matsuri": &Meshi{
			Res:         meshiStr,
			Title:       "祭",
			Description: "祭. 西調布にある居酒屋",
			IsClose:     false,
		},
		"cowboyfamily": &Meshi{
			Res:         meshiStr,
			Title:       "カウボーイ家族 調布店",
			Description: "カウボーイ家族 調布店. ファミレス. ステーキとかハンバーグとかありそう",
			IsClose:     false,
		},
		"kurazushi": &Meshi{
			Res:         meshiStr,
			Title:       "無添くら寿司 調布つつじヶ丘店",
			Description: "無添くら寿司 調布つつじヶ丘店. お寿司屋さん",
			IsClose:     false,
		},
		"chatnoir": &Meshi{
			Res:         meshiStr,
			Title:       "コーヒーハウス・シャノアール 調布店",
			Description: "コーヒーハウス・シャノアール 調布店. コーヒー屋さん",
			IsClose:     false,
		},
		"starbucks": &Meshi{
			Res:         meshiStr,
			Title:       "スターバックス コーヒー",
			Description: "スターバックス コーヒー. 調布駅周辺に2店舗存在.",
			IsClose:     false,
		},
		"boteju": &Meshi{
			Res:         meshiStr,
			Title:       "ぼてぢゅう 調布パルコ店",
			Description: "ぼてぢゅう 調布パルコ店. お好み焼き屋さん",
			IsClose:     false,
		},
		"tonchinkan": &Meshi{
			Res:         meshiStr,
			Title:       "豚珍館 東口店",
			Description: "豚珍館 東口店. とんかつ屋さん",
			IsClose:     false,
		},
		"dandadan": &Meshi{
			Res:         meshiStr,
			Title:       "肉汁餃子製作所ダンダダン酒場",
			Description: "肉汁餃子製作所ダンダダン酒場. 居酒屋",
			IsClose:     false,
		},
		"honorurushokudo": &Meshi{
			Res:         meshiStr,
			Title:       "ホノルル食堂 調布",
			Description: "ホノルル食堂 調布.",
			IsClose:     false,
		},
		"egao": &Meshi{
			Res:         meshiStr,
			Title:       "東京エガオ食堂",
			Description: "東京エガオ食堂. 2019年9月23日（月・祝）閉店",
			IsClose:     true,
		},
		"kenta": &Meshi{
			Res:         meshiStr,
			Title:       "ケンタッキーフライドチキン 調布店",
			Description: "ケンタッキーフライドチキン 調布店",
			IsClose:     false,
		},
		"wakamatsuya": &Meshi{
			Res:         meshiStr,
			Title:       "そば処 若松屋",
			Description: "そば処 若松屋. そばや",
			IsClose:     false,
		},
		"sakananomanma": &Meshi{
			Res:         meshiStr,
			Title:       "魚の飯 調布店",
			Description: "魚の飯 調布店. 居酒屋っぽい",
			IsClose:     false,
		},
		"ohara": &Meshi{
			Res:         meshiStr,
			Title:       "おはら",
			Description: "おはら. 中華料理屋",
			IsClose:     false,
		},
		"shirataka": &Meshi{
			Res:         meshiStr,
			Title:       "しらたか",
			Description: "しらたか. ラーメン屋",
			IsClose:     true,
		},
		"hidakaya": &Meshi{
			Res:         meshiStr,
			Title:       "日高屋 調布北口店",
			Description: "日高屋 調布北口店. 熱烈中華食堂",
			IsClose:     false,
		},
		"nakau": &Meshi{
			Res:         meshiStr,
			Title:       "なか卯 調布北口店",
			Description: "なか卯 調布北口店. 丼もの",
			IsClose:     false,
		},
		"yoshinoya": &Meshi{
			Res:         meshiStr,
			Title:       "吉野家 調布駅前店",
			Description: "吉野家 調布駅前店. 牛丼屋さん",
			IsClose:     false,
		},
		"yukaku": &Meshi{
			Res:         meshiStr,
			Title:       "牛角 調布店",
			Description: "牛角 調布店. 焼肉屋さん",
			IsClose:     false,
		},
		"akechan": &Meshi{
			Res:         meshiStr,
			Title:       "伸ちゃん",
			Description: "おそらく, 西調布にある伸ちゃん(しんちゃん)という定食屋.",
			IsClose:     false,
		},
		"gyushige": &Meshi{
			Res:         meshiStr,
			Title:       "牛繁 調布店",
			Description: "牛繁 調布店. 焼肉屋さん",
			IsClose:     false,
		},
		"chofudojo": &Meshi{
			Res:         meshiStr,
			Title:       "調布道場",
			Description: "調布道場. ラーメン屋",
			IsClose:     false,
		},
		"shabuyou": &Meshi{
			Res:         meshiStr,
			Title:       "しゃぶ葉 調布南口店",
			Description: "しゃぶ葉 調布南口店. しゃぶしゃぶ",
			IsClose:     false,
		},
		"gyoza-taishoken": &Meshi{
			Res:         meshiStr,
			Title:       "餃子の大勝軒 調布店",
			Description: "餃子の大勝軒 調布店. 中華料理屋",
			IsClose:     false,
		},
	}
	ret := mp[meshiStr]
	jsonBytes, err := json.MarshalIndent(ret, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return jsonBytes
}
