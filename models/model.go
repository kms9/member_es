package models

//$intArr = ['send_num', 'receive_num', 'update_info_time', 'mass_msg_num', 'last_update_time', 'subscribe_time', 'id'];

type WxIdList struct {
	WxId int   `json:"wx_system_user_id_list" orm:"column(wx_system_user_id)"`
	Num  int64 `json:"num" orm:"column(num)"`
}

type WxUser struct {
	Id                 int64  `json:"id" orm:"pk;auto;column(id)"`
	WxOpenId           string `json:"wx_open_id" orm:"column(wx_open_id)"`
	UserTel            string `json:"user_tel" orm:"column(user_tel)"`
	UserAddress        string `json:"user_address" orm:"column(user_address)"`
	SubscribeTime      int    `json:"subscribe_time" orm:"column(subscribe_time)"`
	UnSubscribeTime    int    `json:"unsubscribe_time" orm:"column(unsubscribe_time)"`
	WxUserGroupId      int    `json:"wx_user_group_id" orm:"column(wx_user_group_id)"`
	WxSystemUserId     int    `json:"wx_system_user_id" orm:"column(wx_system_user_id)"`
	WxReplyUserId      int    `json:"wx_reply_context_id" orm:"column(wx_reply_context_id)"`
	RemarkName         string `json:"remark_name" orm:"column(remark_name)"`
	UserPosition       string `json:"user_position" orm:"column(user_position)"`
	IdDel              string `json:"is_del" orm:"column(is_del)"`
	Nickname           string `json:"nickname" orm:"column(nickname)"`
	Sex                int    `json:"sex" orm:"column(sex)"`
	RemarkSex          int    `json:"remark_sex" orm:"column(remark_sex)"`
	Country            int    `json:"country" orm:"column(country)"`
	RemarkCountry      int    `json:"remark_country" orm:"column(remark_country)"`
	Province           int    `json:"province" orm:"column(province)"`
	RemarkProvince     int    `json:"remark_province" orm:"column(remark_province)"`
	City               int    `json:"city" orm:"column(city)"`
	RemarkCity         int    `json:"remark_city" orm:"column(remark_city)"`
	Language           string `json:"language" orm:"column(language)"`
	Headimgurl         string `json:"headimgurl" orm:"column(headimgurl)"`
	RemarkHeadingurl   string `json:"remark_headimgurl" orm:"column(remark_headimgurl)"`
	ImageId            string `json:"image_id" orm:"column(image_id)"`
	LastUpdateTime     int    `json:"last_update_time" orm:"column(last_update_time)"`
	RemarkDesc         string `json:"remark_desc" orm:"column(remark_desc)"`
	Birthday           string `json:"birthday" orm:"column(birthday)"`
	SendNum            int    `json:"send_num" orm:"column(send_num)"`
	ReceiveNum         int    `json:"receive_num" orm:"column(receive_num)"`
	MassMsgNum         int    `json:"mass_msg_num" orm:"column(mass_msg_num)"`
	SubscribeSource    string `json:"subscribe_source" orm:"column(subscribe_source)"`
	Unionid            string `json:"unionid" orm:"column(unionid)"`
	UpdateInfoTime     int    `json:"update_info_time" orm:"column(update_info_time)"`
	Tag                string `json:"tag" orm:"column(tag)"`
	IsAuthorize        string `json:"is_authorize" orm:"column(is_authorize)"`
	SubscribeQrcodeVal string `json:"subscribe_qrcode_val" orm:"column(subscribe_qrcode_val)"`
	IsUnilever         string `json:"is_unilever" orm:"column(is_unilever)"`
	GpsCountry         int    `json:"gps_country" orm:"column(gps_country)"`
	GpsProvince        int    `json:"gps_province" orm:"column(gps_province)"`
	GpsCity            int    `json:"gps_city" orm:"column(gps_city)"`
	UserShopName       string `json:"user_shop_name" orm:"column(user_shop_name)"`
}

//wx_system_user_id, main_id, fans_id, member_id, id_string,create_time
type TagString struct {
	Id             int64  `json:"id" orm:"column(id)"`
	MainId         int    `json:"main_id" orm:"column(main_id)"`
	IdString       string `json:"id_string" orm:"column(id_string)"`
	WxSystemUserId int    `json:"wx_system_user_id" orm:"column(wx_system_user_id)"`
	FansId         int64  `json:"fans_id" orm:"column(fans_id)"`
	MemberId       int    `json:"member_id" orm:"column(member_id)"`
	CreateTime     int    `json:"create_time" orm:"column(create_time)"`
}

type UserTagRel struct {
	MainId     int        `json:"main_id" orm:"column(main_id)"`
	IdString   string     `json:"id_string" orm:"column(id_string)"`
	CreateTime int        `json:"create_time" orm:"column(create_time)"`
	MemberId   int        `json:"member_id" orm:"column(member_id)"`
	FansIdRel  *FansIdRel `json:"fans_id_rel"`
}

type FansIdRel struct {
	Name   string `json:"name"`
	Parent string `json:"parent"`
}

type AppCustomersSetUserTag struct {
	Id             int64  `json:"id" orm:"pk;auto;column(id)"`
	WxSystemUserId int    `json:"wx_system_user_id" orm:"column(wx_system_user_id)"`
	FansId         int64  `json:"fans_id" orm:"column(fans_id)"`
	MemberId       int64  `json:"member_id" orm:"column(member_id)"`
	IdString       string `json:"id_string" orm:"column(id_string)"`
	CatId          int    `json:"cat_id" orm:"column(cat_id)"`
	CreateTime     string `json:"create_time" orm:"column(create_time)"`
	IsDel          string `json:"is_del" orm:"column(is_del)"`
	Content        string `json:"content" orm:"column(content)"`
}

type EsModel struct {
	*WxUser
	TagNameList []string `json:"tagname"`
}
