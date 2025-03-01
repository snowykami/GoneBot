package keyboard

import "fmt"

var AutoId int = 0

// ActionType 按钮操作类型
type ActionType uint32

// PermissionType 按钮的权限类型
type PermissionType uint32

const (
	// ActionTypeURL http 或 小程序 客户端识别 schema, data字段为链接
	ActionTypeURL ActionType = 0
	// ActionTypeCallback 回调互动回调地址, data 传给互动回调地址
	ActionTypeCallback ActionType = 1
	// ActionTypeAtBot at机器人, 根据 at_bot_show_channel_list 决定在当前频道或用户选择频道,自动在输入框 @bot data
	ActionTypeAtBot ActionType = 2
	// PermissionTypeSpecifyUserIDs 仅指定这条消息的人可操作
	PermissionTypeSpecifyUserIDs PermissionType = 0
	// PermissionTypManager  仅频道管理者可操作
	PermissionTypManager PermissionType = 1
	// PermissionTypAll  所有人可操作
	PermissionTypAll PermissionType = 2
	// PermissionTypSpecifyRoleIDs 指定身份组可操作
	PermissionTypSpecifyRoleIDs PermissionType = 3
)

// MessageKeyboard 消息按钮组件
type MessageKeyboard struct {
	ID      string          `json:"id,omitempty"`      // 消息按钮组件模板 ID
	Content *CustomKeyboard `json:"content,omitempty"` // 消息按钮组件自定义内容
}

// CustomKeyboard 自定义 Keyboard
type CustomKeyboard struct {
	Rows []*Row `json:"rows,omitempty"` // 行数组
}

// Row 每行结构
type Row struct {
	Buttons []*Button `json:"buttons,omitempty"` // 每行按钮
}

// Button 单个按纽
type Button struct {
	ID         string      `json:"id,omitempty"`          // 按钮 ID
	RenderData *RenderData `json:"render_data,omitempty"` // 渲染展示字段
	Action     *Action     `json:"action,omitempty"`      // 该按纽操作相关字段
}

// RenderData  按纽渲染展示
type RenderData struct {
	Label        string `json:"label,omitempty"`         // 按纽上的文字
	VisitedLabel string `json:"visited_label,omitempty"` // 点击后按纽上文字
	Style        int    `json:"style,omitempty"`         // 按钮样式，0：灰色线框，1：蓝色线框
}

// Action 按纽点击操作
type Action struct {
	Type                 ActionType  `json:"type,omitempty"`           // 操作类型
	Permission           *Permission `json:"permission,omitempty"`     // 可操作
	UnsupportTips        string      `json:"unsupport_tips,omitempty"` // 可点击的次数, 默认不限
	Data                 string      `json:"data,omitempty"`           // 操作相关数据
	Reply                bool        `json:"reply,omitempty"`          // false:当前 true:弹出展示子频道选择器
	Enter                bool        `json:"enter,omitempty"`
	AtBotShowChannelList bool        `json:"at_bot_show_channel_list,omitempty"`
}

// Permission 按纽操作权限
type Permission struct {
	// Type 操作权限类型
	Type PermissionType `json:"type,omitempty"`
	// SpecifyRoleIDs 身份组
	SpecifyRoleIDs []string `json:"specify_role_ids,omitempty"`
	// SpecifyUserIDs 指定 UserID
	SpecifyUserIDs []string `json:"specify_user_ids,omitempty"`
}

func NewRow() *Row {
	return &Row{
		Buttons: make([]*Button, 0),
	}
}

func NewKeyBoard() *CustomKeyboard {
	return &CustomKeyboard{
		Rows: make([]*Row, 0),
	}
}

// 通用按钮，自由度最高
func (r *Row) Button(label, visitedLabel, data string, style, actionType, permissionType int, reply, enter, atBotShowChannelList bool) *Row {
	AutoId++
	id := fmt.Sprintf("%v", AutoId)
	r.Buttons = append(r.Buttons, &Button{
		ID: id,
		RenderData: &RenderData{
			Label:        label,
			VisitedLabel: visitedLabel,
			Style:        style,
		},
		Action: &Action{
			Type: ActionType(actionType),
			Permission: &Permission{
				Type: PermissionType(permissionType),
			},
			Data:                 data,
			Reply:                reply,
			Enter:                enter,
			AtBotShowChannelList: atBotShowChannelList,
		},
	})
	return r
}

// 文本按钮，所有人可用
func (r *Row) TextButton(label, visitedLabel, data string, reply, enter bool) *Row {
	AutoId++
	id := fmt.Sprintf("%v", AutoId)
	r.Buttons = append(r.Buttons, &Button{
		ID: id,
		RenderData: &RenderData{
			Label:        label,
			VisitedLabel: visitedLabel,
			Style:        0,
		},
		Action: &Action{
			Type: ActionTypeAtBot,
			Permission: &Permission{
				Type: PermissionTypAll,
			},
			Data:                 data,
			Reply:                reply,
			Enter:                enter,
			AtBotShowChannelList: false,
		},
	})
	return r
}

// 文本按钮，管理可用
func (r *Row) TextButtonAdmin(label, visitedLabel, data string, reply, enter bool) *Row {
	AutoId++
	id := fmt.Sprintf("%v", AutoId)
	r.Buttons = append(r.Buttons, &Button{
		ID: id,
		RenderData: &RenderData{
			Label:        label,
			VisitedLabel: visitedLabel,
			Style:        0,
		},
		Action: &Action{
			Type: ActionTypeAtBot,
			Permission: &Permission{
				Type: PermissionTypManager,
			},
			Data:                 data,
			Reply:                reply,
			Enter:                enter,
			AtBotShowChannelList: false,
		},
	})
	return r
}

// 链接按钮，所有人可用
func (r *Row) UrlButton(label, visitedLabel, url string, reply, enter bool) *Row {
	AutoId++
	id := fmt.Sprintf("%v", AutoId)
	r.Buttons = append(r.Buttons, &Button{
		ID: id,
		RenderData: &RenderData{
			Label:        label,
			VisitedLabel: visitedLabel,
			Style:        0,
		},
		Action: &Action{
			Type: ActionTypeURL,
			Permission: &Permission{
				Type: PermissionTypAll,
			},
			Data:                 url,
			Reply:                reply,
			Enter:                enter,
			AtBotShowChannelList: false,
		},
	})
	return r
}

// 链接按钮，管理可用
func (r *Row) UrlButtonAdmin(label, visitedLabel, url string, reply, enter bool) *Row {
	AutoId++
	id := fmt.Sprintf("%v", AutoId)
	r.Buttons = append(r.Buttons, &Button{
		ID: id,
		RenderData: &RenderData{
			Label:        label,
			VisitedLabel: visitedLabel,
			Style:        0,
		},
		Action: &Action{
			Type: ActionTypeURL,
			Permission: &Permission{
				Type: PermissionTypManager,
			},
			Data:                 url,
			Reply:                reply,
			Enter:                enter,
			AtBotShowChannelList: false,
		},
	})
	return r
}

func (k *CustomKeyboard) Row(r *Row) *CustomKeyboard {
	k.Rows = append(k.Rows, r)
	return k
}

func (k *CustomKeyboard) ResetAutoId() {
	AutoId = 0
}
