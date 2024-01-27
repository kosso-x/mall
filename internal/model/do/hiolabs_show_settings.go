// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsShowSettings is the golang structure of table hiolabs_show_settings for DAO operations like Where/Data.
type HiolabsShowSettings struct {
	g.Meta         `orm:"table:hiolabs_show_settings, do:true"`
	Id             interface{} //
	Banner         interface{} // 滚动banner
	Channel        interface{} // 是否开启menu,几个图标
	IndexBannerImg interface{} // 首页的img图片是否显示
	Notice         interface{} //
}
