package omdata

/*IUpgrade 实现数据升级接口 */
type IUpgrade interface {
	Upgrade() interface{}
}
