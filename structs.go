package TextNowAPI

import "time"

type LoginPayload struct {
	Json string `json:"json"`
}

type LoginResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	GuidHex  string `json:"guid_hex"`
}

type MessagesResponse struct {
	Status struct {
		UserTimestamp   time.Time `json:"user_timestamp"`
		SettingsVersion string    `json:"settings_version"`
		FeaturesVersion string    `json:"features_version"`
		LatestMessageId int64     `json:"latest_message_id"`
		NumDevices      int       `json:"num_devices"`
	} `json:"status"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Id                    int64     `json:"id"`
	Username              string    `json:"username"`
	ContactValue          string    `json:"contact_value"`
	E164ContactValue      *string   `json:"e164_contact_value"`
	ContactType           int       `json:"contact_type"`
	ContactName           string    `json:"contact_name"`
	MessageDirection      int       `json:"message_direction"`
	MessageType           int       `json:"message_type"`
	Message               string    `json:"message"`
	Read                  bool      `json:"read"`
	Date                  time.Time `json:"date"`
	ConversationFiltering struct {
		FirstTimeContact           bool        `json:"first_time_contact"`
		PreviousNumberOwnerContact interface{} `json:"previous_number_owner_contact"`
	} `json:"conversation_filtering"`
}

type UserInformationResponse struct {
	UserId                        int64     `json:"user_id"`
	Username                      string    `json:"username"`
	AccountStatus                 string    `json:"account_status"`
	GuidHex                       string    `json:"guid_hex"`
	Expiry                        string    `json:"expiry"`
	Email                         string    `json:"email"`
	EmailVerified                 int       `json:"email_verified"`
	FirstName                     string    `json:"first_name"`
	LastName                      string    `json:"last_name"`
	CaptchaRequired               bool      `json:"captcha_required"`
	LastUpdate                    string    `json:"last_update"`
	Ringtone                      string    `json:"ringtone"`
	Signature                     string    `json:"signature"`
	ShowTextPreviews              bool      `json:"show_text_previews"`
	ForwardMessages               int       `json:"forward_messages"`
	IncentivizedShareDateTwitter  string    `json:"incentivized_share_date_twitter"`
	IncentivizedShareDateFacebook string    `json:"incentivized_share_date_facebook"`
	Credits                       int       `json:"credits"`
	Timestamp                     string 	`json:"timestamp"`
	PurchasesTimestamp            string 	`json:"purchases_timestamp"`
	HasPassword                   bool      `json:"has_password"`
	PhoneNumber                   string    `json:"phone_number"`
	PhoneAssignedDate             string    `json:"phone_assigned_date"`
	PhoneLastUnassigned           string    `json:"phone_last_unassigned"`
	Sip                           struct {
		Id           int    `json:"id"`
		Host         string `json:"host"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		Proxy        string `json:"proxy"`
		Stun         string `json:"stun"`
		VoicemailUrl string `json:"voicemail_url"`
		ClientIp     string `json:"client_ip"`
		IceServers   []struct {
			Endpoint   string `json:"endpoint"`
			Credential struct {
				Expiry   int    `json:"expiry"`
				Username string `json:"username"`
				Password string `json:"password"`
			} `json:"credential"`
		} `json:"ice_servers"`
	} `json:"sip"`
	DisableCalling                 string `json:"disable_calling"`
	MytempnumberDnd                bool   `json:"mytempnumber_dnd"`
	SipMinutes                     int    `json:"sip_minutes"`
	SipIP                          string `json:"sip_IP"`
	MytempnumberVoicemailUploadUrl string `json:"mytempnumber_voicemail_upload_url"`
	SipUsername                    string `json:"sip_username"`
	SipPassword                    string `json:"sip_password"`
	MytempnumberVoicemailV2        int    `json:"mytempnumber_voicemail_v2"`
	ReferringAmount                int    `json:"referring_amount"`
	ReferredAmount                 int    `json:"referred_amount"`
	MytempnumberStatus             int    `json:"mytempnumber_status"`
	MytempnumberExpiry             string `json:"mytempnumber_expiry"`
	Features                       struct {
		CdmaFallback bool `json:"cdma_fallback"`
		E911Accepted bool `json:"e911_accepted"`
		IsEmployee   bool `json:"is_employee"`
	} `json:"features"`
	ForwardingExpiry           string        `json:"forwarding_expiry"`
	ForwardingStatus           string        `json:"forwarding_status"`
	PremiumCalling             bool          `json:"premium_calling"`
	ForwardingNumber           string        `json:"forwarding_number"`
	Voicemail                  string        `json:"voicemail"`
	VoicemailTimestamp         string     	 `json:"voicemail_timestamp"`
	ShowAds                    bool          `json:"show_ads"`
	IsPersistent               bool          `json:"is_persistent"`
	MytempnumberFreeCalling    int           `json:"mytempnumber_free_calling"`
	IncentivizedShareDate      string        `json:"incentivized_share_date"`
	AppendFooter               int           `json:"append_footer"`
	AdsAutorenew               string        `json:"ads_autorenew"`
	VoiceAutorenew             string        `json:"voice_autorenew"`
	ForwardEmail               string        `json:"forward_email"`
	MytempnumberVoicemail      bool          `json:"mytempnumber_voicemail"`
	PhoneExpiry                string        `json:"phone_expiry"`
	AreaCode                   interface{}   `json:"area_code"`
	UnlimitedCalling           bool          `json:"unlimited_calling"`
	VmTranscriptionEnabled     bool          `json:"vm_transcription_enabled"`
	VmTranscriptionUserEnabled interface{}   `json:"vm_transcription_user_enabled"`
	AdCategories               []interface{} `json:"ad_categories"`
	MessagingEmail             string        `json:"messaging_email"`
}