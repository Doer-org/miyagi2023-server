package mock

import (
	"time"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
	"github.com/google/uuid"
)

var spot *model.Spot = &model.Spot{
	ID:        uuid.New(),
	Name:      "hoge mock",
	Detail:    "hogehoge",
	Like:      23,
	ImgURL:    "http://example.com",
	Address:   "hoge",
	CreatedAt: time.Now(),
}

var spots []*model.Spot = []*model.Spot{
	{
		ID:        uuid.New(),
		Name:      "hoge1 mock",
		Detail:    "hogehoge1",
		Like:      123,
		ImgURL:    "http://example.com",
		Address:   "hoge",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "hoge2 mock",
		Detail:    "hogehoge2",
		Like:      223,
		ImgURL:    "http://example.com",
		Address:   "hoge",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "hoge3 mock",
		Detail:    "hogehoge3",
		Like:      323,
		ImgURL:    "http://example.com",
		Address:   "hoge",
		CreatedAt: time.Now(),
	},
	//ここから実在spots
	{
		ID:        uuid.New(),
		Name:      "秋保大滝",
		Detail:    "秋保大滝は、宮城県仙台市太白区秋保町にある滝。蔵王国定公園および県立自然公園二口峡谷の地域内を流れる名取川上流部において、落差55m、幅6mで流れ落ちる。 国の名勝に指定されており、日本の滝百選の1つにも数えられる。諸説あるが「日本三大名瀑」あるいは「日本三名瀑」の1つに数えられることがある。",
		Like:      323,
		ImgURL:    "https://www.sentabi.jp/wp-content/uploads/2019/11/1%E5%A4%A7%E6%BB%9D.jpg",
		Address:   "〒982-0244 宮城県仙台市太白区秋保町馬場字大滝 ",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "大崎八幡宮",
		Detail:    "大崎八幡宮は、宮城県仙台市青葉区八幡にある神社である。旧社格は村社。社殿は国宝に指定されており、どんと祭の裸参りで知られる。",
		Like:      456,
		ImgURL:    "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcQuka-Qf_5xgOlw9CaHFcIr1h0VLkyXYJgyNw6yciAIhExRCZ2HvEzJXYBUPX7ABZPrb0DjGmEcQOQcgGwL5Io4_w",
		Address:   "〒980-0871 宮城県仙台市青葉区八幡４丁目６−１",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "仙台うみの杜水族館",
		Detail:    "仙台うみの杜水族館は、宮城県仙台市宮城野区高砂中央公園内に所在する民間水族館である。かつて宮城県宮城郡松島町に存在したマリンピア松島水族館の後継館として2015年にオープンした。横浜八景島と仙台水族館開発が連携して運営する。",
		Like:      298,
		ImgURL:    "https://encrypted-tbn2.gstatic.com/licensed-image?q=tbn:ANd9GcSxovJPwnXIUw79Sa_fmV1UV1BTSApK2PZ_uATAVHOG4JY3PhGemSkJLlAYRsbNH2ryoVth6yq5H29QQKi1T_dt8g",
		Address:   "〒983-0013 宮城県仙台市宮城野区中野４丁目６",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "御釜",
		Detail:    "御釜は、宮城県刈田郡蔵王町と同県柴田郡川崎町の境界付近にある火口湖で、五色沼とも呼ばれる。蔵王連峰の中央部の最も標高の高いエリアにあり、しばしば蔵王連峰の象徴として見られている。",
		Like:      154,
		ImgURL:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTrCbww90TEpBV7uN9NiLSJqVN2tnVEYCvM55YMAGq0awMs4KojTBXNEFVFNy5qVMZ4ryycA2I4qAO87Ejnm-IS2w",
		Address:   "〒989-0916 宮城県刈田郡蔵王町遠刈田温泉倉石岳国有地内国有林内",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "鹽竈神社",
		Detail:    "鹽竈神社は、宮城県塩竈市にある神社である。志波彦神社との二社が同一境内に鎮座している。志波彦神社は式内社。鹽竈神社は式外社、陸奥国一宮。両社合わせて旧社格は国幣中社で、現在は神社本庁の別表神社。",
		Like:      223,
		ImgURL:    "https://tabi-mag.jp/wp-content/uploads/MG006901.jpg",
		Address:   "〒985-8510 宮城県塩竈市一森山１−１",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "瑞巌寺（臨済宗）",
		Detail:    "瑞巌寺は、宮城県宮城郡松島町にある臨済宗妙心寺派の仏教寺院である。",
		Like:      878,
		ImgURL:    "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcRiBX_NVMEa0r5nZzKeMy05YL8hSoXQtuX4sBPvKFox0j61n3poD1VR3P1-JhHLLuoQ4ybhYq0bQvhQ6eMreGvpCA",
		Address:   "〒981-0213 宮城県宮城郡松島町松島町内９１",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "円通院",
		Detail:    "円通院は、宮城県宮城郡松島町にある臨済宗妙心寺派の寺院である。本尊は聖観世音菩薩。瑞巌寺の南側に隣接してある。通称「バラ寺」「薔薇寺」。",
		Like:      83,
		ImgURL:    "https://encrypted-tbn3.gstatic.com/licensed-image?q=tbn:ANd9GcT_chqlrbYtOCTUR2d2Eeu9m1pFo0tRMjsX8avtKd5-hFd_TJc0UGfeHhhCdssFC3ee0Fcq3ADFCspyBCck3JByNg",
		Address:   "〒981-0213 宮城県宮城郡松島町松島町内６７−６７",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "瑞鳳殿",
		Detail:    "瑞鳳殿は、宮城県仙台市青葉区霊屋下にある仙台藩祖伊達政宗の霊廟である。広瀬川の蛇行部を挟んで、仙台城の本丸跡と向かい合う経ヶ峯に位置する。元々あった建物は戦災で焼失したが、後に復元された。 瑞鳳殿の周辺には伊達忠宗の霊廟である感仙殿、伊達綱宗の霊廟である善応殿、妙雲界廟、御子様御廟といった伊達氏に関連する霊廟や付属施設があり、一帯が「経ヶ峯伊達家墓所」として仙台市指定史跡となっている。また、瑞鳳寺が隣接する。",
		Like:      564,
		ImgURL:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSkstKO7w8iB867t_mnuuZaiAohEZPSsHnl3JKai3GpL1F75fccOp5DrwVHtpTUIH-yGtwULdWYhDEo8Kqq3xaymQ",
		Address:   "〒980-0814 宮城県仙台市青葉区霊屋下２３−２",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "仙台城跡（青葉城址）",
		Detail:    "仙台城は、宮城県仙台市青葉区の青葉山にあった日本の城。雅称は「青葉城」で「五城楼」の別名もある。国の史跡。",
		Like:      532,
		ImgURL:    "https://encrypted-tbn3.gstatic.com/licensed-image?q=tbn:ANd9GcSjNuEGC75SI2aOCIffiKW5JIopTnXHpxfVs3_zAFWcZqCwzg9f32TtVdh_CqmtP4l5dFITicoya1v6LUCcOmDIcQ",
		Address:   "〒980-0862 宮城県仙台市青葉区川内１",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "回顧橋 (かいこばし)",
		Detail:    "回顧橋（かいこばし）。大深沢橋の真下にある橋で、下から一気に見上げる迫力の景色を楽しむことができます",
		Like:      750,
		ImgURL:    "https://encrypted-tbn1.gstatic.com/licensed-image?q=tbn:ANd9GcRnE-DfmTBnB36Jje61R5WuZ5F9zBkqWo7B-F9VRBPrUKom2Ro9R6441tVtScrXn6WsUG3Mw6vzXyIJmHVOWqK0Jw",
		Address:   "〒989-6100 宮城県大崎市鳴子温泉",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "金蛇水神社",
		Detail:    "金蛇水神社は、宮城県岩沼市に鎮座する神社である。旧社格は無格社。",
		Like:      671,
		ImgURL:    "http://kanahebi.cdx.jp/img144.jpg",
		Address:   "〒989-2464 宮城県岩沼市三色吉水神７",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "石ノ森萬画館",
		Detail:    "石ノ森萬画館は、日本・宮城県石巻市中瀬に立地する、宮城県出身の漫画家・石ノ森章太郎の記念館である。英表記はIshinomori Manga Museum。 石ノ森作品の原画などを所蔵、展示する。",
		Like:      433,
		ImgURL:    "https://upload.wikimedia.org/wikipedia/commons/thumb/f/fd/MangattanMuseum002.jpg/1200px-MangattanMuseum002.jpg",
		Address:   "〒986-0823 宮城県石巻市中瀬２−７",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "白石川堤一目千本桜",
		Detail:    "桜が美しい景勝地",
		Like:      723,
		ImgURL:    "https://i2.wp.com/dora-tabi.com/wp-content/uploads/2021/09/a30946580622710179ace92a1c3a5a7c.jpg?ssl=1",
		Address:   "〒989-1603 宮城県柴田郡柴田町船岡１２",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "瑞巌寺五大堂",
		Detail:    "五大堂は、宮城県宮城郡松島町の景勝地松島にある仏堂。同町にある臨済宗妙心寺派の寺院・瑞巌寺の境外仏堂である。",
		Like:      883,
		ImgURL:    "https://encrypted-tbn2.gstatic.com/licensed-image?q=tbn:ANd9GcRwmOQghb_rstkqY0O54EVgeSx3HH2sr3h1frCrhP-0LKhBYfCwi54y7gglHNwuV5kWqet-uFypIIe-L35I4l7ABQ",
		Address:   "〒981-0213 宮城県宮城郡松島町松島町内１１１",
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "松島 (日本三景 )",
		Detail:    "松島は、宮城県北東部の松島湾内外にある約260の島々からなる諸島やそれを擁する多島海。または、湾周囲を囲む松島丘陵も含めた修景地区のことである。日本三景の一つに数えられている。2019年の観光入込客数は約598万人。",
		Like:      1072,
		ImgURL:    "https://www.nippon.com/ja/ncommon/contents/guide-to-japan/3811/3811.jpg",
		Address:   "〒981-0213 宮城県宮城郡松島町松島字町内115",
		CreatedAt: time.Now(),
	},
}

func NewSpotGetOutput() *usecase.SpotGetOutput {
	return &usecase.SpotGetOutput{
		Spot: spot,
	}
}

func NewSpotCreateOutput() *usecase.SpotCreateOutput {
	return &usecase.SpotCreateOutput{
		Spot: spot,
	}
}

func NewSpotListOutput() *usecase.SpotListOutput {
	return &usecase.SpotListOutput{
		Spots: spots,
	}
}
