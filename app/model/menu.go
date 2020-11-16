// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

// 顶部菜单数据结构
type TopMenuItem struct {
	Name   string         `json:"name"`   // 显示名称
	Url    string         `json:"url"`    // 链接地址
	Target string         `json:"target"` // 打开方式: 空, _blank
	Active bool           `json:"active"` // 是否被选中
	Items  []*TopMenuItem `json:"items"`  // 子级菜单
}
